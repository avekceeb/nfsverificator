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
	globalFile   string
	globalDir    string
	globalFileFH NfsFh4
	globalDirFH  NfsFh4
	notExisting  string
	globalContent string
)

func init() {
	flag.Parse()
	Config = ReadConfig(ConfigFile)
	notExisting = "v4-not-exists-" + RandString(16)
	globalFile  = "v4-global-" + RandString(16)
	globalDir   = "v4-global-" + RandString(16)
	globalContent = "v4-" + RandString(20)
}

var _ = Describe("Before-After", func(){

	BeforeSuite(func() {
		c = NewNFSv40Client(Config.GetHost(), Config.GetPort(),
			RandString(8)+".fake.net", 0, 0, RandString(8))
		c.SetAndConfirmClientId()
		c.GetSomeAttr()
		export = Config.GetRWExport()
		rootFH = c.GetExportFH(export)
		By("Global File")
		openArgs := c.OpenArgs()
		openArgs.Opopen.Claim.File = globalFile
		r := c.Pass(Putfh(rootFH), openArgs, Getfh())
		globalFileFH = GrabFh(&r)
		sid := r[1].Opopen.Resok4.Stateid
		// !!! v4.0 specific:
		// RFC 7530 16.16.5
		// OPEN4_RESULT_CONFIRM indicates that the client MUST execute an
   		// OPEN_CONFIRM operation before using the open file.
		if CheckFlag(r[1].Opopen.Resok4.Rflags,
			OPEN4_RESULT_CONFIRM) {
			r = c.Pass(Putfh(globalFileFH), OpenConfirm(sid, c.Seq))
			sid = r[1].OpopenConfirm.Resok4.OpenStateid
		}
		By("Write to file")
		r = c.Pass(Putfh(globalFileFH),
				Write(sid, 0, UNSTABLE4, []byte(globalContent)))
		By("Close file")
		c.Pass(Putfh(globalFileFH), Close(c.Seq, sid))
		By("Global dir")
		createArgs := c.CreateArgs()
		createArgs.Opcreate.Objname = globalDir
		r = c.Pass(Putfh(rootFH), createArgs, Getfh())
		globalDirFH = LastRes(&r).Opgetfh.Resok4.Object

	})

	AfterSuite(func() {
		// TODO: delete global
		c.Pass(Putfh(rootFH), Remove(globalDir))
		c.Pass(Putfh(rootFH), Remove(globalFile))
		c.Close()
	})

})

func TestSanity(t *testing.T) {
	Config.SuiteName = "NFSv4.0"
	RunSpecs(t, "NFSv4.0")
}
