package v40tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
	"time"
)

var _ = Describe("Reboot", func() {

	Context("Basic", func() {

		It("Reboot", func() {
			r := c.Pass(Putfh(rootFH), c.OpenArgs(), Getfh())
			fh := GrabFh(&r)
			stateId := r[1].Opopen.Resok4.Stateid
			if CheckFlag(r[1].Opopen.Resok4.Rflags,
				OPEN4_RESULT_CONFIRM) {
				r = c.Pass(Putfh(fh), OpenConfirm(stateId, c.Seq))
				stateId = r[1].OpopenConfirm.Resok4.OpenStateid
			}
			// TODO: ???
			By("Locking file")
			c.Pass(Putfh(fh), c.LockArgs(stateId))

			By("Lock Test")
			c.Pass(Putfh(fh), c.LocktArgs(c.Id))

			By("Reboot Server")
			Config.RebootServer()

			// TODO: ping until become available
			time.Sleep(time.Second * time.Duration(10))

			By("Reconnect")
			c.Reconnect()

			By("Try again")
			c.Pass(Putfh(fh), Getfh())

			c.Fail(NFS4ERR_GRACE,
				Putfh(fh), c.LocktArgs("Other Owner"))

			c.Fail(NFS4ERR_GRACE,
				Putfh(fh), c.LocktArgs(c.Id))

			By("Waiting grace period")
			time.Sleep(time.Second * time.Duration(c.LeaseTime))
			c.Fail(NFS4ERR_STALE_CLIENTID,
				Putfh(fh), c.LocktArgs("Other Owner"))

			By("Re get client id")
			c.SetAndConfirmClientId()

			By("Network partition recovery finished by now")
			c.Pass(Putfh(fh), c.LocktArgs(c.Id))

		})

	})
})