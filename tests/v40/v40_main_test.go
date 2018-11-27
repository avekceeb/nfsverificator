package v40tests

import (
	. "github.com/onsi/ginkgo"
	"testing"
	"flag"
	. "github.com/avekceeb/nfsverificator/common"
	. "github.com/avekceeb/nfsverificator/v40"
)

var (
	c *NFSv40Client
	export string
	rootFH NfsFh4
)

func init() {
	flag.Parse()
	Config = ReadConfig(ConfigFile)
}

var _ = Describe("Before-After", func(){

	BeforeSuite(func() {
		c = NewNFSv40Client(Config.GetHost(), Config.GetPort(), RandString(8)+".fake.net", 0, 0, RandString(8))
		c.SetAndConfirmClientId()
		c.GetSomeAttr()
		export = Config.GetRWExport()
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
