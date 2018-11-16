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
	// TODO: create file, dir, link
	globalFile   string
	globalDir    string
	globalFileFH NfsFh4
	globalDirFH  NfsFh4
	notExisting  string
)

func init() {
    flag.Parse()
    Config = ReadConfig(ConfigFile)
	notExisting = "not-exists-" + RandString(16)
	globalFile  = "global-" + RandString(16)
	globalDir   = "global-" + RandString(16)
}

var _ = Describe("Before-After", func() {

	BeforeSuite(func() {
		Config.RunExternalCommands()
		c = NewNFSv41DefaultClient()
		c.ExchangeId()
		c.CreateSession()
		c.GetSomeAttr()
		By("Saving some fh for future use in tests")
		rootFH = c.LookupFromRoot(Config.GetRWExport())
		blockExport = Config.GetBlockExport()
		if "" != blockExport {
			rootBlockFH = c.LookupFromRoot(blockExport)
		}
		By("Creating some global objects for tests")
		By("Global File")
		openArgs := c.OpenArgs()
		openArgs.Opopen.Claim.File = globalFile
		r := c.Pass(
			c.SequenceArgs(),
			Putfh(rootFH),
			openArgs, Getfh())
		resok := LastRes(&r).Opgetfh.Resok4
		globalFileFH = resok.Object
		sid := r[2].Opopen.Resok4.Stateid
		c.Pass(c.SequenceArgs(),
			Putfh(globalFileFH),
			Close(c.Seq, sid))
		By("Global dir")
		r = c.Pass(c.SequenceArgs(),
			Putfh(rootFH),
			c.CreateArgs(globalDir),
			Getfh())
		globalDirFH = LastRes(&r).Opgetfh.Resok4.Object
	})

	AfterSuite(func() {
		// TODO: DestroySession
		// DestroyClientId
		// Delete global
		c.Close()
		Config.StopExternalCommands()
	})
})

func TestSanity(t *testing.T) {
	Config.SuiteName = SuiteName
    RunSpecs(t, SuiteName)
}

