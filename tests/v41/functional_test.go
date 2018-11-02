package v41tests

import (
    . "github.com/onsi/ginkgo"
 	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/util"
)

var _ = Describe("Functional", func() {

	//BeforeSuite(func() {
	//})
	//
	//AfterSuite(func() {
	//})

	Context("Basic", func() {

		It("Try", func() {
			c := NewV41(Config.ServerHost, Config.ServerPort, RandString(8) + ".fake.net", 0, 0, RandString(8))
			c.Connect()
		})

	})

})


