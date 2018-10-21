package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/nfs40"
)

var (
	client NFSv40Client
	export string
	rootFH FH4
)

var _ = Describe("NFSv4.0", func() {

	BeforeSuite(func() {
		client = NewNFSv40Client()
		Expect(len(Config.Exports) > 0).To(BeTrue())
		export = Config.Exports[0]
		rootFH = client.getExportFH(export)
	})

	AfterSuite(func() {
		client.Close()
	})

	Context("Basic", func() {

		It("Read Dir", func() {
			ret := client.Compound(PutFH(rootFH), CreateDir(RandString(16)))
			ExpectOK(ret)
			ret = client.Compound(PutFH(rootFH), ReadDir())
			ExpectOK(ret)
			// TODO
			//for _, e := range ret[1].ReadDir.Result.DirList.Entries {
			//	fmt.Printf(" entry: %s\n", e.Name)
			//}
		})

		It("Write File", func() {
			// Client1
			cli1 := NewNFSv40Client()
			fh := cli1.getExportFH(export)

			// Client2
			cli2 := NewNFSv40Client()
			//cli2 := createNFSv40Client()
			// 1
			fileName := RandString(8)
			ret := cli1.Compound(
				PutFH(fh),
				Open(cli1.Seq, cli1.ClientId, cli1.Id, fileName),
				GetFH(),
				GetAttr(FATTR4_FH_EXPIRE_TYPE))
			cli1.Seq += 1
			ExpectOK(ret)
			stateId := ret.ResArray[1].Open.Result.StateId
			ExpectOK(ret)
			newFH := ret.ResArray[2].GetFH.FH
			ret = cli1.Compound(PutFH(newFH), OpenConfirm(stateId, cli1.Seq))
			cli1.Seq += 1
			newStateId := ret.ResArray[1].OpenConfirm.Result.State

			// 2
			retOther := cli2.Compound(
				PutFH(fh),
				Open(cli2.Seq, cli2.ClientId, cli2.Id, fileName),
				GetFH())
			ExpectOK(retOther)
			cli2.Seq += 1
			stateIdOther := retOther.ResArray[1].Open.Result.StateId
			newFHOther := retOther.ResArray[2].GetFH.FH

			d := []byte{0x41,0x41,0x41,0x41,0x41,0x41}
			ret = cli1.Compound(PutFH(newFH), Write(newStateId, &d, uint64(32)))
			// Close 1
			ret = cli1.Compound(PutFH(newFH), Close(cli1.Seq, newStateId))
			cli1.Seq += 1


			retOther = cli2.Compound(PutFH(newFHOther),
				OpenConfirm(stateIdOther, cli2.Seq))
			cli2.Seq += 1
			newStateIdOther := retOther.ResArray[1].OpenConfirm.Result.State
			retOther = cli2.Compound(PutFH(newFHOther),
				Write(newStateIdOther, &d, uint64(0)))
			retOther = cli2.Compound(PutFH(newFHOther), Close(cli2.Seq, newStateIdOther))
			cli2.Seq += 1


			//ret = cli40.Compound(GetPutFH(newFH),
			//		GetSetAttr(stateId, GetBitmap(FATTR4_MODE), GetPermAttrList(0700)))
			//cli40_another := NewNFSv40Client()
			//cli40_another.Compound(GetPutFH(newFH),
			//	GetGetAttr(FATTR4_MODE, FATTR4_SIZE))

			cli1.Close()
			cli2.Close()

		})

		It("Sequence Id", func() {
			Skip("This is scale")
			Expect(len(Config.Exports) > 0).To(BeTrue())
			export := Config.Exports[0]
			cli1 := NewNFSv40Client()
			fh := cli1.getExportFH(export)
			fileName := RandString(8)
			ret := cli1.Compound(
				PutFH(fh),
				Open(cli1.Seq, cli1.ClientId, cli1.Id, fileName),
				GetFH())
			cli1.Seq += 1
			ExpectOK(ret)
			stateId := ret.ResArray[1].Open.Result.StateId
			ExpectOK(ret)
			newFH := ret.ResArray[2].GetFH.FH
			ret = cli1.Compound(PutFH(newFH), OpenConfirm(stateId, cli1.Seq))
			cli1.Seq += 1
			newStateId := ret.ResArray[1].OpenConfirm.Result.State
			ret = cli1.Compound(PutFH(newFH), Close(cli1.Seq, newStateId))
			cli1.Seq += 1
			var i uint64 = 0
			for i < (4294967296 + 5) {
				ret = cli1.Compound(
					PutFH(fh),
					Open(cli1.Seq, cli1.ClientId, cli1.Id, fileName),
					GetFH())
				cli1.Seq += 1
				ExpectOK(ret)
				stateId := ret.ResArray[1].Open.Result.StateId
				newFH := ret.ResArray[2].GetFH.FH
				ret = cli1.Compound(PutFH(newFH), Close(cli1.Seq, stateId))
				cli1.Seq += 1
				i += 1
			}
		})

	})
})
