
// based on go-libvirt

package codegen

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

// ConstItem stores an const's symbol and value from the parser. This struct is
// also used for enums.
type ConstItem struct {
	Name     string
	OrigName string
	Val      string
}

// Generator holds all the information parsed out of the protocol file.
type Generator struct {
	// Enums holds the enum declarations. The type of enums is always int32.
	Enums []Decl
	// EnumVals holds the list of enum values found by the parser. In sunrpc as
	// in go, these are not separately namespaced.
	EnumVals []ConstItem
	// Consts holds all the const items found by the parser.
	Consts []ConstItem
	// Structs holds a list of all the structs found by the parser
	Structs []Structure
	// StructMap is a map of the structs we find for quick searching.
	StructMap map[string]int
	// Typedefs holds all the type definitions from 'typedef ...' lines.
	Typedefs []Typedef
	// Unions holds all the discriminated unions.
	Unions []Union
	// UnionMap is a map of the unions we find for quick searching.
	UnionMap map[string]int
	// Procs holds all the discovered libvirt procedures.
	Procs []Proc

	ArgFunc []Func
}

// Gen accumulates items as the parser runs, and is then used to produce the
// output.
var Gen Generator

// CurrentEnumVal is the auto-incrementing value assigned to enums that aren't
// explicitly given a value.
var CurrentEnumVal int64

// goEquivTypes maps the basic types defined in the rpc spec to their golang
// equivalents.
var goEquivTypes = map[string]string{
	// Some of the identifiers in the rpc specification are reserved words or
	// pre-existing types in go. This renames them to something safe.
	"type":  "xtype",
	"error": "xerror",
	"nil":   "xnil",

	// The libvirt spec uses this NonnullString type, which is a string with a
	// specified maximum length. This makes the go code more confusing, and
	// we're not enforcing the limit anyway, so collapse it here. This also
	// requires us to ditch the typedef that would otherwise be generated.

	// TODO: Get rid of these. They're only needed because we lose information
	// that the parser has (the parser knows it has emitted a go type), and then
	// we capitalize types to make them public.
	"Int":     "int",
	"Uint":    "uint",
	"Int8":    "int8",
	"Uint8":   "uint8",
	"Int16":   "int16",
	"Uint16":  "uint16",
	"Int32":   "int32",
	"Uint32":  "uint32",
	"Int64":   "int64",
	"Uint64":  "uint64",
	"Float32": "float32",
	"Float64": "float64",
	"Bool":    "bool",
	"Byte":    "byte",
	"Int8T":    "int8",
	"Uint8T":   "uint8",
	"Int16T":   "int16",
	"Uint16T":  "uint16",
	"Int32T":   "int32",
	"Uint32T":  "uint32",
	"Int64T":   "int64",
	"Uint64T":  "uint64",
	"Float32T": "float32",
	"Float64T": "float64",
}

// Decl records a declaration, like 'int x' or 'remote_nonnull_string str'
type Decl struct {
	Name, OrigName, Type string
}

// NewDecl returns a new declaration struct.
func NewDecl(identifier, itype string) *Decl {
	goidentifier := identifierTransform(identifier)
	itype = typeTransform(itype)
	return &Decl{Name: goidentifier, OrigName: identifier, Type: itype}
}

// Structure records the name and members of a struct definition.
type Structure struct {
	Name     string
	OrigName string
	Members  []Decl
}

// Typedef holds the name and underlying type for a typedef.
type Typedef struct {
	Decl
}

// Union holds a "discriminated union", which consists of a discriminant, which
// tells you what kind of thing you're looking at, and a number of encodings.
type Union struct {
	Name             string
	//OrigName         string
	DiscriminantName string
	DiscriminantType string
	Cases            []Case
}

// Case holds a single case of a discriminated union.
type Case struct {
	CaseName        string
	DiscriminantVal string
	Decl
}

// Proc holds information about a libvirt procedure the parser has found.
type Proc struct {
	Num            int64  // The libvirt procedure number.
	Name           string // The name of the go func.
	OrigName       string // The name of the libvirt proc this wraps.
	Args           []Decl // The contents of the args struct for this procedure.
	Ret            []Decl // The contents of the ret struct for this procedure.
	ArgsStruct     string // The name of the args struct for this procedure.
	RetStruct      string // The name of the ret struct for this procedure.
	ReadStreamIdx  int    // The index of read stream in function argument list
	WriteStreamIdx int    // The index of read stream in function argument list
}

