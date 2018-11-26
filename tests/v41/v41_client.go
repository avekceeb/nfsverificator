package v41tests

import (
	"math/rand"
	"errors"
	"fmt"
	"github.com/avekceeb/nfsverificator/rpc"
	"github.com/avekceeb/nfsverificator/xdr"
	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/common"
	"github.com/onsi/ginkgo"
	"strings"
	"github.com/davecgh/go-spew/spew"
)

const (
	Nfs4ClientError = 10099
)

var (
	DefProtect = StateProtect4A{SpaHow:0}
	DefImpl = []NfsImplID4{{
				NiiDate:Nfstime4{Seconds: 0, Nseconds: 0},
				NiiDomain:"kernel.org",
				NiiName:"Linux"}}
	DefExchgFlags = MakeUint32Flags(
		EXCHGID4_FLAG_BIND_PRINC_STATEID,
		EXCHGID4_FLAG_SUPP_MOVED_MIGR,
		EXCHGID4_FLAG_SUPP_MOVED_REFER)
	DefCsFlags = MakeUint32Flags(
		CREATE_SESSION4_FLAG_PERSIST,
		CREATE_SESSION4_FLAG_CONN_BACK_CHAN)
	DefChannelAttrs = ChannelAttrs4{
		CaHeaderpadsize:0,
		CaMaxrequestsize:1049620,
		CaMaxresponsesize:1049480,
		CaMaxresponsesizeCached:3428,
		CaMaxoperations:32,
		CaMaxrequests:64,
	}
)

/**********************************************
	Minimalist Assertion infrastructure
	TODO: duplication of v4 v4.1 v4.2
***********************************************/

func Assert(condition bool, errMessage string) {
	if ! condition {
		ginkgo.Fail(errMessage)
	}
}

func AssertStatus(actual int32, expected int32) {
	Assert(actual == expected,
		fmt.Sprintf("Expected: %s  Got: %s",
			ErrorName(expected), ErrorName(actual)))
}

func AssertStatusOneOf(actual int32, expected []int32) {
	list := []string{}
	for _, err := range expected {
		list = append(list, ErrorName(err))
		if actual == err {
			return
		}
	}
	ginkgo.Fail(fmt.Sprintf("Expected one of: %s  Got: %s",
			strings.Join(list, ", "), ErrorName(actual)))
}

func AssertNfsOK(actual int32) {
	AssertStatus(actual, NFS4_OK)
}

func AssertNoErr(err error) {
	if err != nil {
		ginkgo.Fail(fmt.Sprintf(
			"Unexpected error: %s", err.Error()))
	}
}

//////////////////////////////////

type ArgArrayT struct {
	Args []NfsArgop4
}

type CompoundMessage struct {
	Head rpc.Header
	Args COMPOUND4args
}

type NFSv41Client struct {
	RpcClient      *rpc.Client
	Auth           rpc.Auth
	AuthSys        AuthsysParms
	ClientId       uint64
	Seq            uint32 // TODO ??
	Verifier       Verifier4
	Id             string
	EidFlags       uint32
	// TODO: per server / per session
	// Now only one session
	Sid            Sessionid4
	// Now only fore channel
	ForeChAttr     ChannelAttrs4
	Server         string
	// TODO:
	LeaseTime      uint32
	sentNum        uint32
	recvNum        uint32
	DL             bool
	XT             bool
	LU             bool
	MD             bool
	RD             bool
}

func (cli *NFSv41Client) MockReboot() {
	cli.Seq = 0
	cli.Verifier = Uint64ToVerifier(rand.Uint64())
}

func (cli *NFSv41Client) Close() {
	cli.RpcClient.Close()
}

