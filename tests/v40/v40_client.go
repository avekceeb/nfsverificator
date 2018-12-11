package v40tests

import (
	. "github.com/avekceeb/nfsverificator/tests"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
	"github.com/avekceeb/nfsverificator/rpc"
	"github.com/avekceeb/nfsverificator/xdr"
	"strings"
	"math/rand"
	"errors"
	"fmt"
)


const (
	Nfs4ClientError = 10099
)

var (
	opsIncSeq     []uint32
	anonStateId   Stateid4
	bypassStateId Stateid4
	timeOnce      Nfstime4
	fattrSize     Fattr4
	fattrMTime    Fattr4
	Assert40      Assertion
)

//func gobEncode(e interface{}) ([]byte) {
//	var buf bytes.Buffer
//	enc := gob.NewEncoder(&buf)
//	err := enc.Encode(e)
//	if nil != err {
//		panic(err.Error())
//	}
//	return buf.Bytes()
//}

func init() {
	Assert40 = Assertion{ErrorName:ErrorNameNfs40}
	opsIncSeq = []uint32{OP_CLOSE, OP_LOCK, OP_LOCKU, OP_OPEN, OP_OPEN_CONFIRM}
	bypassStateId.Seqid = NFS4_UINT32_MAX
	for i:= range bypassStateId.Other {
		bypassStateId.Other[i] = 0xff
	}
	timeOnce.Seconds = 1540057110
	timeOnce.Nseconds = 858740700
	fattrSize.Attrmask = GetBitmap(FATTR4_SIZE)
	fattrSize.AttrVals = []byte{0,0,0,0,0,0,0,128}
	fattrMTime.Attrmask = GetBitmap(FATTR4_TIME_MODIFY)
	fattrMTime.AttrVals = []byte{0,0,0,0,0x5b,0xcb,0x68,0x16/*sec*/,
								0x33,0x2c,0x57,0xdc/*nsec*/}
}

type NFSv40Client struct {
	*Nfs4Client
	Verifier       Verifier4
}

type ArgArrayT struct {
	Args []NfsArgop4
}

type CompoundMessage struct {
	Head rpc.Header
	Args COMPOUND4args
}

func opsNames(args ...NfsArgop4) string {
	s := []string{}
	for i := range args {
		s = append(s, OpNameNfs40(args[i].Argop))
	}
	return strings.Join(s, "|")
}

func DefaultClient40() (*NFSv40Client) {
	return NewNFSv40Client(Config.Server, Config.Port,
		RandString(8) + ".fake.net", 0, 0, RandString(8), Config.Trace)
}

func NewNFSv40Client(srvHost string, srvPort int,
	authHost string, uid uint32, gid uint32, cid string, trace bool) (*NFSv40Client) {
	r := rand.Uint64()
	client := NFSv40Client{NewNfs4Client(
		srvHost, srvPort, authHost, uid, gid, cid, trace), Verifier4{
			byte(r&0xff), byte((r&0xff00)>>8),
			byte((r&0xff0000)>>16), byte((r&0xff000000)>>24),
			byte((r&0xff000000)>>32), byte((r&0xff0000000000)>>40),
			byte((r&0xff000000000000)>>48), byte((r&0xff00000000000000)>>56),
	}}
	return &client
}

func (cli *NFSv40Client) Close() {
	cli.RpcClient.Close()
}

// all these functions are not so minor version specific
// and could be moved to common v4 client
// if only I find a way to properly wrap NfsArgop4

