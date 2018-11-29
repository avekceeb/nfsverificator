package v40tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
)

var _ = Describe("Functional", func() {

	Context("Basic", func() {

		It("Special stateid", func() {
			c.Pass(Putfh(globalFileFH),
				Write(anonStateId,
					0, UNSTABLE4, []byte(RandString(32))))
			c.Pass(Putfh(globalFileFH),
				Write(bypassStateId,
					0, UNSTABLE4, []byte(RandString(32))))
			c.Pass(Putfh(globalFileFH),
				Read(bypassStateId, 0, 32))
			c.Pass(Putfh(globalFileFH),
				Read(anonStateId, 0, 32))
			c.Pass(Putfh(globalFileFH),
				Setattr(anonStateId, fattrSize))
			// TODO: wrong time val???
			//c.Pass(Putfh(globalFileFH),
			//	Setattr(anonStateId, fattrMTime))
		})

		It("PyNFS LOOK9", func() {
			createArgs := c.CreateArgs()
			dirName := createArgs.Opcreate.Objname
			r := c.Pass(Putfh(rootFH), createArgs, Getfh())
			dirFH := GrabFh(&r)
			r = c.Pass(Putfh(dirFH), createArgs, Getfh())
			dirFH2 := GrabFh(&r)
			c.Pass(Putfh(dirFH2),
					Setattr(anonStateId,
						Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
							AttrVals:GetPermAttrList(0000)}))
			res, _ := c.Compound(Putfh(rootFH),
				Lookup(dirName), Lookup(dirName))
			if c.AuthSys.Uid == 0 {
					// root can do everything
					AssertNfsOK(res.Status)
			} else {
					AssertStatus(res.Status, NFS4ERR_ACCESS)
			}
		})

		It("PyNFS LOCK1", func() {
			By("New client. TODO: toxic test, spoils seqid")
			var c *NFSv40Client
			c = NewNFSv40Client(
				Config.GetHost(), Config.GetPort(),
				RandString(8) + ".fake.net", 0, 0, RandString(8))
			c.SetAndConfirmClientId()

			By("Open/create file for write")
			openArgs := c.OpenArgs()
			r := c.Pass(Putfh(rootFH), openArgs, Getfh())
			fh := GrabFh(&r)
			stateId := c.OpenConfirmMacro(&r)

			By("Lock file for write")
			r = c.Pass(Putfh(fh), c.LockArgs(stateId))
			stateId = r[1].Oplock.Resok4.LockStateid

			By("Check file is locked")
			c.Fail(NFS4ERR_DENIED,
					Putfh(fh), c.LocktArgs("Other Owner"))

			By("One more lock for write")
			lockArgs2 := Lock(
				WRITE_LT,
				false, /*reclaim*/
				0, /*offset*/
				NFS4_UINT64_MAX, /*length*/
				Locker4{
					NewLockOwner:0,
					LockOwner: ExistLockOwner4{
						LockSeqid: r[1].Oplock.Resok4.LockStateid.Seqid,
						LockStateid: stateId}})
			r = c.Pass(Putfh(fh), lockArgs2)

			By("Check file is still locked")
			c.Fail(NFS4ERR_DENIED,
					Putfh(fh), c.LocktArgs("Other Owner"))

			By("Clean up")
			c.Pass(Putfh(rootFH), Remove(openArgs.Opopen.Claim.File))
		})


	})

})