package v41tests

import (
	. "github.com/onsi/gomega"
    . "github.com/onsi/ginkgo"
 	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/util"
	"time"
	"math/rand"
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
			c.Fail(
				NFS4ERR_NOENT,
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Lookup(RandString(12)))
		})

		It("PyNFS::LOOK3", func(){
			c.Fail(
				NFS4ERR_INVAL,
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Lookup(""))
		})

		It("Not in session error", func(){
			c.Fail(
				NFS4ERR_OP_NOT_IN_SESSION,
				Putrootfh(),
				Sequence(c.Sid, c.Seq, 0, 0, false))
		})

		It("DestroySession is not the only error", func(){
			c.Fail(
				NFS4ERR_NOT_ONLY_OP,
				DestroySession(c.Sid),
				Sequence(c.Sid, c.Seq, 0, 0, false))
		})

		It("Sequence in non-first position", func(){
			c.Fail(
				NFS4ERR_SEQUENCE_POS,
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Sequence(c.Sid, c.Seq, 0, 0, false))
		})


		It("CreateSession is not the only error", func() {
			c.Fail(
				NFS4ERR_NOT_ONLY_OP,
				CreateSession(0, 0, 0, DefChannelAttrs,
					DefChannelAttrs, 0x40000000,
					[]CallbackSecParms4{}),
				Sequence(c.Sid, c.Seq, 0, 0, false))
		})

		It("ExchangeId is not the only error", func() {
			c.Fail(
				NFS4ERR_NOT_ONLY_OP,
				ExchangeId(
					ClientOwner4{CoOwnerid: RandString(14),
							CoVerifier: Verifier4{}},
					DefExchgFlags,
					DefProtect,	DefImpl),
				Sequence(c.Sid, c.Seq, 0, 0, false))
		})

		It("PyNFS::LOOK4", func(){
			c.Fail(
				NFS4ERR_NAMETOOLONG,
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Lookup(RandString(4000)))
		})

		It("TODO: NFS4ERR_TOO_MANY_OPS", func(){
			args := []NfsArgop4{Sequence(c.Sid, c.Seq, 0, 0, false)}
			for i := uint32(0);i<c.ForeChAttr.CaMaxoperations/2 + 1;i++ {
				args = append(args, Putrootfh(), Getfh())
			}
			c.Fail(
				NFS4ERR_TOO_MANY_OPS,
				args...)
		})

		It("TODO: Access", func(){
			absentMask := MakeUint32Flags(ACCESS4_DELETE,
				ACCESS4_EXTEND, ACCESS4_LOOKUP, ACCESS4_MODIFY)
			r := c.Pass(
				Sequence(c.Sid, c.Seq, 0, 0, false),
				Putrootfh(),
				Access(MakeUint32Flags(ACCESS4_READ, ACCESS4_EXECUTE)),
			)
			res := LastRes(&r).Opaccess.Resok4
			// ensure not asked bits are not set
			Expect(res.Supported & absentMask).To(Equal(uint32(0)))
			Expect(res.Access & absentMask).To(Equal(uint32(0)))
			//if 1 == res.Supported && uint32(ACCESS4_READ) {
			//}
		})

		It("TODO: bad session error", func(){
			badSid := [NFS4_SESSIONID_SIZE]byte{}
			for i:=range badSid {
				badSid[i] = byte(rand.Uint32())
			}
			c.Fail(
				NFS4ERR_BADSESSION,
				Sequence(badSid, c.Seq, 0, 0, false),
				Putrootfh(), Getfh())
		})

		It("TODO: NFS4ERR_RETRY_UNCACHED_REP", func(){
			c.Fail(
				NFS4ERR_RETRY_UNCACHED_REP,
				Sequence(c.Sid, c.Seq - 1, 0, 0, false),
				Putrootfh(), Access(MakeUint32Flags(ACCESS4_READ)))
		})

		//It("TODO: NFS4ERR_SEQ_FALSE_RETRY", func(){
		//	c.Pass(
		//		Sequence(c.Sid, c.Seq, 0, 0, true),
		//		Putrootfh(), Getfh())
		//	c.Fail(
		//		NFS4ERR_SEQ_FALSE_RETRY,
		//		Sequence(c.Sid, c.Seq - 1, 0, 0, false),
		//		Putrootfh(), Getfh())
		//})

	})

	Context("Slow", func() {

		It("CreateSession Timeout (PyNFS::EID9)", func() {
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