func (cli *NFSv41Client) Compound(args ...NfsArgop4) (reply COMPOUND4res, err error) {
	/*
		tracing here (not in rpc client) because
		we are not interested im rpc header
		and reply casted to compound res here
	 */
	if (Config.Trace) {
		fmt.Println()
		fmt.Println("#", cli.sentNum, Tm(),
			">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Println()
		spew.Dump(args)
	}

	res, err := cli.RpcClient.Call(CompoundMessage{
		Head: rpc.Header{
			Rpcvers: 2,
			Prog:    NFS4_PROGRAM,
			Vers:    NFS_V4,
			Proc:    NFSPROC4_COMPOUND,
			Cred:    cli.Auth,
			Verf:    rpc.AuthNull,
		},
		Args: COMPOUND4args{
			Tag: "",
			Minorversion: 1,
			Argarray: args,
		},
	})
	if nil != err {
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	if nil == res {
		return COMPOUND4res{Status:Nfs4ClientError}, errors.New("Reply is nil")
	}
	// Parse reply at last
	err = xdr.Read(res, &reply)
	cli.recvNum++
	if (Config.Trace) {
		fmt.Println()
		fmt.Println("#", cli.recvNum, Tm(),
			"<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println()
		spew.Dump(reply)
	}
	cli.sentNum++
	if nil != err {
		fmt.Printf("%s", err.Error())
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	if len(args) > 0 && len(reply.Resarray) > 0 {
		if args[0].Argop == OP_SEQUENCE {
			if reply.Resarray[0].Opsequence.SrStatus == int32(NFS4_OK) {
				cli.Seq++
			}
		}
	}
	return reply, nil
}

func (cli*NFSv41Client) Null() (error) {
	res, err := cli.RpcClient.Call(rpc.Header{
		Rpcvers: 2,
		Prog:    NFS4_PROGRAM,
		Vers:    NFS_V4,
		Proc:    NFSPROC4_NULL,
		Cred:    cli.Auth,
		Verf:    rpc.AuthNull,
	})
	if nil != err {
		return err
	}
	if nil == res {
		return errors.New("rpc returned nil")
	}
	var b []byte
	res.Read(b)
	if len(b) != 0 {
		return errors.New("NFSv4.1 NULL returned non empty")
	}
	return nil
}


func NewNFSv41Client(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string) (*NFSv41Client) {
	client := NFSv41Client{}
	u := rpc.NewAuthUnix(authHost, uid, gid)
	client.Auth = u.Auth()
	client.AuthSys = AuthsysParms{
		Stamp:u.Stamp, Uid:uid, Gid:gid, Machinename:authHost, GidLen:0}
	var err error
	if 0 == srvPort {
		srvPort = 2049
	}
	client.RpcClient, err = rpc.DialService(srvHost, srvPort)
	if err != nil {
		panic(err.Error())
	}
	client.Id = cid
	client.MockReboot()
	client.Server = srvHost
	return &client
}

func NewNFSv41DefaultClient() (*NFSv41Client) {
	return NewNFSv41Client(Config.GetHost(), Config.GetPort(),
		RandString(8) + ".fake.net", 0, 0, RandString(8))
}

// HELPERS FOR CHECK COMPOUND STATUS //

func (t *NFSv41Client) Pass(args ...NfsArgop4) ([]NfsResop4) {
	reply, err := t.Compound(args...)
	AssertNoErr(err)
	AssertNfsOK(reply.Status)
	// Note: not checking every op in compound, only overall status
	return reply.Resarray
}


func (t *NFSv41Client) Fail(stat int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	AssertNoErr(err)
	AssertStatus(res.Status, stat)
	return res.Resarray
}


func (t *NFSv41Client) FailOneOf(listOfErr []int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	AssertNoErr(err)
	AssertStatusOneOf(res.Status, listOfErr)
	return res.Resarray
}

// HELPERS TO EXECUTE REQUESTS //

/// TODO: this could not handle unexported 'holes' in path
func (t *NFSv41Client) LookupFromRoot(path string) (fh NfsFh4) {
	ret := t.Pass(
		Sequence(t.Sid, t.Seq, 0, 0, false), Putrootfh(), Getfh())
	fh = LastRes(&ret).Opgetfh.Resok4.Object
	for _, k := range strings.Split(path, "/") {
		if "" == k {
			continue
		}
		ret = t.Pass(
			Sequence(t.Sid, t.Seq, 0, 0, false), Putfh(fh), Lookup(k), Getfh())
		fh = LastRes(&ret).Opgetfh.Resok4.Object
	}
	return fh
}

func (cli *NFSv41Client) ExchangeId() {
	r := cli.Pass(ExchangeId(
		ClientOwner4{
			CoOwnerid: RandString(14),
			CoVerifier: Verifier4{}},
		DefExchgFlags,
		DefProtect,
		DefImpl))
	ei := LastRes(&r).OpexchangeID.EirResok4
	cli.ClientId = ei.EirClientid
	cli.Seq = ei.EirSequenceid
	cli.EidFlags = ei.EirFlags
}

func (cli *NFSv41Client) CreateSession() {
	/*
As previously stated, CREATE_SESSION can be sent with or without a
preceding SEQUENCE operation.  Even if a SEQUENCE precedes
CREATE_SESSION, the server MUST maintain the CREATE_SESSION reply
cache, which is separate from the reply cache for the session
associated with a SEQUENCE.  If CREATE_SESSION was originally sent by
itself, the client MAY send a retry of the CREATE_SESSION operation
within a COMPOUND preceded by a SEQUENCE.  If CREATE_SESSION was
originally sent in a COMPOUND that started with a SEQUENCE, then the
client SHOULD send a retry in a COMPOUND that starts with a SEQUENCE
that has the same session ID as the SEQUENCE of the original request.
However, the client MAY send a retry in a COMPOUND that either has no
preceding SEQUENCE, or has a preceding SEQUENCE that refers to a
different session than the original CREATE_SESSION.  This might be
necessary if the client sends a CREATE_SESSION in a COMPOUND preceded
by a SEQUENCE with session ID X, and session X no longer exists.
Regardless, any retry of CREATE_SESSION, with or without a preceding
SEQUENCE, MUST use the same value of csa_sequence as the original.
*/
	s := cli.Pass(CreateSession(
		cli.ClientId,
		cli.Seq,
		DefCsFlags,
		DefChannelAttrs,
		DefChannelAttrs,
		0x40000000, // CallBack Program
		[]CallbackSecParms4{{
			CbSecflavor:1,
			CbspSysCred:cli.AuthSys}}))
	resok := s[0].OpcreateSession.CsrResok4
	cli.Sid = resok.CsrSessionid
	// TODO: now only fore channel
	cli.ForeChAttr = resok.CsrForeChanAttrs

/*
   Once the session is created, the first SEQUENCE or CB_SEQUENCE
   received on a slot MUST have a sequence ID equal to 1; if not, the
   replier MUST return NFS4ERR_SEQ_MISORDERED.

   BTW, Linux disregards this

 */
	cli.Seq = 1
	cli.Pass(
		Sequence(cli.Sid, cli.Seq, 0, 0, false),
		ReclaimComplete(false))
}

func (cli *NFSv41Client) GetSomeAttr() {
	cli.Pass(
		Sequence(cli.Sid, cli.Seq, 0, 0, false),
		Putrootfh(),
		SecinfoNoName(0))
	// TODO
	l := cli.Pass(
		Sequence(cli.Sid, cli.Seq, 0, 0, false),
		Putrootfh(),
		Getfh(),
		Getattr(MakeGetAttrFlags(FATTR4_LEASE_TIME)))
	cli.LeaseTime = BytesToUint32(LastRes(&l).Opgetattr.Resok4.ObjAttributes.AttrVals)
	r := cli.Pass(
		Sequence(cli.Sid, cli.Seq, 0, 0, false),
		Putrootfh(),
		Access(MakeUint32Flags(
			ACCESS4_DELETE, ACCESS4_EXTEND, ACCESS4_LOOKUP,
			ACCESS4_MODIFY,	ACCESS4_READ, ACCESS4_EXECUTE)),
	)
	access := LastRes(&r).Opaccess.Resok4.Access
	cli.DL = CheckFlag(access, ACCESS4_DELETE)
	cli.XT = CheckFlag(access, ACCESS4_EXTEND)
	cli.LU = CheckFlag(access, ACCESS4_LOOKUP)
	cli.MD = CheckFlag(access, ACCESS4_MODIFY)
	cli.RD = CheckFlag(access, ACCESS4_READ)
	// TODO : execute ???

}

// HELPERS FOR BUILDING NfsArgop4 FOR COMPOUNDS //

func (t *NFSv41Client) SequenceArgs() NfsArgop4 {
	return Sequence(t.Sid, t.Seq, 0, 0, false)
}

/*
OPEN4resok::Rflags

o  OPEN4_RESULT_CONFIRM is deprecated and MUST NOT be returned by an
  NFSv4.1 server.

o  OPEN4_RESULT_LOCKTYPE_POSIX indicates that the server's byte-range
  locking behavior supports the complete set of POSIX locking
  techniques [24].  From this, the client can choose to manage byte-
  range locking state in a way to handle a mismatch of byte-range
  locking management.

o  OPEN4_RESULT_PRESERVE_UNLINKED indicates that the server will
  preserve the open file if the client (or any other client) removes
  the file as long as it is open.  Furthermore, the server promises
  to preserve the file through the grace period after server
  restart, thereby giving the client the opportunity to reclaim its
  open.

o  OPEN4_RESULT_MAY_NOTIFY_LOCK indicates that the server may attempt
  CB_NOTIFY_LOCK callbacks for locks on this file.  This flag is a
  hint only, and may be safely ignored by the client.
 */
func (t *NFSv41Client) OpenArgs() NfsArgop4 {
	return Open(t.Seq,
			OPEN4_SHARE_ACCESS_WRITE,
			OPEN4_SHARE_DENY_NONE,
			OpenOwner4{
				Clientid: t.ClientId,
				Owner: t.Id},
			Openflag4{
				Opentype: OPEN4_CREATE,
				How: Createhow4{
					Mode: UNCHECKED4,
					CreateattrsUnchecked: Fattr4{
						Attrmask: GetBitmap(FATTR4_MODE),
						AttrVals: GetPermAttrList(0644)},
				},
			},
			OpenClaim4{Claim:CLAIM_NULL, File: RandString(12)})
}

func (t *NFSv41Client) OpenNoCreateArgs() NfsArgop4 {
	return Open(t.Seq,
			OPEN4_SHARE_ACCESS_READ,
			OPEN4_SHARE_DENY_NONE,
			OpenOwner4{
				Clientid: t.ClientId,
				Owner: t.Id},
			Openflag4{Opentype: OPEN4_NOCREATE},
			OpenClaim4{Claim:CLAIM_NULL, File: RandString(12)})
}

func (t *NFSv41Client) LockArgs(stateId Stateid4) NfsArgop4 {
		return Lock(
				WRITE_LT,
				false, /*reclaim*/
				0, /*offset*/
				0xffffffffffffffff, /*length*/
				Locker4{
					NewLockOwner:1,
					OpenOwner: OpenToLockOwner4{
						OpenSeqid: t.Seq,
						OpenStateid: stateId,
						LockSeqid: 0,
						LockOwner:LockOwner4{
							Clientid: t.ClientId,
							Owner: t.Id}}})
}

func (t *NFSv41Client) LocktArgs(owner string) (NfsArgop4) {
	return Lockt(
		WRITE_LT, 0, 0xffffffffffffffff, LockOwner4{
			Clientid: t.ClientId, Owner: owner})
}

/*
   Any legal value for locktype
   has no effect on the success or failure of the LOCKU operation.

   The seqid parameter MAY be any value and the server MUST ignore it.

 */
func (t *NFSv41Client) LockuArgs(sid Stateid4) (NfsArgop4) {
	return Locku(WRITE_LT, t.Seq, sid, 0, 0xffffffffffffffff)
}

func (t *NFSv41Client) CreateArgs(name string) (NfsArgop4) {
	return Create(Createtype4{Type:NF4DIR},	name,
			Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
					AttrVals: GetPermAttrList(0777)})
}