func (cli *NFSv40Client) Compound(args ...NfsArgop4) (reply COMPOUND4res, err error) {
	fmt.Println()
	fmt.Println(opsNames(args...))
	cli.Trace(args)
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
			Minorversion: 0,
			Argarray: args,
		},
	})
	if nil != err {
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	if nil == res {
		return COMPOUND4res{Status:Nfs4ClientError}, errors.New("Reply is nil")
	}
	/*
	RFC 7530 9.1.7.  Sequencing of Lock Requests
	The client MUST advance the sequence number for the
	CLOSE, LOCK, LOCKU, OPEN, OPEN_CONFIRM, and OPEN_DOWNGRADE operations.
	This is true even in the event that the previous operation that used the
	sequence number received an error.  The only exception to this rule
	is if the previous operation received one of the following errors:
	NFS4ERR_STALE_CLIENTID, NFS4ERR_STALE_STATEID, NFS4ERR_BAD_STATEID,
	NFS4ERR_BAD_SEQID, NFS4ERR_BADXDR, NFS4ERR_RESOURCE,
	NFS4ERR_NOFILEHANDLE, or NFS4ERR_MOVED.
	*/
	if len(args) > 0 {
		for _, a := range args {
			if InSliceUint32(uint32(a.Argop), opsIncSeq) {
				// TODO: check result
				cli.Seq++
			}
		}
	}
	cli.SentNum++
	// Parse reply at last
	err = xdr.Read(res, &reply)
	cli.RecvNum++
	cli.Trace(reply)
	if nil != err {
		fmt.Printf("%s", err.Error())
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	if nil != err {
		fmt.Printf("%s", err.Error())
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	return reply, nil
}

func (t *NFSv40Client) Pass(args ...NfsArgop4) ([]NfsResop4) {
	reply, err := t.Compound(args...)
	Assert40.AssertNoErr(err)
	Assert40.AssertNfsOK(reply.Status)
	return reply.Resarray
}


func (t *NFSv40Client) Fail(stat int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	Assert40.AssertNoErr(err)
	Assert40.AssertStatus(res.Status, stat)
	return res.Resarray
}

func (t *NFSv40Client) FailOneOf(listOfErr []int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	Assert40.AssertNoErr(err)
	Assert40.AssertStatusOneOf(res.Status, listOfErr)
	return res.Resarray
}

/////// these are minor version (v4.0) specific /////////////////////////////

func (c *NFSv40Client) SetAndConfirmClientId() {
	r := c.Pass(Setclientid(c.GetClientID(), c.GetCallBack(), 1))
	c.ClientId = r[0].Opsetclientid.Resok4.Clientid
	c.Verifier = r[0].Opsetclientid.Resok4.SetclientidConfirm
	c.Pass(SetclientidConfirm(c.ClientId, c.Verifier))
}

func (cli *NFSv40Client) GetClientID() (NfsClientID4) {
	// TODO: ???
	return NfsClientID4{
		Verifier: cli.Verifier,
		ID: cli.Id}
}

func (cli *NFSv40Client) GetCallBack() (CbClient4) {
	// TODO: real client, calculate address
	return CbClient4{
		CbProgram:0x40000000,
		CbLocation: Clientaddr4{RNetid:"tcp", RAddr:"127.0.0.1.139.249"}}
}

func (t *NFSv40Client) GetRootFH() (NfsFh4) {
	ret := t.Pass(Putrootfh(), Getfh())
	return ret[1].Opgetfh.Resok4.Object
}

func (t *NFSv40Client) GetExportFH(export string) (fh NfsFh4) {
	fh = t.GetRootFH()
	for _, k := range strings.Split(export, "/") {
		if "" == k {
			continue
		}
		ret := t.Pass(Putfh(fh), Lookup(k), Getfh())
		fh = ret[2].Opgetfh.Resok4.Object
	}
	return fh
}

func (t *NFSv40Client) GetFHType(fh NfsFh4) ([]byte) {
    ret := t.Pass(Putfh(fh), Getattr([]uint32{FATTR4_FH_EXPIRE_TYPE}))
    return ret[1].Opgetattr.Resok4.ObjAttributes.AttrVals
}

func (t *NFSv40Client) CreateArgs() (NfsArgop4) {
	return Create(Createtype4{Type:NF4DIR},
				RandString(12),
				Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
					AttrVals: GetPermAttrList(0777)})
}

func (cli *NFSv40Client) GetSomeAttr() {
	l := cli.Pass(
		Putrootfh(),
		Getfh(),
		Getattr(MakeGetAttrFlags(FATTR4_LEASE_TIME)))
	cli.LeaseTime = BytesToUint32(LastRes(&l).Opgetattr.Resok4.ObjAttributes.AttrVals)
}


func (t *NFSv40Client) OpenArgs() (NfsArgop4) {
	return Open(t.Seq,
		OPEN4_SHARE_ACCESS_WRITE,
		OPEN4_SHARE_DENY_NONE,
		OpenOwner4{
			Clientid: t.ClientId,
			Owner: t.Id},
		Openflag4{
			Opentype:OPEN4_CREATE,
			How: Createhow4{
				Mode: UNCHECKED4,
				CreateattrsUnchecked: Fattr4{
					Attrmask: GetBitmap(FATTR4_MODE),
					AttrVals: GetPermAttrList(0644)},
			},
		},
		OpenClaim4{Claim:CLAIM_NULL, File: RandString(8)})
}

func (t *NFSv40Client) OpenNoCreateArgs() NfsArgop4 {
	return Open(t.Seq,
			OPEN4_SHARE_ACCESS_READ,
			OPEN4_SHARE_DENY_NONE,
			OpenOwner4{
				Clientid: t.ClientId,
				Owner: t.Id},
			Openflag4{Opentype: OPEN4_NOCREATE},
			OpenClaim4{Claim:CLAIM_NULL, File: RandString(12)})
}

func (t *NFSv40Client) LockArgs(stateId Stateid4) NfsArgop4 {
		return Lock(
				WRITE_LT,
				false, /*reclaim*/
				0, /*offset*/
				NFS4_UINT64_MAX, /*length = whole file*/
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

func (t *NFSv40Client) LocktArgs(owner string) (NfsArgop4) {
	return Lockt(
		WRITE_LT, 0, NFS4_UINT64_MAX, LockOwner4{
			Clientid: t.ClientId, Owner: owner})
}

func (t *NFSv40Client) LockuArgs(sid Stateid4) (NfsArgop4) {
	return Locku(WRITE_LT, t.Seq, sid, 0, NFS4_UINT64_MAX)
}

////// macro commands ////////////////////////////////////

// Note: supposing r = PUTFH;OPEN;GETFH
func (c *NFSv40Client) OpenConfirmMacro(r *([]NfsResop4)) (Stateid4) {
	fh := GrabFh(r)
	stateId := (*r)[1].Opopen.Resok4.Stateid
	if CheckFlag((*r)[1].Opopen.Resok4.Rflags,
		OPEN4_RESULT_CONFIRM) {
		rc := c.Pass(Putfh(fh), OpenConfirm(stateId, c.Seq))
		stateId = rc[1].OpopenConfirm.Resok4.OpenStateid
	}
	return stateId
}

