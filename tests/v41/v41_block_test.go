package v41tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/common"
)


var _ = Describe("pNFS", func() {

	skipIfNoBlock := func() {
		if "" == blockExport {
			Skip("No block layout pNFS export")
		}
	}

	Context("Block", func() {

		It("Layoutget", func() {
			skipIfNoBlock()
			if ! CheckFlag(c.EidFlags, EXCHGID4_FLAG_USE_PNFS_MDS) {
				Skip("Server is not MDS")
			}
			r := c.Pass(c.SequenceArgs(), Putfh(rootBlockFH), c.OpenArgs(), Getfh())
			resok := r[2].Opopen.Resok4
			stateId := resok.Stateid
			fh := LastRes(&r).Opgetfh.Resok4.Object
			c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				Layoutget(false,
					LAYOUT4_BLOCK_VOLUME,
					LAYOUTIOMODE4_RW,
					0, 4096, 4096, stateId, 4096 /*maxcount*/))
			c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				Close(c.Seq, stateId))
		})

		It("GetDeviceInfo", func() {
			skipIfNoBlock()
			Skip("TODO")
		})

	})
})