// ProcMeta holds information from annotations attached to a libvirt procedure
type ProcMeta struct {
	ReadStream  int
	WriteStream int
}

type Func struct {
	Name  string
	OpNum string
	Args  string
	Init  string
	Op    string
	Type  string
}

type structStack []*Structure

// CurrentStruct will point to a struct record if we're in a struct declaration.
// When the parser adds a declaration, it will be added to the open struct if
// there is one.
var CurrentStruct structStack

// Since it's possible to have an embedded struct definition, this implements
// a stack to keep track of the current structure.
func (s *structStack) empty() bool {
	return len(*s) == 0
}
func (s *structStack) push(st *Structure) {
	*s = append(*s, st)
}
func (s *structStack) pop() *Structure {
	if s.empty() {
		return nil
	}
	st := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return st
}
func (s *structStack) peek() *Structure {
	if s.empty() {
		return nil
	}
	return (*s)[len(*s)-1]
}

// CurrentTypedef will point to a typedef record if we're parsing one. Typedefs
// can define a struct or union type, but the preferred for is struct xxx{...},
// so we may never see the typedef form in practice.
var CurrentTypedef *Typedef

// CurrentUnion holds the current discriminated union record.
var CurrentUnion *Union

// CurrentCase holds the current case record while the parser is in a union and
// a case statement.
var CurrentCase *Case

var CurrentFunc *Func

// Generate will output go bindings for libvirt. The lvPath parameter should be
// the path to the root of the libvirt source directory to use for the
// generation.
func Generate(proto io.Reader, outFileName string) error {
	Gen.StructMap = make(map[string]int)
	Gen.UnionMap = make(map[string]int)
	lexer, err := NewLexer(proto)
	if err != nil {
		return err
	}
	go lexer.Run()
	parser := yyNewParser()
	yyErrorVerbose = true
	// Turn this on if you're debugging.
	//yyDebug = 3
	rv := parser.Parse(lexer)
	if rv != 0 {
		return fmt.Errorf("Error %v", rv)
	}

	// Generate and write the output.
	outFile, err := os.Create(outFileName)
	if err != nil {
		return err
	}
	defer outFile.Close()
	err = genGo(outFile)
	return err
}

// genGo is called when the parsing is done; it generates the golang output
// files using templates.
func genGo(file io.Writer) error {
	t, err := template.ParseFiles("yacc2go.tmpl")
	if err != nil {
		return err
	}
	if err = t.Execute(file, Gen); err != nil {
		return err
	}
	return nil
}

// constNameTransform changes an upcased, snake-style name like
// REMOTE_PROTOCOL_VERSION to a comfortable Go name like ProtocolVersion. It
// also tries to upcase abbreviations so a name like DOMAIN_GET_XML becomes
// DomainGetXML, not DomainGetXml.
func constNameTransform(name string) string {
	//decamelize := strings.ContainsRune(name, '_')
	//nn := strings.TrimPrefix(name, "REMOTE_")
	//if decamelize {
	//	nn = fromSnakeToCamel(nn)
	//}
	//nn = fixAbbrevs(nn)
	//return nn
	return name
}

func identifierTransform(name string) string {
	decamelize := strings.ContainsRune(name, '_')
	nn := strings.TrimPrefix(name, "remote_")
	if decamelize {
		nn = fromSnakeToCamel(nn)
	} else {
		nn = publicize(nn)
	}
	nn = fixAbbrevs(nn)
	nn = checkIdentifier(nn)
	// Many types in libvirt are prefixed with "Nonnull" to distinguish them
	// from optional values. We add "Opt" to optional values and strip "Nonnull"
	// because this makes the go code clearer.
	//nn = strings.TrimPrefix(nn, "Nonnull")
	return nn
}

func typeTransform(name string) string {
	nn := strings.TrimLeft(name, "*[]")
	diff := len(name) - len(nn)
	nn = identifierTransform(nn)
	return name[0:diff] + nn
}

func publicize(name string) string {
	if len(name) <= 0 {
		return name
	}
	r, n := utf8.DecodeRuneInString(name)
	name = string(unicode.ToUpper(r)) + name[n:]
	return name
}

