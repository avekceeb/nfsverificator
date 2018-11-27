package v40tests

import (
	. "github.com/onsi/ginkgo"
	"testing"
	"flag"
	. "github.com/avekceeb/nfsverificator/common"
	. "github.com/avekceeb/nfsverificator/v40"
)

func init() {
	flag.Parse()
	Config = ReadConfig(ConfigFile)
}

var _ = Describe("Before-After", func(){

	BeforeSuite(func() {
		c = NewNFSv40Client(Config.GetHost(), Config.GetPort(), RandString(8)+".fake.net", 0, 0, RandString(8))
		// Set Client ID
		r := c.Pass(Setclientid(c.GetClientID(), c.GetCallBack(), 1))
		c.ClientId = r[0].Opsetclientid.Resok4.Clientid
		c.Verifier = r[0].Opsetclientid.Resok4.SetclientidConfirm
		c.Pass(SetclientidConfirm(c.ClientId, c.Verifier))
		export = Config.GetRWExport()
		// Get exported dir
		rootFH = c.GetExportFH(export)
	})

	AfterSuite(func() {
		c.Close()
	})

})

func TestSanity(t *testing.T) {
	Config.SuiteName = "NFSv4.0"
	RunSpecs(t, "NFSv4.0")
}
