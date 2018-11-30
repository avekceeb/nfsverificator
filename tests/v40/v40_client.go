package v40tests

import (
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
	"github.com/avekceeb/nfsverificator/rpc"
	"github.com/avekceeb/nfsverificator/xdr"
    "github.com/onsi/ginkgo"
	"strings"
	"math/rand"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
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
	RpcClient      *rpc.Client
	Auth           rpc.Auth
	AuthSys        AuthsysParms
	ClientId       uint64
	Seq            uint32 // TODO ??
	Verifier       Verifier4
	Id             string
	sentNum        uint32
	recvNum        uint32
	LeaseTime      uint32
	server         string
	port           int

}

type ArgArrayT struct {
	Args []NfsArgop4
}

type CompoundMessage struct {
	Head rpc.Header
	Args COMPOUND4args
}

/**********************************************
	Minimalist Assertion infrastructure
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

func (c *NFSv40Client) Reconnect() {
	var err error
	c.RpcClient, err = rpc.DialService(c.server, c.port)
	if err != nil {
		panic(err.Error())
	}
}

func NewNFSv40Client(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string) (*NFSv40Client) {
	client := NFSv40Client{server:srvHost}
	u := rpc.NewAuthUnix(authHost, uid, gid)
	client.Auth = u.Auth()
	client.AuthSys = AuthsysParms{
		Stamp:u.Stamp, Uid:uid, Gid:gid, Machinename:authHost, GidLen:0}
	var err error
	if 0 == srvPort {
		srvPort = 2049
	}
	client.port = srvPort
	client.RpcClient, err = rpc.DialService(srvHost, srvPort)
	if err != nil {
		panic(err.Error())
	}
	client.Id = cid
	client.MockReboot()
	return &client
}

func (cli *NFSv40Client) MockReboot() {
	cli.Seq = 0
	r := rand.Uint64()
	cli.Verifier = Verifier4{
		byte(r&0xff), byte((r&0xff00)>>8),
		byte((r&0xff0000)>>16), byte((r&0xff000000)>>24),
		byte((r&0xff000000)>>32), byte((r&0xff0000000000)>>40),
		byte((r&0xff000000000000)>>48), byte((r&0xff00000000000000)>>56),
	}
}

func (cli *NFSv40Client) Close() {
	cli.RpcClient.Close()
}

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

func (cli *NFSv40Client) Compound(args ...NfsArgop4) (reply COMPOUND4res, err error) {
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
	if nil != err {
		fmt.Printf("%s", err.Error())
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	return reply, nil
}

func (cli *NFSv40Client) Null() (error) {
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
		return errors.New("NFSv4.0 NULL returned non empty")
	}
	return nil
}


func (t *NFSv40Client) Pass(args ...NfsArgop4) ([]NfsResop4) {
	reply, err := t.Compound(args...)
	AssertNoErr(err)
	AssertNfsOK(reply.Status)
	return reply.Resarray
}


func (t *NFSv40Client) Fail(stat int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	AssertNoErr(err)
	AssertStatus(res.Status, stat)
	return res.Resarray
}

func (t *NFSv40Client) FailOneOf(listOfErr []int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	AssertNoErr(err)
	AssertStatusOneOf(res.Status, listOfErr)
	return res.Resarray
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