// fromSnakeToCamel transmutes a snake-cased string to a camel-cased one. All
// runes that follow an underscore are up-cased, and the underscores themselves
// are omitted.
//
// ex: "PROC_DOMAIN_GET_METADATA" -> "ProcDomainGetMetadata"
func fromSnakeToCamel(s string) string {
	buf := make([]rune, 0, len(s))
	// Start rune will be upper case - we generate all public symbols.
	hump := true

	for _, r := range s {
		if r == '_' {
			hump = true
		} else {
			var transform func(rune) rune
			if hump == true {
				transform = unicode.ToUpper
			} else {
				transform = unicode.ToLower
			}
			buf = append(buf, transform(r))
			hump = false
		}
	}

	return string(buf)
}

// abbrevs is a list of abbreviations which should be all upper-case in a name.
// (This is really just to keep the go linters happy and to produce names that
// are intuitive to a go developer.)
var abbrevs = []string{"Xml", "Io", "Uuid", "Cpu", "Id", "Ip"}

// fixAbbrevs up-cases all instances of anything in the 'abbrevs' array. This
// would be a simple matter, but we don't want to upcase an abbreviation if it's
// actually part of a larger word, so it's not so simple.
func fixAbbrevs(s string) string {
	for _, a := range abbrevs {
		for loc := 0; ; {
			loc = strings.Index(s[loc:], a)
			if loc == -1 {
				break
			}
			r := 'A'
			if len(a) < len(s[loc:]) {
				r, _ = utf8.DecodeRune([]byte(s[loc+len(a):]))
			}
			if unicode.IsLower(r) == false {
				s = s[:loc] + strings.Replace(s[loc:], a, strings.ToUpper(a), 1)
			}
			loc++
		}
	}
	return s
}


//---------------------------------------------------------------------------
// Routines called by the parser's actions.
//---------------------------------------------------------------------------

// StartEnum is called when the parser has found a valid enum.
func StartEnum(name string) {
	// Enums are always signed 32-bit integers.
	goname := identifierTransform(name)
	Gen.Enums = append(Gen.Enums, Decl{goname, name, "int32"})
	// Set the automatic value var to -1; it will be incremented before being
	// assigned to an enum value.
	CurrentEnumVal = -1
}

// AddEnumVal will add a new enum value to the list.
func AddEnumVal(name, val string) error {
	ev, err := parseNumber(val)
	if err != nil {
		return fmt.Errorf("invalid enum value %v = %v", name, val)
	}
	return addEnumVal(name, ev, nil)
}

// AddEnumValMeta will add a new enum value with attached metadata to the list.
// Metadata is parsed from annotations in libvirt RPC description file that are
// in block comment preceding every function in enum, it looks like this:
// /**
//  * @generate: both
//  * @readstream: 1
//  * @sparseflag: VIR_STORAGE_VOL_DOWNLOAD_SPARSE_STREAM
//  * @acl: storage_vol:data_read
//  */
// See full description of possible annotations in libvirt's src/remote/remote_protocol.x
// at the top of remote_procedure enum.
// We're parsing only @readstream and @writestream annotations at the moment.
func AddEnumValMeta(name, val, meta string) error {
	ev, err := parseNumber(val)
	if err != nil {
		return fmt.Errorf("invalid enum value %v = %v", name, val)
	}
	metaObj, err := parseMeta(meta)
	if err != nil {
		return fmt.Errorf("invalid metadata for enum value %v: %v", name, err)
	}
	return addEnumVal(name, ev, metaObj)
}

// AddEnumAutoVal adds an enum to the list, using the automatically-incremented
// value. This is called when the parser finds an enum definition without an
// explicit value.
func AddEnumAutoVal(name string) error {
	CurrentEnumVal++
	return addEnumVal(name, CurrentEnumVal, nil)
}

func addEnumVal(name string, val int64, meta *ProcMeta) error {
	goname := constNameTransform(name)
	Gen.EnumVals = append(Gen.EnumVals, ConstItem{goname, name, fmt.Sprintf("%d", val)})
	CurrentEnumVal = val
	addProc(goname, name, val, meta)
	return nil
}

// AddConst adds a new constant to the parser's list.
func AddConst(name, val string) error {
	//_, err := parseNumber(val)
	//if err != nil {
	//	return fmt.Errorf("invalid const value %v = %v", name, val)
	//}
	goname := constNameTransform(name)
	Gen.Consts = append(Gen.Consts, ConstItem{goname, name, val})
	return nil
}

