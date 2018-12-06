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
			cli := DefaultClient40()
			cli.SetAndConfirmClientId()

			By("New client 2")
			cli2 := DefaultClient40()
			cli2.SetAndConfirmClientId()

			openArgs := cli.OpenArgs()
			openArgs2 := cli2.OpenNoCreateArgs()
			openArgs2.Opopen.Claim.File = openArgs.Opopen.Claim.File
			openArgs2.Opopen.ShareAccess = OPEN4_SHARE_ACCESS_WRITE

			r := cli.Pass(Putfh(rootFH), openArgs, Getfh())
			r2 := cli2.Pass(Putfh(rootFH), openArgs2, Getfh())
			fh := GrabFh(&r)
			Assert40.Assert(AreFhEqual(fh, GrabFh(&r2)),
				"Fh should be the same")

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

	})

})