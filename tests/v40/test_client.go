package v40tests

import (
    . "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/v40"
	"strings"
)

type V40Test struct {
	Client *V40
	Uid uint32 // TODO: ???
	Gid uint32
}

func NewV40Test(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string) (*V40Test) {
	tc := V40Test{Uid:uid, Gid:gid}
	tc.Client = NewV40(srvHost, srvPort, authHost, uid, gid, cid)
	return &tc
}

func (t *V40Test) ExpectOK(args ...NfsArgop4) (reply COMPOUND4res) {
	reply = t.Client.Compound(args...)
    Expect(reply).ToNot(BeNil())
    Expect(reply.Status).To(Equal(NFS4_OK))
    // TODO: this probably is not needed
    //for _, k := range reply.Resarray {
    //    Expect(GetStatus(&k)).To(Equal(NFS4_OK))
    //}
	return reply
}


func (t *V40Test) ExpectErr(stat int32, args ...NfsArgop4) (res COMPOUND4res) {
	res = t.Client.Compound(args...)
    Expect(res).ToNot(BeNil())
    Expect(res.Status).To(Equal(stat))
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
        ret := t.ExpectOK(Putfh(fh), Lookup(k), Getfh())
        fh = ret.Resarray[2].Opgetfh.Resok4.Object
    }
    return fh
}

func (t *V40Test) GetFHType(fh NfsFh4) ([]byte) {
    ret := t.ExpectOK(Putfh(fh), Getattr([]uint32{FATTR4_FH_EXPIRE_TYPE}))
    return ret.Resarray[1].Opgetattr.Resok4.ObjAttributes.AttrVals
}

func (t *V40Test) CreateDir(fh NfsFh4, name string, perm uint) (NfsFh4) {
	r := t.ExpectOK(
		Putfh(fh),
		Create(Createtype4{Type:NF4DIR},
			name,
			Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
                   AttrVals: GetPermAttrList(0777)}),
		Getfh())
    return r.Resarray[2].Opgetfh.Resok4.Object
}

func (t *V40Test) SetAttr(fh NfsFh4, perm uint) {
	t.ExpectOK(Putfh(fh),
		Setattr(Stateid4{}, // TODO: ????
			Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
				AttrVals:GetPermAttrList(perm)}))
}

func (t *V40Test) OpenSimple(fh NfsFh4, name string) (newFH NfsFh4, stateId Stateid4) {
    r := t.ExpectOK(
		Putfh(fh),
        Open(t.Client.Seq,
			OPEN4_SHARE_ACCESS_WRITE,
			OPEN4_SHARE_DENY_NONE,
            OpenOwner4{
				Clientid: t.Client.ClientId,
				Owner: t.Client.Id},
            Openflag4{
				Opentype:OPEN4_CREATE,
                How: Createhow4{
					Mode: UNCHECKED4,
                    CreateattrsUnchecked: Fattr4{
						Attrmask: GetBitmap(FATTR4_MODE),
						AttrVals: GetPermAttrList(0644)},
                },
			},
            OpenClaim4{Claim:CLAIM_NULL, File: name}),
		Getfh())
	newFH = r.Resarray[2].Opgetfh.Resok4.Object
	stateId = r.Resarray[1].Opopen.Resok4.Stateid
	t.Client.Seq += 1
	r = t.ExpectOK(Putfh(newFH), OpenConfirm(stateId, t.Client.Seq))
	stateId = r.Resarray[1].Opopen.Resok4.Stateid
	return newFH, stateId
}

func (t *V40Test) LockSimple(fh NfsFh4, ltype int32, off uint32, length uint64, stateId Stateid4) (COMPOUND4res) {
		return t.ExpectOK(
			Putfh(fh),
			Lock(
				ltype,
				false,
				off,
				length,
				Locker4{
					NewLockOwner:true,
					OpenOwner: OpenToLockOwner4{
						OpenSeqid: t.Client.Seq,
						OpenStateid: stateId,
						LockSeqid:0,
						LockOwner:LockOwner4{
							Clientid: t.Client.ClientId,
							Owner: t.Client.Id}}}))
}


func (t *V40Test) BuildLockt(ltype int32, off uint32, length uint64, owner string) (LOCKT4args) {
    return LOCKT4args{
		Locktype:ltype,
		Offset:off,
		Length:length,
		Owner:LockOwner4{Clientid: t.Client.ClientId, Owner: owner}}
}

//func Write(stateId StateId4, data *[]byte, offset uint64) (NfsArgOp4) {
//    return NfsArgOp4{ArgOp: OP_WRITE,
//        Write: WRITE4args{State:stateId, Offset: offset, Stable: FILE_SYNC4, Data: *data}}
//}

//func Close(seq uint32, stateId StateId4) (NfsArgOp4) {
//    return NfsArgOp4{ArgOp: OP_CLOSE, Close:CLOSE4args{SeqId:seq, StateId: stateId}}
//}
//
//