// addProc checks an enum value to see if it's a procedure number. If so, we
// add the procedure to our list for later generation.
func addProc(goname, lvname string, val int64, meta *ProcMeta) {
	if !strings.HasPrefix(goname, "Proc") {
		return
	}
	goname = goname[4:]
	proc := &Proc{Num: val, Name: goname, OrigName: lvname, ReadStreamIdx: -1, WriteStreamIdx: -1}
	if meta != nil {
		proc.ReadStreamIdx = meta.ReadStream
		proc.WriteStreamIdx = meta.WriteStream
	}
	Gen.Procs = append(Gen.Procs, *proc)
}

// parseNumber makes sure that a parsed numerical value can be parsed to a 64-
// bit integer.
func parseNumber(val string) (int64, error) {
	base := 10
	if strings.HasPrefix(val, "0x") {
		base = 16
		val = val[2:]
	}
	n, err := strconv.ParseInt(val, base, 64)
	return n, err
}

// parseMeta parses procedure metadata to simple string mapping
func parseMeta(meta string) (*ProcMeta, error) {
	res := &ProcMeta{
		ReadStream:  -1,
		WriteStream: -1,
	}
	for _, line := range strings.Split(meta, "\n") {
		atInd := strings.Index(line, "@")
		if atInd == -1 {
			// Should be only first and last line of comment
			continue
		}
		spl := strings.SplitN(line[atInd+1:], ":", 2)
		if len(spl) != 2 {
			return nil, fmt.Errorf("invalid annotation: %s", meta)
		}
		spl[1] = strings.Trim(spl[1], " ")
		switch spl[0] {
		case "readstream":
			var err error
			res.ReadStream, err = strconv.Atoi(spl[1])
			if err != nil {
				return nil, fmt.Errorf("invalid value for readstream: %s", spl[1])
			}
		case "writestream":
			var err error
			res.WriteStream, err = strconv.Atoi(spl[1])
			if err != nil {
				return nil, fmt.Errorf("invalid value for writestream: %s", spl[1])
			}
		}
	}
	return res, nil
}

// StartStruct is called from the parser when a struct definition is found, but
// before the member declarations are processed.
func StartStruct(name string) {
	goname := identifierTransform(name)
	CurrentStruct.push(&Structure{Name: goname, OrigName: name})
}

// AddStruct is called when the parser has finished parsing a struct. It adds
// the now-complete struct definition to the generator's list.
func AddStruct() {
	st := *CurrentStruct.pop()
	Gen.StructMap[st.Name] = len(Gen.Structs)
	Gen.Structs = append(Gen.Structs, st)
}

// StartTypedef is called when the parser finds a typedef.
func StartTypedef() {
	CurrentTypedef = &Typedef{}
}

// StartUnion is called by the parser when it finds a union declaraion.
func StartUnion(name string) {
	name = identifierTransform(name)
	CurrentUnion = &Union{Name: name}
}

// AddUnion is called by the parser when it has finished processing a union
// type. It adds the union to the generator's list and clears the CurrentUnion
// pointer. We handle unions by declaring an interface for the union type, and
// adding methods to each of the cases so that they satisfy the interface.
func AddUnion() {
	Gen.UnionMap[CurrentUnion.Name] = len(Gen.Unions)
	Gen.Unions = append(Gen.Unions, *CurrentUnion)
	CurrentUnion = nil
}

// StartCase is called when the parser finds a case statement within a union.
func StartCase(dvalue string) {
	// In libvirt, the discriminant values are all C pre- processor definitions.
	// Since we don't run the C pre-processor on the protocol file, they're
	// still just names when we get them - we don't actually have their integer
	// values. We'll use the strings to build the type names, although this is
	// brittle, because we're defining a type for each of the case values, and
	// that type needs a name.
	caseName := dvalue
	if ix := strings.LastIndexByte(caseName, '_'); ix != -1 {
		caseName = caseName[ix+1:]
	}
	caseName = fromSnakeToCamel(caseName)
	//dv, ok := lvTypedParams[dvalue]
	//if ok {
	//	dvalue = strconv.FormatUint(uint64(dv), 10)
	//}
	// TODO: search func
	for _, v := range Gen.Consts {
		if v.OrigName == dvalue {
			dvalue = v.Val
			break
		}
	}
	for _, v := range Gen.EnumVals {
		if v.OrigName == dvalue {
			dvalue = v.Val
			break
		}
	}
	// special case for NfsArgop4
	if "NfsArgop4" == CurrentUnion.Name {
		CurrentFunc = &Func{Name: caseName, OpNum: dvalue, Args:"", Op:"TODO", Type:"TODO"}

	}
	CurrentCase = &Case{CaseName: caseName, DiscriminantVal: dvalue}
}

