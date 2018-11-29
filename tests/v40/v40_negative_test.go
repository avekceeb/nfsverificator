package v40tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v40"
)

var _ = Describe("Negative", func() {

	Context("Basic", func() {

		It("Bad FH (PyNFS::PUTFH2)", func() {
			c.Fail(NFS4ERR_BADHANDLE, Putfh(FhFromString("bad")), Getfh())
		})

		It("Lookup empty", func() {
			c.Fail(NFS4ERR_INVAL, Putfh(rootFH), Lookup(""))
		})

		It("No fh (PyNFS::GF9)", func() {
			c.Fail(NFS4ERR_NOFILEHANDLE, Getfh())
		})

		It("Renew Op (PyNFS::RENEW1,2)", func() {
			c.Pass(Renew(c.ClientId))
			c.Fail(NFS4ERR_STALE_CLIENTID, Renew(0))
		})

		It("Look dots", func() {
			c.Fail(NFS4ERR_BADNAME, Putfh(rootFH), Lookup("."))
			c.Fail(NFS4ERR_BADNAME, Putfh(rootFH), Lookup(".."))
		})

	})

})
