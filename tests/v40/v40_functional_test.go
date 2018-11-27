package v40tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
	"time"
)

var (
	c *NFSv40Client
	export string
	rootFH NfsFh4
)

var _ = Describe("Functional", func() {

	Context("Basic", func() {

		It("Bad FH (PyNFS::PUTFH2)", func() {
			c.Fail(NFS4ERR_BADHANDLE, Putfh(FhFromString("bad")), Getfh())
		})

		It("Lookup empty", func() {
			c.Fail(NFS4ERR_INVAL, Putfh(rootFH), Lookup(""))
		})

		It("No fh (PyNFS::GF9)", func() {
			c.Fail(NFS4ERR_NOFILEHANDLE, Getfh())
		})

		It("Renew Op (PyNFS::RENEW1,2)", func(){
			c.Pass(Renew(c.ClientId))
			c.Fail(NFS4ERR_STALE_CLIENTID, Renew(0))
		})

		It("Look dots", func() {
			createArgs := c.CreateArgs()
			r := c.Pass(Putfh(rootFH), createArgs, Getfh())
			fh := GrabFh(&r)
			c.Fail(NFS4ERR_BADNAME, Putfh(fh), Lookup("."))
			c.Fail(NFS4ERR_BADNAME, Putfh(fh), Lookup(".."))
			c.Pass(Putfh(rootFH), Remove(createArgs.Opcreate.Objname))
		})

		It("PyNFS::LOOK9", func() {
			createArgs := c.CreateArgs()
			dirName := createArgs.Opcreate.Objname
			r := c.Pass(Putfh(rootFH), createArgs, Getfh())
			dirFH := GrabFh(&r)
			r = c.Pass(Putfh(dirFH), createArgs, Getfh())
			dirFH2 := GrabFh(&r)
			c.Pass(Putfh(dirFH2),
					Setattr(Stateid4{},
						Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
							AttrVals:GetPermAttrList(0000)}))
			res, _ := c.Compound(Putfh(rootFH), Lookup(dirName), Lookup(dirName))
			if c.AuthSys.Uid == 0 {
					// root can do everything
					AssertNfsOK(res.Status)
			} else {
					AssertStatus(res.Status, NFS4ERR_ACCESS)
			}
		})

		It("PyNFS::LOCK1", func() {
			r := c.Pass(Putfh(rootFH), c.OpenArgs(), Getfh())
			fh := GrabFh(&r)
			stateId := r[1].Opopen.Resok4.Stateid
			// TODO: automatically???
			c.Seq++
			r = c.Pass(Putfh(fh), OpenConfirm(stateId, c.Seq))
			stateId = r[1].OpopenConfirm.Resok4.OpenStateid
			c.Seq++
			c.Pass(Putfh(fh), c.LockArgs(stateId))
			c.Fail(NFS4ERR_DENIED,
					Putfh(fh), c.LocktArgs("Other Owner"))
		})

	})


	Context("Slow", func() {

		It("Renew expired (PyNFS::RENEW3)", func() {
			cliExpiring := NewNFSv40Client(
				Config.GetHost(), Config.GetPort(),
				RandString(8)+".fake.net", 0, 0, RandString(8))
			r := c.Pass(Setclientid(cliExpiring.GetClientID(),
				cliExpiring.GetCallBack(), 1))
			cliExpiring.ClientId = r[0].Opsetclientid.Resok4.Clientid
			cliExpiring.Verifier = r[0].Opsetclientid.Resok4.SetclientidConfirm
			cliExpiring.Pass(SetclientidConfirm(cliExpiring.ClientId,
				cliExpiring.Verifier))
			cliExpiring.Pass(Renew(cliExpiring.ClientId))
			By("pinging server in default client and abandon in cliExpiring")
			// supposing LeaseTime is the same
			interval := time.Second * time.Duration(90 / 6)
			for i:=0;i<7;i++ {
				time.Sleep(interval)
				c.Pass(Renew(c.ClientId))
			}
			c.Pass(Putrootfh(), Getfh())
			cliExpiring.Fail(
				NFS4ERR_EXPIRED,
				Renew(cliExpiring.ClientId))
			cliExpiring.Close()
		})

		It("Expired xxx", func() {
			cliExpiring := NewNFSv40Client(
				Config.GetHost(), Config.GetPort(),
				RandString(8)+".fake.net", 0, 0, RandString(8))
			r := c.Pass(Setclientid(cliExpiring.GetClientID(),
				cliExpiring.GetCallBack(), 1))
			cliExpiring.ClientId = r[0].Opsetclientid.Resok4.Clientid
			cliExpiring.Verifier = r[0].Opsetclientid.Resok4.SetclientidConfirm
			cliExpiring.Pass(SetclientidConfirm(cliExpiring.ClientId,
				cliExpiring.Verifier))
			cliExpiring.Pass(Renew(cliExpiring.ClientId))
			r = cliExpiring.Pass(Putfh(rootFH), cliExpiring.OpenArgs(), Getfh())
			fh := GrabFh(&r)
			stateId := r[1].Opopen.Resok4.Stateid
			cliExpiring.Seq++
			r = cliExpiring.Pass(Putfh(fh), OpenConfirm(stateId, cliExpiring.Seq))
			stateId = r[1].OpopenConfirm.Resok4.OpenStateid
			cliExpiring.Seq++
			interval := time.Second * time.Duration(90 / 6)
			for i:=0;i<7;i++ {
				time.Sleep(interval)
				c.Pass(Renew(c.ClientId))
			}
			c.Fail(NFS4ERR_EXPIRED,
				Putfh(fh), Close(c.Seq, stateId))
		})

	})

})