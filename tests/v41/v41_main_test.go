package v41tests

import (
    "testing"
    "flag"
    . "github.com/onsi/ginkgo"
    . "github.com/avekceeb/nfsverificator/common"
	. "github.com/avekceeb/nfsverificator/v41"
)

const (
	SuiteName = "NFSv4.1"
)

var (
	c            *NFSv41Client
	rootFH       NfsFh4
	blockExport  string
	rootBlockFH  NfsFh4
	// TODO: list of fh to clean up
	// get name by fh and remove in AfterSuite
)

func init() {
    flag.Parse()
    Config = ReadConfig(ConfigFile)
}

var _ = Describe("Before-After", func() {

	BeforeSuite(func() {
		Config.RunExternalCommands()
		c = NewNFSv41DefaultClient()
		c.ExchangeId()
		c.CreateSession()
		c.GetSomeAttr()
		rootFH = c.LookupFromRoot(Config.GetRWExport())
		blockExport = Config.GetBlockExport()
		if "" != blockExport {
			rootBlockFH = c.LookupFromRoot(blockExport)
		}
	})

	AfterSuite(func() {
		// TODO: DestroySession
		// DestroyClientId
		c.Close()
		Config.StopExternalCommands()
	})
})

func TestSanity(t *testing.T) {
	Config.SuiteName = SuiteName
    RunSpecs(t, SuiteName)
}

