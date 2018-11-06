package v41tests

import (
    . "github.com/onsi/ginkgo"
 	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/util"
	"time"
)

var (
	c *NFSv41Client
)

var _ = Describe("Functional", func() {

	BeforeSuite(func() {
		c = NewNFSv41Client(
			Config.ServerHost, Config.ServerPort,
				RandString(8) + ".fake.net", 0, 0, RandString(8))
		c.ExchangeId()
		c.CreateSession()
		c.Seq++
	})

	AfterSuite(func() {
		c.Close()
	})

	Context("Basic", func() {

		It("PyNFS::LOOK1", func(){
			c.Fail(
				NFS4ERR_NOFILEHANDLE,
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Lookup(RandString(12)))
		})

		It("PyNFS::LOOK2", func(){
			c.Seq++
			c.Fail(
				NFS4ERR_NOENT,
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Lookup(RandString(12)))
		})

		It("PyNFS::LOOK3", func(){
			c.Seq++
			c.Fail(
				NFS4ERR_INVAL,
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Lookup(""))
		})

		It("PyNFS::LOOK4", func(){
			c.Seq++
			c.Fail(
				NFS4ERR_NAMETOOLONG,
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Lookup(RandString(4000)))
		})

		It("PyNFS::EID9", func() {
			Skip("just too long")
			c.Seq++
			l := c.Pass(
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Getfh(),
				Getattr([]uint32{MakeGetAttrFlags(FATTR4_LEASE_TIME)}))
			leaseTime := BytesToUint32(
				LastRes(&l).Opgetattr.Resok4.ObjAttributes.AttrVals)
			cliStale := NewNFSv41Client(
				Config.ServerHost, Config.ServerPort,
					RandString(8) + ".fake.net", 0, 0, RandString(8))
			time.Sleep(time.Second * time.Duration(leaseTime + 5))

			cliStale.Fail(
				NFS4ERR_STALE_CLIENTID,
				CreateSession(
					cliStale.ClientId,
					cliStale.Seq,
					DefCsFlags,
					DefChannelAttrs,
					DefChannelAttrs,
					0x40000000,
					[]CallbackSecParms4{{
						CbSecflavor:1,
						CbspSysCred:cliStale.AuthSys}}))
			cliStale.Close()
		})

	})

})


