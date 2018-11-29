package v40tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
)

var _ = Describe("Functional", func() {

	Context("Basic", func() {

		It("PyNFS LOOK9", func() {
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

		It("Range locks", func() {

			By("Open/create file for write")
			openArgs := c.OpenArgs()
			r := c.Pass(Putfh(rootFH), openArgs, Getfh())
			fh := GrabFh(&r)
			openStateId := c.OpenConfirmMacro(&r)
			openSeqId := c.Seq

			By("Write to file")
			r = c.Pass(Putfh(fh),
				Write(openStateId, 0, UNSTABLE4, make([]byte, 64)))

			By("Lock range 0-32 for write")
			lockArgs1 := c.LockArgs(openStateId)
			lockArgs1.Oplock.Offset = 0
			lockArgs1.Oplock.Length = 32
			lockArgs1.Oplock.Locker.OpenOwner.OpenSeqid = openSeqId
			r = c.Pass(Putfh(fh), lockArgs1)
			lock1StateId := r[1].Oplock.Resok4.LockStateid

			By("Lock range 32-64 for write")
			lockArgs2 := c.LockArgs(openStateId)
			lockArgs2.Oplock.Offset = 32
			lockArgs2.Oplock.Length = 32
			lockArgs2.Oplock.Locker.OpenOwner.OpenSeqid = openSeqId
			r = c.Pass(Putfh(fh), lockArgs2)
			lock2StateId := r[1].Oplock.Resok4.LockStateid

			By("Check file is locked")
			c.Fail(NFS4ERR_DENIED,
					Putfh(fh), c.LocktArgs("Other Owner"))

			By("Write to 0-32")
			r = c.Pass(Putfh(fh),
				Write(lock1StateId, 0, UNSTABLE4, []byte(RandString(32))))
			By("Write to 32-64")
			r = c.Pass(Putfh(fh),
				Write(lock2StateId, 0, UNSTABLE4, []byte(RandString(32))))

			By("Clean up")
			c.Pass(Putfh(fh), Close(openSeqId, openStateId))
			c.Pass(Putfh(rootFH), Remove(openArgs.Opopen.Claim.File))
		})

	})

})