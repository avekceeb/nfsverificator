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
	"fmt"
	"flag"
	"strings"
)

var (
	Config TestConfig
	// TODO: config? ; rotate?
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

func (cli* NFSv40Client) getRootFH() (nfs40.FH4) {
	ret := cli.Compound(nfs40.PutRootFH(), nfs40.GetFH())
	ExpectOK(ret)
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

func newNFSv40Client() (client NFSv40Client) {
	client = NewNFSv40Client(Config.ServerHost, Config.ServerPort, RandString(8)+".fake.net", Uid, Gid, RandString(8))
	Expect(client).NotTo(BeNil())
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
	r := client.Compound(nfs40.NfsArgOp4{
		ArgOp:nfs40.OP_SETCLIENTID, SetClientId:setClientArgs})
	ExpectOK(r)
	client.ClientId = r.ResArray[0].SetClientId.ResOk.ClientId
	client.Verifier = r.ResArray[0].SetClientId.ResOk.Verifier
	ExpectOK(client.Compound(nfs40.SetClientConfirm(client.ClientId, client.Verifier)))
	return client
}


func TestSanity(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Sanity")
}
