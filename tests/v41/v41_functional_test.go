package v41tests

import (
	. "github.com/onsi/ginkgo"
 	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/common"
	"time"
	"math/rand"
)

var (
	c *NFSv41Client
	rootFH NfsFh4
	blockExport  string
	rootBlockFH  NfsFh4
	// TODO: list of fh to clean up
	// get name by fh and remove in AfterSuite
)

var _ = Describe("Functional", func() {

	BeforeSuite(func() {
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
	})

	Context("Basic", func() {

		It("Layout Unavailable", func(){
			// TODO: check Config.ShareIsNotPNFS
			r := c.Pass(c.SequenceArgs(), Putfh(rootFH), c.OpenArgs(), Getfh())
			resok := r[2].Opopen.Resok4
			stateId := resok.Stateid
			fh := LastRes(&r).Opgetfh.Resok4.Object
			for layout := 0; layout < 6; layout++ {
				c.Fail(
					NFS4ERR_LAYOUTUNAVAILABLE,
					c.SequenceArgs(),
					Putfh(fh),
					Layoutget(false,
						int32(layout),
						2 /*RW*/,
						0, 4096, 4096, stateId, 4096 /*maxcount*/))
			}
			c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				Close(c.Seq, stateId))
		})

		It("PyNFS::LOOK1", func(){
			c.Fail(
				NFS4ERR_NOFILEHANDLE,
				c.SequenceArgs(),
				Lookup(RandString(12)))
		})

		It("PyNFS::LOOK2", func(){
			c.Fail(
				NFS4ERR_NOENT,
				c.SequenceArgs(),
				Putrootfh(),
				Lookup(RandString(12)))
		})

		It("PyNFS::LOOK3", func(){
			c.Fail(
				NFS4ERR_INVAL,
				c.SequenceArgs(),
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
				c.SequenceArgs(),
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
				c.SequenceArgs(),
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
				c.SequenceArgs(),
				Putrootfh(),
				Access(MakeUint32Flags(ACCESS4_READ, ACCESS4_EXECUTE)),
			)
			res := LastRes(&r).Opaccess.Resok4
			By("ensure not asked bits are not set")
			Assert(0 == (res.Supported & absentMask),
				"Wrong Supported mask")
			Assert(0 == (res.Access & absentMask),
				"Wrong Access mask")
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
			savedSeq := c.Seq
			c.Fail(
				NFS4ERR_RETRY_UNCACHED_REP,
				Sequence(c.Sid, c.Seq - 1, 0, 0, false),
				Putrootfh(), Access(MakeUint32Flags(ACCESS4_READ)))
			// Sequence ops result is OK, but compound status is fail, so
			c.Seq = savedSeq
		})

		It("TODO: Write", func(){
			r := c.Pass(c.SequenceArgs(), Putfh(rootFH), c.OpenArgs(), Getfh())
			resok := r[2].Opopen.Resok4
			stateId := resok.Stateid
			// TODO: CB_NOTIFY_LOCK
			//if ! CheckFlag(resok.Rflags, OPEN4_RESULT_MAY_NOTIFY_LOCK) {
			//	fmt.Println("TODO: Server does not Notify Lock Call")
			//}
			fh := LastRes(&r).Opgetfh.Resok4.Object
			r = c.Pass(
				c.SequenceArgs(), Putfh(fh),
				Write(stateId, 0, UNSTABLE4, []byte(RandString(128))))
		})

		It("TODO: NFS4ERR_SEQ_FALSE_RETRY", func(){
			Skip("TODO: toxic test")
			c.Pass(
				c.SequenceArgs(),
				Putrootfh(), Getfh())
			c.Fail(
				NFS4ERR_SEQ_FALSE_RETRY,
				Sequence(c.Sid, c.Seq - 1, 0, 0, false),
				Putrootfh(), Getfh())
		})

        It("Lock Sanity (PyNFS::LOCK1)", func() {
            r := c.Pass(
				c.SequenceArgs(),
				Putfh(rootFH),
				c.OpenArgs(), Getfh())
			resok := LastRes(&r).Opgetfh.Resok4
			fh := resok.Object
			sid := r[2].Opopen.Resok4.Stateid
            c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				c.LockArgs(sid))
            c.Fail(
				NFS4ERR_DENIED,
				c.SequenceArgs(),
				Putfh(fh),
				c.LocktArgs("Other owner"))
        })

		It("DestroySession", func(){
			newClient := NewNFSv41DefaultClient()
			newClient.ExchangeId()
			newClient.CreateSession()
			newClient.GetSomeAttr()
			By("Destroing session...")
			newClient.Pass(newClient.SequenceArgs(),
				DestroySession(newClient.Sid))
			By("Trying to use destroyed...")
			newClient.Fail(NFS4ERR_BADSESSION, newClient.SequenceArgs())
			newClient.Pass(DestroyClientid(newClient.ClientId))
		})

	})

	Context("Slow", func() {

        It("Session Expiration, Lock Release", func() {
			cliExpiring := NewNFSv41DefaultClient()
			cliExpiring.ExchangeId()
			cliExpiring.CreateSession()
			cliExpiring.GetSomeAttr()

            r := cliExpiring.Pass(
				cliExpiring.SequenceArgs(),
				Putfh(rootFH),
				cliExpiring.OpenArgs(), Getfh())
			resok := LastRes(&r).Opgetfh.Resok4
			fh := resok.Object
			sid := r[2].Opopen.Resok4.Stateid
            cliExpiring.Pass(
				cliExpiring.SequenceArgs(),
				Putfh(fh),
				cliExpiring.LockArgs(sid))

			By("pinging server in default client and abandon in cliExpiring")
			// supposing LeaseTime is the same
			interval := time.Second * time.Duration(c.LeaseTime / 6)
			for i:=0;i<7;i++ {
				time.Sleep(interval)
				c.Pass(c.SequenceArgs())
			}

            c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				c.LocktArgs("Other Owner"))

            cliExpiring.Fail(
				NFS4ERR_BADSESSION,
				cliExpiring.SequenceArgs(),
			)
			cliExpiring.Close()
        })

		It("CreateSession Timeout (PyNFS::EID9)", func() {
			cliStale := NewNFSv41DefaultClient()

			time.Sleep(time.Second * time.Duration(c.LeaseTime + 5))

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
