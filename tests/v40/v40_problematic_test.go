package v40tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
)

var _ = Describe("Problematic", func() {

	Context("Basic", func() {

		It("Release lock owner RFC7530 16.37", func() {

			By("New client 1")
			cli := NewNFSv40Client(
				Config.GetHost(), Config.GetPort(),
				RandString(8) + ".fake.net", 0, 0, RandString(8))
			cli.SetAndConfirmClientId()

			By("New client 2")
			cli2 := NewNFSv40Client(
				Config.GetHost(), Config.GetPort(),
				RandString(8) + ".fake.net", 0, 0, RandString(8))
			cli2.SetAndConfirmClientId()

			openArgs := cli.OpenArgs()
			openArgs2 := cli2.OpenNoCreateArgs()
			openArgs2.Opopen.Claim.File = openArgs.Opopen.Claim.File
			openArgs2.Opopen.ShareAccess = OPEN4_SHARE_ACCESS_WRITE

			r := cli.Pass(Putfh(rootFH), openArgs, Getfh())
			r2 := cli2.Pass(Putfh(rootFH), openArgs2, Getfh())
			fh := GrabFh(&r)
			Assert(AreFhEqual(fh, GrabFh(&r2)), "Fh should be the same")

			stateId := r[1].Opopen.Resok4.Stateid
			stateId2 := r2[1].Opopen.Resok4.Stateid

			if CheckFlag(r[1].Opopen.Resok4.Rflags,
				OPEN4_RESULT_CONFIRM) {
				r = cli.Pass(Putfh(fh), OpenConfirm(stateId, cli.Seq))
				stateId = r[1].OpopenConfirm.Resok4.OpenStateid
			}

			if CheckFlag(r2[1].Opopen.Resok4.Rflags,
				OPEN4_RESULT_CONFIRM) {
				r2 = cli2.Pass(Putfh(fh), OpenConfirm(stateId2, cli2.Seq))
				stateId2 = r2[1].OpopenConfirm.Resok4.OpenStateid
			}

			By("Locking file")
			r = cli.Pass(Putfh(fh), cli.LockArgs(stateId))
			stateId = r[1].Oplock.Resok4.LockStateid

			By("Trying to lock in client2")
			cli2.Fail(NFS4ERR_DENIED, Putfh(fh), cli2.LockArgs(stateId2))

			By("Check ReleaseLockowner returns NFS4ERR_LOCKS_HELD")
			By(`
			RFC7530 16.37: If file locks associated with the lock_owner
			are held at the server,	the error NFS4ERR_LOCKS_HELD will be returned
			and no further action will be taken`)
			cli.Fail(NFS4ERR_LOCKS_HELD,
				ReleaseLockowner(LockOwner4{
					Clientid:cli.ClientId, Owner:cli.Id}))

			By("Again trying to lock in client2")
			cli2.Fail(NFS4ERR_DENIED, Putfh(fh), cli2.LockArgs(stateId2))

			//cli.Seq = 0 //=> NFS4ERR_SERVERFAULT (ubuntu 4.15.0-39-generic)
			// TODO: NFS4ERR_BAD_SEQID, so previous Release was done
			//cli.Pass(Putfh(fh), cli.LockuArgs(stateId))
			//cli.Pass(ReleaseLockowner(LockOwner4{
			//		Clientid:cli.ClientId, Owner:cli.Id}))
			//cli2.Pass(Putfh(fh), cli2.LockArgs(stateId2))
			cli.Pass(Putfh(rootFH), Remove(openArgs.Opopen.Claim.File))
			cli.Close()
			cli2.Close()
		})

		It("Range locks", func() {

			By("Open/create file for write, denying writes")
			openArgs := c.OpenArgs()
			openArgs.Opopen.ShareDeny = OPEN4_SHARE_DENY_WRITE
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

			By(`
			TODO: RFC7530 9.1.4.5 lock stateid SHOULD be used
			... but it is seemingly not enforced
			`)
			c.Pass(Putfh(fh),
				Write(openStateId, 0, UNSTABLE4, []byte(RandString(32))))
			c.Fail(NFS4ERR_LOCKED,
				Putfh(fh),
				Write(anonStateId,
					0, UNSTABLE4, []byte(RandString(32))))

			By("TODO: write to range of other lock ???")
			r = c.Pass(Putfh(fh),
				Write(lock2StateId/*!!!*/, 0, UNSTABLE4, []byte(RandString(32))))
			By("Write to 0-32")
			r = c.Pass(Putfh(fh),
				Write(lock1StateId, 0, UNSTABLE4, []byte(RandString(32))))
			By("Write to 32-64")
			r = c.Pass(Putfh(fh),
				Write(lock2StateId, 32, UNSTABLE4, []byte(RandString(32))))

			By("Unlock range. TODO: this seems to be not REQUIRED")
			c.Pass(Putfh(fh), c.LockuArgs(lock1StateId))
			c.Pass(Putfh(fh), c.LockuArgs(lock2StateId))

			By("Clean up")
			By(`TODO: RFC7530 9.10
			The CLOSE operation removes all share reservations held by the
			open-owner on that file.  If byte-range locks are held, the client
			SHOULD release all locks before issuing a CLOSE.  The server MAY free
			all outstanding locks on CLOSE, but some servers may not support the
			CLOSE of a file that still has byte-range locks held.  The server
			MUST return failure, NFS4ERR_LOCKS_HELD, if any locks would exist
			after the CLOSE.
			`)
			c.Pass(Putfh(fh), Close(openSeqId, openStateId))
			c.Pass(Putfh(rootFH), Remove(openArgs.Opopen.Claim.File))
		})

	})

})