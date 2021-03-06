package v41tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/common"
)

var _ = Describe("Errors", func() {

	Context("Basic", func() {

		It("PyNFS::LOOKP2r", func(){
			c.Fail(NFS4ERR_NOTDIR,
				c.SequenceArgs(),
				Putfh(globalFileFH),
				Lookupp())
		})

		It("Remove not existent (PyNFS::RM6)", func() {
			c.Fail(NFS4ERR_NOENT,
				c.SequenceArgs(),
				Putfh(rootFH),
				Remove(notExisting))
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

		It("Too many operations", func(){
			By("TODO: NFS4ERR_RESOURCE seems not right here (fedora 29)")
			args := []NfsArgop4{Sequence(c.Sid, c.Seq, 0, 0, false)}
			for i := uint32(0);i<c.ForeChAttr.CaMaxoperations/2 + 1;i++ {
				args = append(args, Putrootfh(), Getfh())
			}
			c.FailOneOf([]int32{NFS4ERR_TOO_MANY_OPS, NFS4ERR_RESOURCE},
				args...)
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

	})
})