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
	"strings"
)

var (
	Config TestConfig
	Uid uint32 = 0
	Gid uint32 = 0
)

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
	fmt.Printf("### New Client: %x/%x\n", uint32(h), uint32(l))
}

func ExpectOK(res nfs40.COMPOUND4res) {
	Expect(res).ToNot(BeNil())
	Expect(res.Status).To(Equal(int32(nfs40.NFS4_OK)))
	// TODO: this probably is not needed
	for _, k := range res.ResArray {
		Expect(nfs40.GetResStatus(&k)).To(Equal(int32(nfs40.NFS4_OK)))
	}
}

func ExpectErr(res nfs40.COMPOUND4res, stat uint32) {
	Expect(res).ToNot(BeNil())
	Expect(res.Status).To(Equal(int32(stat)))
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
	// TODO: increment Seq automatically ?
	/*
	   The client MUST monotonically increment the sequence number for the
	   CLOSE, LOCK, LOCKU, OPEN, OPEN_CONFIRM, and OPEN_DOWNGRADE
	   operations.  This is true even in the event that the previous
	   operation that used the sequence number received an error.  The only
	   exception to this rule is if the previous operation received one of
	   the following errors: NFS4ERR_STALE_CLIENTID, NFS4ERR_STALE_STATEID,
	   NFS4ERR_BAD_STATEID, NFS4ERR_BAD_SEQID, NFS4ERR_BADXDR,
	   NFS4ERR_RESOURCE, NFS4ERR_NOFILEHANDLE
	 */
	Expect(err).To(BeNil())
	Expect(res).ToNot(BeNil())
	err = xdr.Read(res, &reply)
	Expect(err).To(BeNil())
	return reply
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

func (cli* NFSv40Client) getRootFH() (nfs40.FH4) {
	ret := cli.Compound(nfs40.PutRootFH(), nfs40.GetFH())
	Expect(ret.ResArray[1].GetFH.Status).To(Equal(int32(nfs40.NFS4_OK)))
	return ret.ResArray[1].GetFH.FH
}

func (cli *NFSv40Client) getExportFH(export string) (fh nfs40.FH4) {
	fh = cli.getRootFH()
	for _, k := range strings.Split(export, "/") {
		if "" == k {
			continue
		}
		ret := cli.Compound(nfs40.PutFH(fh), nfs40.Lookup(k), nfs40.GetFH())
		fh = ret.ResArray[2].GetFH.FH
	}
	return fh
}

func (cli *NFSv40Client) getFHType(fh nfs40.FH4) ([]byte) {
	ret := cli.Compound(nfs40.PutFH(fh),
		nfs40.GetAttr(nfs40.FATTR4_FH_EXPIRE_TYPE))
	return ret.ResArray[1].GetAttr.Attr.AttrList
}

func createNFSv40Client(uid uint32, gid uint32, cid string) (client NFSv40Client) {
	client.Auth = rpc.NewAuthUnix(RandString(8)+".fake.net", uid, gid).Auth()
	var err error
	client.RpcClient, err = rpc.DialService(Config.ServerHost, Config.ServerPort)
	if err != nil {
		panic(err.Error())
	}
	client.Id = cid
	client.MockReboot()
	return client
}

func NewNFSv40Client() (client NFSv40Client) {
	client = createNFSv40Client(Uid, Gid, RandString(8))
	// TODO: uid gid
	//Uid += 1
	//Gid += 1
	cb_cl := nfs40.NfsClientId{
		Verifier: client.Verifier,
		Id: client.Id}
	// TODO: real callback server
	cb := nfs40.CallbackClient{
		Program:0x40000000,
		Location: nfs40.ClientAddr{NetId:"tcp", Addr:"127.0.0.1.139.249"}}
	setClientArgs := nfs40.SETCLIENTID4args{
		Client: cb_cl,
		Callback: cb,
		CallbackIdent: 1}
	ret := client.Compound(nfs40.NfsArgOp4{
		ArgOp:nfs40.OP_SETCLIENTID, SetClientId:setClientArgs})
	ExpectOK(ret)
	client.ClientId = ret.ResArray[0].SetClientId.ResOk.ClientId
	client.Verifier = ret.ResArray[0].SetClientId.ResOk.Verifier
	hexClientId(client.ClientId)
	// TODO
	_ = client.Compound(nfs40.SetClientConfirm(client.ClientId, client.Verifier))
	return client
}

func TestSanity(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Sanity")
}
