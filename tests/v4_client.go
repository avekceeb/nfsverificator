package tests

import (
	"github.com/avekceeb/nfsverificator/v4"
	"github.com/avekceeb/nfsverificator/rpc"
	"errors"
	"github.com/davecgh/go-spew/spew"
	"fmt"
	"time"
)

type Nfs4Client struct {
	RpcClient      *rpc.Client
	Auth           rpc.Auth
	AuthSys        v4.AuthsysParms
	ClientId       uint64
	Seq            uint32 // TODO ??
	Id             string
	SentNum        uint32
	RecvNum        uint32
	LeaseTime      uint32
	Server         string
	Port           int
	DoTrace        bool
}

func (c *Nfs4Client) Reconnect() {
	var err error
	c.RpcClient, err = rpc.DialService(c.Server, c.Port)
	if err != nil {
		panic(err.Error())
	}
}

func (cli *Nfs4Client) Close() {
	cli.RpcClient.Close()
}


func NewNfs4Client(srvHost string, srvPort int,
	authHost string,
	uid uint32, gid uint32, cid string,
	trace bool) (*Nfs4Client) {
	client := Nfs4Client{Server:srvHost, DoTrace:trace}
	u := rpc.NewAuthUnix(authHost, uid, gid)
	client.Auth = u.Auth()
	client.AuthSys = v4.AuthsysParms{
		Stamp:u.Stamp, Uid:uid, Gid:gid, Machinename:authHost, GidLen:0}
	var err error
	if 0 == srvPort {
		srvPort = 2049
	}
	client.Port = srvPort
	client.RpcClient, err = rpc.DialService(srvHost, srvPort)
	if err != nil {
		panic(err.Error())
	}
	client.Id = cid
	//client.MockReboot()
	return &client
}

func (cli *Nfs4Client) Null() (error) {
	null := rpc.Header{
		Rpcvers: 2,
		Prog:    100003 /*NFS4_PROGRAM*/,
		Vers:    4 /*NFS_V4*/,
		Proc:    0 /*NFSPROC4_NULL*/,
		Cred:    cli.Auth,
		Verf:    rpc.AuthNull,
	}
	cli.Trace(null)
	res, err := cli.RpcClient.Call(null)
	if nil != err {
		return err
	}
	if nil == res {
		return errors.New("rpc returned nil")
	}
	cli.Trace(res)
	var b []byte
	res.Read(b)
	if len(b) != 0 {
		return errors.New("NFSv4.0 NULL returned non empty")
	}
	return nil
}

func (cli *Nfs4Client) Trace(args... interface{}) {
	if cli.DoTrace {
		t := time.Now()
		fmt.Println()
		fmt.Println("#", cli.SentNum,
			fmt.Sprintf("[%02d:%02d:%02d]",
				t.Hour(), t.Minute(), t.Second()),
			"-----------------------------------")
		fmt.Println()
		spew.Dump(args)
	}
}