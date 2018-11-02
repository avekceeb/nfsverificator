package v40tests

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/util"
)

var (
    c *V40Test
    export string
    rootFH NfsFh4
)

var _ = Describe("Functional", func() {

	BeforeSuite(func() {
		c = NewV40Test(Config.ServerHost, Config.ServerPort, RandString(8)+".fake.net", 0, 0, RandString(8))
		// Set Client ID
		r := c.ExpectOK(Setclientid(c.Client.GetClientID(), c.Client.GetCallBack(), 1))
		c.Client.ClientId = r.Resarray[0].Opsetclientid.Resok4.Clientid
		c.Client.Verifier = r.Resarray[0].Opsetclientid.Resok4.SetclientidConfirm
		c.ExpectOK(SetclientidConfirm(c.Client.ClientId, c.Client.Verifier))
		Expect(len(Config.Exports) > 0).To(BeTrue())
		export = Config.Exports[0]
		// Get exported dir
		rootFH = c.GetExportFH(export)
	})

	AfterSuite(func() {
		c.Client.Close()
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

        It("Get Same FH", func() {
            r := c.ExpectOK(Putfh(rootFH), Getfh())
            Expect(r.Resarray[1].Opgetfh.Resok4.Object).To(Equal(rootFH))
            c.ExpectErr(NFS4ERR_BADHANDLE, Putfh(FhFromString("bad")), Getfh())
        })

        It("Lookup empty", func() {
            c.ExpectErr(NFS4ERR_INVAL, Putfh(rootFH), Lookup(""))
        })

        It("No fh", func() {
            c.ExpectErr(NFS4ERR_NOFILEHANDLE, Getfh())
        })

		It("Look dots", func() {
			dir := RandString(8)
			fh := c.CreateDir(rootFH, dir, 0777)
			c.ExpectErr(NFS4ERR_BADNAME, Putfh(fh), Lookup("."))
			c.ExpectErr(NFS4ERR_BADNAME, Putfh(fh), Lookup(".."))
			c.ExpectOK(Putfh(rootFH), Remove(dir))
		})

        It("PyNFS::LOOK9", func() {
            dir := (RandString(16))
            dirFH := c.CreateDir(rootFH, dir, 0777)
            dirFH2 := c.CreateDir(dirFH, dir, 0777)
            c.SetAttr(dirFH2, 0000)
			c.Comp(0)
			r := c.Client.Compound(Putfh(rootFH), Lookup(dir), Lookup(dir))
            if c.Uid == 0 {
				Expect(r.Status).To(BeOK)
            } else {
				Expect(r.Status).To(Equal(int32(NFS4ERR_ACCESS)))
            }
        })

        It("PyNFS::LOCK1", func() {
            fileName := RandString(8)
            newFH, stateId := c.OpenSimple(rootFH, fileName)
            c.LockSimple(
				newFH,
                WRITE_LT, 0, 10, stateId)
            c.ExpectErr(
				NFS4ERR_DENIED,
				Putfh(newFH),
				Lockt(WRITE_LT, 0, 10, LockOwner4{
					Clientid: c.Client.ClientId, Owner: "Other Owner"}))
        })


	})
})