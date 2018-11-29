package v40tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
	. "github.com/avekceeb/nfsverificator/common"
	"time"
)

var _ = Describe("Expiration", func() {

	Context("Basic", func() {

		It("Renew and state_id expired PyNFS RENEW3", func() {

			By("Creating new client")
			cliExpiring := NewNFSv40Client(
				Config.GetHost(), Config.GetPort(),
				RandString(8)+".fake.net", 0, 0, RandString(8))
			cliExpiring.SetAndConfirmClientId()

			By("Check that renew works")
			cliExpiring.Pass(Renew(cliExpiring.ClientId))

			By("Create file")
			r := cliExpiring.Pass(Putfh(rootFH), cliExpiring.OpenArgs(), Getfh())
			fh := GrabFh(&r)
			stateId := r[1].Opopen.Resok4.Stateid
			r = cliExpiring.Pass(Putfh(fh), OpenConfirm(stateId, cliExpiring.Seq))
			stateId = r[1].OpopenConfirm.Resok4.OpenStateid

			By("Imitate network partition in new client")
			interval := time.Second * time.Duration(c.LeaseTime / 6)
			for i:=0;i<7;i++ {
				time.Sleep(interval)
				c.Pass(Renew(c.ClientId))
			}
			c.Pass(Putrootfh(), Getfh())

			By("Expired client could not renew")
			cliExpiring.Fail(
				NFS4ERR_EXPIRED,
				Renew(cliExpiring.ClientId))
			cliExpiring.Close()

			//By("Old state id should not work")
			//c.Fail(NFS4ERR_EXPIRED,
			//	Putfh(fh), Close(c.Seq, stateId))
			//c.Seq-- // previous op was for other cli id...
		})

	})

})
