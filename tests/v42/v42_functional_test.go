package v42tests

import (
	. "github.com/onsi/ginkgo"
)

var (
	c *NFSv42Client
)

var _ = Describe("Functional", func() {

	BeforeSuite(func() {
		c = NewNFSv42DefaultClient()
	})

	AfterSuite(func() {
		c.Close()
	})

	Context("Basic", func() {
		It("TODO", func() {
			c.Null()
			c.ExchangeId()
			c.CreateSession()
			c.GetSomeAttr()
		})
	})
})
