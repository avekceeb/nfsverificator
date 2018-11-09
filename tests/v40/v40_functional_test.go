package v40tests

import (
    . "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
)

var (
    c *NFSv40Client
    export string
    rootFH NfsFh4
)

var _ = Describe("Functional", func() {

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

	Context("Basic", func() {

		//It("Read Dir", func() {
		//	//Skip("Not ready")
		//	ret := c.ExpectOK(Putfh(rootFH), Create..(RandString(16)))
		//	ret = c.ExpectOK(Putfh(rootFH), Readdir(...))
		//	// TODO
		//	//for _, e := range ret[1].ReadDir.Result.DirList.Entries {
		//	//  fmt.Printf(" entry: %s\n", e.Name)
		//	//}
		//})

        //It("Same FH (PyNFS::PUTFH1r)", func() {
        //    r := c.Pass(Putfh(rootFH), Getfh())
        //    Assert(bytes.Equal(r[1].Opgetfh.Resok4.Object, rootFH),
			//	"fh should be same")
        //})

        It("Bad FH (PyNFS::PUTFH2)", func() {
            c.Fail(NFS4ERR_BADHANDLE, Putfh(FhFromString("bad")), Getfh())
        })

        It("Lookup empty", func() {
            c.Fail(NFS4ERR_INVAL, Putfh(rootFH), Lookup(""))
        })

        It("No fh (PyNFS::GF9)", func() {
            c.Fail(NFS4ERR_NOFILEHANDLE, Getfh())
        })

		It("Look dots", func() {
			dir := RandString(8)
			fh := c.CreateDir(rootFH, dir, 0777)
			c.Fail(NFS4ERR_BADNAME, Putfh(fh), Lookup("."))
			c.Fail(NFS4ERR_BADNAME, Putfh(fh), Lookup(".."))
			c.Pass(Putfh(rootFH), Remove(dir))
		})

        It("PyNFS::LOOK9", func() {
            dir := (RandString(16))
            dirFH := c.CreateDir(rootFH, dir, 0777)
            dirFH2 := c.CreateDir(dirFH, dir, 0777)
            c.SetAttr(dirFH2, 0000)
			r, _ := c.Compound(Putfh(rootFH), Lookup(dir), Lookup(dir))
            if c.AuthSys.Uid == 0 {
				AssertNfsOK(r.Status)
            } else {
				AssertStatus(r.Status, NFS4ERR_ACCESS)
            }
        })

        It("PyNFS::LOCK1", func() {
            fileName := RandString(8)
            newFH, stateId := c.OpenSimple(rootFH, fileName)
            c.LockSimple(
				newFH,
                WRITE_LT, 0, 10, stateId)
            c.Fail(
				NFS4ERR_DENIED,
				Putfh(newFH),
				Lockt(WRITE_LT, 0, 10, LockOwner4{
					Clientid: c.ClientId, Owner: "Other Owner"}))
        })


	})
})