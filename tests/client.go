package tests

import (
	"github.com/avekceeb/nfsverificator/nfs40"
	"github.com/avekceeb/nfsverificator/rpc"
	"math/rand"
	"errors"
	"github.com/avekceeb/nfsverificator/xdr"
)

type NFSv40Client struct {
	RpcClient      *rpc.Client
	Auth           rpc.Auth
	ClientId       uint64
	Verifier       nfs40.Verifier4
	Id             string
	Seq            uint32
	// TODO: callback server ; thread to send RENEW ?
}


func (cli *NFSv40Client) MockReboot() {
	cli.Seq = 0
	r := rand.Uint64()
	cli.Verifier = nfs40.Verifier4{
		byte(r&0xff), byte((r&0xff00)>>8),
		byte((r&0xff0000)>>16), byte((r&0xff000000)>>24),
		byte((r&0xff000000)>>32), byte((r&0xff0000000000)>>40),
		byte((r&0xff000000000000)>>48), byte((r&0xff00000000000000)>>56),
	}
}

func (cli *NFSv40Client) Close() {
	cli.RpcClient.Close()
}

func (cli *NFSv40Client) Compound(args ...nfs40.NfsArgOp4) (reply nfs40.COMPOUND4res) {
	res, err := cli.RpcClient.Call(nfs40.CompoundMessage{
		Head: rpc.Header{
			Rpcvers: 2,
			Prog:    nfs40.NFS4_PROGRAM,
			Vers:    nfs40.NFS_V4,
			Proc:    nfs40.NFSPROC4_COMPOUND,
			Cred:    cli.Auth,
			Verf:    rpc.AuthNull,
		},
		Args: nfs40.COMPOUND4args{
			Tag: "",
			MinorVersion: 0,
			ArgArray: nfs40.ArgArrayT{Args:args},
		},
	})
	// TODO: log err
	// We are not interested in pushing this error information further
	// so this is the place where all rpc errors stop
	if nil != err {
		return nfs40.COMPOUND4res{Status:10049}
	}
	if nil == res {
		return nfs40.COMPOUND4res{Status:10049}
	}
	err = xdr.Read(res, &reply)
	if nil != err {
		return nfs40.COMPOUND4res{Status:10049}
	}
	return reply
}

func (cli* NFSv40Client) Null() (error) {
	res, err := cli.RpcClient.Call(rpc.Header{
		Rpcvers: 2,
		Prog:    nfs40.NFS4_PROGRAM,
		Vers:    nfs40.NFS_V4,
		Proc:    nfs40.NFSPROC4_NULL,
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


func NewNFSv40Client(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string) (client NFSv40Client) {
	client.Auth = rpc.NewAuthUnix(authHost, uid, gid).Auth()
	var err error
	client.RpcClient, err = rpc.DialService(srvHost, srvPort)
	if err != nil {
		panic(err.Error())
	}
	client.Id = cid
	client.MockReboot()
	return client
}