// AddCase is called when the parser finishes parsing a case.
func AddCase() {
	CurrentUnion.Cases = append(CurrentUnion.Cases, *CurrentCase)
	// special case for NfsArgop4
	if nil != CurrentFunc {
		CurrentFunc.Op = CurrentCase.Name
		CurrentFunc.Type = CurrentCase.Type
		ix, ok := Gen.StructMap[CurrentCase.Type]
		if ok {
			a := make([]string, 0)
			i := make([]string, 0)
			for _, v := range Gen.Structs[ix].Members {
				a = append(a, fmt.Sprintf("%s %s", strings.ToLower(v.Name), v.Type))
				i = append(i, fmt.Sprintf("%s:%s", v.Name, strings.ToLower(v.Name)))
			}
			CurrentFunc.Args = strings.Join(a, ", ")
			CurrentFunc.Init = strings.Join(i, ", ")
		}
		Gen.ArgFunc = append(Gen.ArgFunc, *CurrentFunc)
		CurrentFunc = nil
	}
	CurrentCase = nil
}

// AddDeclaration is called by the parser when it find a declaration (int x).
// The declaration will be added to any open container (such as a struct, if the
// parser is working through a struct definition.)
func AddDeclaration(identifier, itype string) {
	addDecl(NewDecl(identifier, itype))
}

// addDecl adds a declaration to the current container.
func addDecl(decl *Decl) {
	if !CurrentStruct.empty() {
		st := CurrentStruct.peek()
		st.Members = append(st.Members, *decl)
	} else if CurrentTypedef != nil {
		CurrentTypedef.Name = decl.Name
		CurrentTypedef.OrigName = decl.OrigName
		CurrentTypedef.Type = decl.Type
		if CurrentTypedef.Name != "string" {
			// Omit recursive typedefs. These happen because we're massaging
			// some of the type names.
			Gen.Typedefs = append(Gen.Typedefs, *CurrentTypedef)
		}
		CurrentTypedef = nil
	} else if CurrentCase != nil {
		CurrentCase.Name = decl.Name
		CurrentCase.Type = decl.Type
	} else if CurrentUnion != nil {
		CurrentUnion.DiscriminantName = decl.Name
		CurrentUnion.DiscriminantType = decl.Type
	}
}

// AddFixedArray is called by the parser to add a fixed-length array to the
// current container (struct, union, etc). Fixed-length arrays are not length-
// prefixed.
func AddFixedArray(identifier, itype, len string) {
	atype := fmt.Sprintf("[%v]%v", len, itype)
	addDecl(NewDecl(identifier, atype))
}

// AddVariableArray is called by the parser to add a variable-length array.
// Variable-length arrays are prefixed with a 32-bit unsigned length, and may
// also have a maximum length specified.
func AddVariableArray(identifier, itype, len string) {
	// This code ignores the length restriction (array<MAXLEN>), so as of now we
	// can't check to make sure that we're not exceeding that restriction when
	// we fill in message buffers. That may not matter, if libvirt's checking is
	// careful enough.
	atype := "[]" + itype
	// Handle strings specially. In the rpcgen definition a string is specified
	// as a variable-length array, either with or without a max length. We want
	// these to be go strings, so we'll just remove the array specifier.
	if itype == "string" {
		atype = itype
	}
	addDecl(NewDecl(identifier, atype))
}

// AddOptValue is called by the parser to add an optional value. These are
// declared in the protocol definition file using a syntax that looks like a
// pointer declaration, but are actually represented by a variable-sized array
// with a maximum size of 1.
func AddOptValue(identifier, itype string) {
	atype := "[]" + itype
	decl := NewDecl(identifier, atype)
	newType := "Opt" + decl.Name
	goEquivTypes[decl.Name] = newType
	decl.Name = newType
	addDecl(decl)
}

// checkIdentifier determines whether an identifier is in our translation list.
// If so it returns the translated name. This is used to massage the type names
// we're emitting.
func checkIdentifier(i string) string {
	nn, reserved := goEquivTypes[i]
	if reserved {
		return nn
	}
	return i
}

// GetUnion returns the type information for a union. If the provided type name
// isn't a union, this will return a zero-value Union type.
func (decl *Decl) GetUnion() Union {
	ix, ok := Gen.UnionMap[decl.Type]
	if ok {
		return Gen.Unions[ix]
	}
	return Union{}
}
