package tests

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
	. "github.com/avekceeb/nfsverificator/util"
	"github.com/avekceeb/nfsverificator/nfs40"
	"path/filepath"
	"os"
	"math/rand"
	"time"
	"github.com/avekceeb/nfsverificator/rpc"
	"github.com/avekceeb/nfsverificator/xdr"
	"fmt"
	"flag"
)

var Config TestConfig

func init() {
	var configFile string
	flag.StringVar(&configFile, "config",
		filepath.Join(os.Getenv("GOPATH"),
        "src/github.com/avekceeb/nfsverificator/config.json"), "Config File")
	flag.Parse()
	Config = ReadConfig(configFile)
    rand.Seed(time.Now().UnixNano())
}

const letters = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

func RandString(n int) string {
    var l int64 = int64(len(letters))
    b := make([]byte, n)
    for i := 0; i < n; i++ {
        b[i] = letters[int(rand.Int63n(l))]
    }
    return string(b)
}

func hexClientId(id uint64) {
	// TODO: reverse byte order
	var h, l uint64
	l = 0xffffffff & id
	h = (0xffffffff00000000 & id) >> 32
	fmt.Printf("CLIENT: %x/%x\n", uint32(h), uint32(l))
}

type NFSv40Client struct {
	RpcClient      *rpc.Client
	Auth           rpc.Auth
	ClientId       uint64
	Verifier       nfs40.Verifier4
	Id             string
	Seq            uint32
	// TODO: callback server ; thread to send RENEW ?
}

func (cli *NFSv40Client) Close() {
	cli.RpcClient.Close()
}

func (cli *NFSv40Client) Compound(args ...nfs40.NfsArgOp4) ([]nfs40.NfsResOp4) {
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
	// TODO: increment Seq automatically
	Expect(err).To(BeNil())
	Expect(res).ToNot(BeNil())
	var reply nfs40.COMPOUND4res
	err = xdr.Read(res, &reply)
	Expect(err).To(BeNil())
	Expect(reply.Status).To(Equal(int32(nfs40.NFS4_OK)))
	Expect(len(reply.ResArray)).To(Equal(len(args)))
	for k := range args {
		Expect(reply.ResArray[k].ResOp).To(Equal(args[k].ArgOp))
	}
	return reply.ResArray
}

func (cli* NFSv40Client) Null() {
	res, err := cli.RpcClient.Call(rpc.Header{
		Rpcvers: 2,
		Prog:    nfs40.NFS4_PROGRAM,
		Vers:    nfs40.NFS_V4,
		Proc:    nfs40.NFSPROC4_NULL,
		Cred:    cli.Auth,
		Verf:    rpc.AuthNull,
	})
	Expect(err).To(BeNil())
	Expect(res).ToNot(BeNil())
	var b []byte
	res.Read(b)
	Expect(len(b)).To(Equal(0))
}

func NewNFSv40Client() (client NFSv40Client) {
	client.Auth = rpc.NewAuthUnix(RandString(8)+".fake.net", 0, 0).Auth()
	var err error
	client.RpcClient, err = rpc.DialService(Config.ServerHost, Config.ServerPort)
	//defer client.Close()
	if err != nil {
		panic(err.Error())
	}
	client.Id = RandString(8)
	client.Seq = 0
	cb_cl := nfs40.NfsClientId{
		Verifier:nfs40.Verifier4{
			0x04, 0x05, 0x06, 0x07, 0x08, 0x19, 0x1a, 0x1b},
		Id: client.Id}
	// TODO: real callback server
	cb := nfs40.CallbackClient{
		Program:0x40000000,
		Location: nfs40.ClientAddr{NetId:"tcp", Addr:"127.0.0.1.139.249"}}

	a := nfs40.SETCLIENTID4args{
		Client: cb_cl,
		Callback: cb,
		CallbackIdent: 1}

	ret := client.Compound(nfs40.NfsArgOp4{ArgOp:nfs40.OP_SETCLIENTID, SetClientId:a})

	Expect(ret[0].SetClientId.Status).To(Equal(int32(nfs40.NFS4_OK)))
	client.ClientId = ret[0].SetClientId.ResOk.ClientId
	client.Verifier = ret[0].SetClientId.ResOk.Verifier

	hexClientId(client.ClientId)

	// TODO
	_ = client.Compound(nfs40.GetSetClientConfirm(client.ClientId, client.Verifier))
	//Expect(ret[0].SetClientIdConfirm. TODO: Status).To(Equal())
	return client
}

func TestSanity(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Sanity")
}
