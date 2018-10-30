package v40tests

import (
    . "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/v40"
	"strings"
)

type V40Test struct {
	Client *V40
}

func NewV40Test(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string) (*V40Test) {
	tc := V40Test{}
	tc.Client = NewV40(srvHost, srvPort, authHost, uid, gid, cid)
	return &tc
}

func (t *V40Test) ExpectOK(args ...NfsArgop4) (reply COMPOUND4res) {
	reply = t.Client.Compound(args...)
    Expect(reply).ToNot(BeNil())
    Expect(reply.Status).To(Equal(Nfsstat4(NFS4_OK)))
    // TODO: this probably is not needed
    //for _, k := range reply.Resarray {
    //    Expect(GetStatus(&k)).To(Equal(Nfsstat4(NFS4_OK)))
    //}
	return reply
}


func (t *V40Test) ExpectErr(stat int32, args ...NfsArgop4) (res COMPOUND4res) {
	res = t.Client.Compound(args...)
    Expect(res).ToNot(BeNil())
    Expect(res.Status).To(Equal(Nfsstat4(stat)))
	return res
}


func (t *V40Test) GetRootFH() (NfsFh4) {
    ret := t.ExpectOK(Putrootfh(), Getfh())
    return ret.Resarray[1].Opgetfh.Resok4.Object
}

func (t *V40Test) GetExportFH(export string) (fh NfsFh4) {
    fh = t.GetRootFH()
    for _, k := range strings.Split(export, "/") {
        if "" == k {
            continue
        }
        ret := t.ExpectOK(Putfh(fh), Lookup(Component4(k)), Getfh())
        fh = ret.Resarray[2].Opgetfh.Resok4.Object
    }
    return fh
}

func (t *V40Test) GetFHType(fh NfsFh4) ([]byte) {
    ret := t.ExpectOK(Putfh(fh), Getattr(Bitmap4{FATTR4_FH_EXPIRE_TYPE}))
    return ret.Resarray[1].Opgetattr.Resok4.ObjAttributes.AttrVals
}

func (t *V40Test) CreateDir(fh NfsFh4, name string, perm uint) (NfsFh4) {
	r := t.ExpectOK(
		Putfh(fh),
		Create(Createtype4{Type:NF4DIR},
			Component4(name),
			Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
                   AttrVals: GetPermAttrList(0777)}),
		Getfh())
    return r.Resarray[2].Opgetfh.Resok4.Object
}