package v40

import (
	"github.com/avekceeb/nfsverificator/rpc"
	"github.com/avekceeb/nfsverificator/xdr"
	"math/rand"
	"errors"
)

type ArgArrayT struct {
	Args []NfsArgop4
}

type CompoundMessage struct {
	Head rpc.Header
	Args COMPOUND4args
}

type V40 struct {
	RpcClient      *rpc.Client
	Auth           rpc.Auth
	ClientId       uint64
	Verifier       Verifier4
	Id             string
	Seq            uint32
	// TODO: callback server ; thread to send RENEW ?
}


func (cli *V40) MockReboot() {
	cli.Seq = 0
	r := rand.Uint64()
	cli.Verifier = Verifier4{
		byte(r&0xff), byte((r&0xff00)>>8),
		byte((r&0xff0000)>>16), byte((r&0xff000000)>>24),
		byte((r&0xff000000)>>32), byte((r&0xff0000000000)>>40),
		byte((r&0xff000000000000)>>48), byte((r&0xff00000000000000)>>56),
	}
}

func (cli *V40) Close() {
	cli.RpcClient.Close()
}

func (cli *V40) GetClientID() (NfsClientID4) {
	// TODO: ???
	return NfsClientID4{
		Verifier: cli.Verifier,
		ID: cli.Id}
}

func (cli *V40) GetCallBack() (CbClient4) {
	// TODO: real client, calculate address
	return CbClient4{
		CbProgram:0x40000000,
		CbLocation: Clientaddr4{RNetid:"tcp", RAddr:"127.0.0.1.139.249"}}
}

func (cli *V40) Compound(args ...NfsArgop4) (reply COMPOUND4res) {
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
	// TODO: log err
	// We are not interested in pushing this error information further
	// so this is the place where all rpc errors stop
	if nil != err {
		return COMPOUND4res{Status:10049}
	}
	if nil == res {
		return COMPOUND4res{Status:10049}
	}
	err = xdr.Read(res, &reply)
	if nil != err {
		return COMPOUND4res{Status:10049}
	}
	return reply
}

func (cli*V40) Null() (error) {
	res, err := cli.RpcClient.Call(rpc.Header{
		Rpcvers: 2,
		Prog:    NFS4_PROGRAM,
		Vers:    NFS_V4,
		Proc:    NFSPROC4_NULL,
		Cred:    cli.Auth,
		Verf:    rpc.AuthNull,
	})
	// TODO:
	// ... though these errors are going to be pushed forward
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


func NewV40(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string) (*V40) {
	client := V40{}
	client.Auth = rpc.NewAuthUnix(authHost, uid, gid).Auth()
	var err error
	client.RpcClient, err = rpc.DialService(srvHost, srvPort)
	if err != nil {
		panic(err.Error())
	}
	client.Id = cid
	client.MockReboot()
	return &client
}
