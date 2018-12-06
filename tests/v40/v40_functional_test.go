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
					Assert40.AssertNfsOK(res.Status)
			} else {
					Assert40.AssertStatus(res.Status, NFS4ERR_ACCESS)
			}
		})

		It("PyNFS LOCK1", func() {
			By("New client. TODO: toxic test, spoils seqid")
			var c *NFSv40Client
			c = DefaultClient40()
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

		It("Lock held on close", func() {
			c1 := DefaultClient40()
			c1.SetAndConfirmClientId()
			openArgs := c1.OpenArgs()
			//openArgs.Opopen.ShareAccess = OPEN4_SHARE_ACCESS_BOTH
			//openArgs.Opopen.ShareDeny = OPEN4_SHARE_DENY_WRITE
			//openArgs.Opopen.Openhow.How.Mode = GUARDED4
			r := c1.Pass(Putfh(rootFH), openArgs, Getfh())
			fh := GrabFh(&r)
			openStateId_1 := c1.OpenConfirmMacro(&r)
			lockArgs := c1.LockArgs(openStateId_1)
			c1.Pass(Putfh(fh), lockArgs)
			By(`RFC7530 9.10
			The CLOSE operation removes all share reservations held by the
			open-owner on that file.  If byte-range locks are held, the client
			SHOULD release all locks before issuing a CLOSE.  The server MAY free
			all outstanding locks on CLOSE, but some servers may not support the
			CLOSE of a file that still has byte-range locks held.  The server
			MUST return failure, NFS4ERR_LOCKS_HELD, if any locks would exist
			after the CLOSE.`)
			r = c1.FailOneOf([]int32{NFS4_OK, NFS4ERR_LOCKS_HELD},
				Putfh(fh), Close(c1.Seq, openStateId_1))
			if int32(NFS4_OK) == r[1].Opclose.Status {
				By("Lock should be released")
				c1.Pass(Putfh(fh), c1.LocktArgs(RandString(8)))
			} else {
				By("Lock should remain")
				c1.Fail(NFS4ERR_DENIED, Putfh(fh), c1.LocktArgs(RandString(8)))
			}
			c1.Pass(Putfh(rootFH), Remove(openArgs.Opopen.Claim.File))
		})


		It("Lock ranges", func() {

			By("New client 1")
			c1 := DefaultClient40()
			c1.SetAndConfirmClientId()

			By("New client 2")
			c2 := DefaultClient40()
			c2.SetAndConfirmClientId()

			By("Open/create file for write, denying writes")
			openArgs := c1.OpenArgs()
			openArgs.Opopen.ShareDeny = OPEN4_SHARE_DENY_NONE
			r := c1.Pass(Putfh(rootFH), openArgs, Getfh())
			fh := GrabFh(&r)
			openStateId_1 := c1.OpenConfirmMacro(&r)

			By("Open in client 2")
			openNoCreate := c2.OpenNoCreateArgs()
			openNoCreate.Opopen.Claim.File = openArgs.Opopen.Claim.File
			openNoCreate.Opopen.ShareAccess = OPEN4_SHARE_ACCESS_BOTH
			r = c2.Pass(Putfh(rootFH), openNoCreate, Getfh())
			fh2 := GrabFh(&r)
			openStateId_2 := c2.OpenConfirmMacro(&r)

			Assert40.Assert(AreFhEqual(fh, fh2), "fh not equal")

			By("Write to file")
			r = c1.Pass(Putfh(fh),
				Write(openStateId_1, 0, UNSTABLE4, []byte(RandString(16))))
			By("Write to file in client 2")
			r = c2.Pass(Putfh(fh2),
				Write(openStateId_2, 16, UNSTABLE4, []byte(RandString(16))))

			By("Lock range 0-32 for write")
			lockArgs1 := c1.LockArgs(openStateId_1)
			lockArgs1.Oplock.Offset = 0
			lockArgs1.Oplock.Length = 32
			r = c1.Pass(Putfh(fh), lockArgs1)
			lock1StateId := r[1].Oplock.Resok4.LockStateid

			By("Lock range 32-64 for write")
			lockArgs2 := c2.LockArgs(openStateId_2)
			lockArgs2.Oplock.Offset = 32
			lockArgs2.Oplock.Length = 32
			r = c2.Pass(Putfh(fh), lockArgs2)
			lock2StateId := r[1].Oplock.Resok4.LockStateid

			By("Check file is locked in default client")
			c.Fail(NFS4ERR_DENIED,
					Putfh(fh), c.LocktArgs("Other Owner"))
			By(`
			TODO: RFC7530 9.1.4.5 lock stateid SHOULD be used
			... but it is seemingly not enforced
			`)
			c1.Pass(Putfh(fh),
				Write(openStateId_1, 0, UNSTABLE4, []byte(RandString(8))))
			c2.Pass(Putfh(fh),
				Write(anonStateId,
					0, UNSTABLE4, []byte(RandString(32))))

			By("Write to 0-32")
			c1.Pass(Putfh(fh),
				Write(lock1StateId, 0, UNSTABLE4, []byte(RandString(16))))
			By("Write to 32-64")
			c2.Pass(Putfh(fh),
				Write(lock2StateId, 0/*32!!!*/, UNSTABLE4, []byte(RandString(16))))

			lockuArgs := c1.LockuArgs(lock1StateId)
			lockuArgs.Oplocku.Seqid = 1
			lockuArgs2 := c2.LockuArgs(lock2StateId)
			lockuArgs2.Oplocku.Seqid = 1
			c2.Pass(Putfh(fh), lockuArgs2)
			c1.Pass(Putfh(fh), lockuArgs)

			c1.Pass(Putfh(rootFH), Remove(openArgs.Opopen.Claim.File))
		})


	})

})