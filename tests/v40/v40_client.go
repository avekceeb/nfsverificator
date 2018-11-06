package v40tests

import (
    . "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/v40"
	"strings"
	"github.com/avekceeb/nfsverificator/rpc"
	"math/rand"
	"errors"
	"github.com/avekceeb/nfsverificator/xdr"
	"fmt"
)


const (
	Nfs4ClientError = 10099
)

type NFSv40Client struct {
	RpcClient      *rpc.Client
	Auth           rpc.Auth
	AuthSys        AuthsysParms
	ClientId       uint64
	Seq            uint32 // TODO ??
	Verifier       Verifier4
	Id             string
}

type ArgArrayT struct {
	Args []NfsArgop4
}

type CompoundMessage struct {
	Head rpc.Header
	Args COMPOUND4args
}


var (
	BeOK = Equal(int32(NFS4_OK))
)

func NewNFSv40Client(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string) (*NFSv40Client) {
	client := NFSv40Client{}
	u := rpc.NewAuthUnix(authHost, uid, gid)
	client.Auth = u.Auth()
	client.AuthSys = AuthsysParms{Stamp:u.Stamp, Uid:uid, Gid:gid, Machinename:authHost, GidLen:0}
	var err error
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
	// Parse reply at last
	err = xdr.Read(res, &reply)
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


func (t *NFSv40Client) ExpectOK(args ...NfsArgop4) ([]NfsResop4) {
	reply, err := t.Compound(args...)
	Expect(err).To(BeNil())
	Expect(reply.Status).To(BeOK)
	// TODO: ???
	//for _, k := range reply.Resarray {
	//    Expect(GetStatus(&k)).To(Equal(NFS4_OK))
	//}
	return reply.Resarray
}


func (t *NFSv40Client) ExpectErr(stat int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	Expect(err).To(BeNil())
	Expect(res.Status).To(Equal(int32(stat)))
	return res.Resarray
}


func (t *NFSv40Client) GetRootFH() (NfsFh4) {
    ret := t.ExpectOK(Putrootfh(), Getfh())
    return ret[1].Opgetfh.Resok4.Object
}

func (t *NFSv40Client) GetExportFH(export string) (fh NfsFh4) {
    fh = t.GetRootFH()
    for _, k := range strings.Split(export, "/") {
        if "" == k {
            continue
        }
        ret := t.ExpectOK(Putfh(fh), Lookup(k), Getfh())
        fh = ret[2].Opgetfh.Resok4.Object
    }
    return fh
}

func (t *NFSv40Client) GetFHType(fh NfsFh4) ([]byte) {
    ret := t.ExpectOK(Putfh(fh), Getattr([]uint32{FATTR4_FH_EXPIRE_TYPE}))
    return ret[1].Opgetattr.Resok4.ObjAttributes.AttrVals
}

func (t *NFSv40Client) CreateDir(fh NfsFh4, name string, perm uint) (NfsFh4) {
	r := t.ExpectOK(
		Putfh(fh),
		Create(Createtype4{Type:NF4DIR},
			name,
			Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
                   AttrVals: GetPermAttrList(0777)}),
		Getfh())
    return r[2].Opgetfh.Resok4.Object
}

func (t *NFSv40Client) SetAttr(fh NfsFh4, perm uint) {
	t.ExpectOK(Putfh(fh),
		Setattr(Stateid4{}, // TODO: ????
			Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
				AttrVals:GetPermAttrList(perm)}))
}

func (t *NFSv40Client) OpenSimple(fh NfsFh4, name string) (newFH NfsFh4, stateId Stateid4) {
    r := t.ExpectOK(
		Putfh(fh),
        Open(t.Seq,
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
            OpenClaim4{Claim:CLAIM_NULL, File: name}),
		Getfh())
	newFH = r[2].Opgetfh.Resok4.Object
	stateId = r[1].Opopen.Resok4.Stateid
	t.Seq += 1
	r = t.ExpectOK(Putfh(newFH), OpenConfirm(stateId, t.Seq))
	stateId = r[1].OpopenConfirm.Resok4.OpenStateid
	t.Seq += 1
	return newFH, stateId
}

func (t *NFSv40Client) LockSimple(fh NfsFh4, ltype int32, off uint64, length uint64, stateId Stateid4) ([]NfsResop4) {
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
						OpenSeqid: t.Seq,
						OpenStateid: stateId,
						LockSeqid: 0,
						LockOwner:LockOwner4{
							Clientid: t.ClientId,
							Owner: t.Id}}}))
}


func (t *NFSv40Client) BuildLockt(ltype int32, off uint64, length uint64, owner string) (LOCKT4args) {
    return LOCKT4args{
		Locktype:ltype,
		Offset:off,
		Length:length,
		Owner:LockOwner4{Clientid: t.ClientId, Owner: owner}}
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
