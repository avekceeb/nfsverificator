package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/nfs40"
	"strings"
)

func findExportedRoot(export string, cli40 *NFSv40Client) (fh FH4) {
	if nil == cli40 {
		*cli40 = NewNFSv40Client()
		defer cli40.Close()
	}
	ret := cli40.Compound(GetPutRootFH(), GetGetFH())
	Expect(ret[1].GetFH.Status).To(Equal(int32(NFS4_OK)))
	fh = ret[1].GetFH.FH
	for _, k := range strings.Split(export, "/") {
		if "" == k {
			continue
		}
		ret = cli40.Compound(GetPutFH(fh), GetLookup(k), GetGetFH())
		fh = ret[2].GetFH.FH
	}
	return fh
}

var _ = Describe("NFSv4.0", func() {

	Context("Basic", func() {

		It("NULL Call", func() {
			cli40 := NewNFSv40Client()
			cli40.Null()
		})

		It("Read Dir", func() {
			Expect(len(Config.Exports) > 0).To(BeTrue())
			export := Config.Exports[0]
			cli40 := NewNFSv40Client()
			fh := findExportedRoot(export, &cli40)
			ret := cli40.Compound(GetPutFH(fh), GetReadDir())
			Expect(ret[1].ReadDir.Status).To(Equal(int32(NFS4_OK)))
			// TODO
			//for _, e := range ret[1].ReadDir.Result.DirList.Entries {
			//	fmt.Printf(" entry: %s\n", e.Name)
			//}

			ret = cli40.Compound(GetPutFH(fh), CreateDir(RandString(16)))
			Expect(ret[1].Create.Status).To(Equal(int32(NFS4_OK)))

			// Create File:
			ret = cli40.Compound(
				GetPutFH(fh),
				GetOpen(cli40.ClientId, cli40.Id, RandString(8)),
				GetGetFH())
			Expect(ret[1].Open.Status).To(Equal(int32(NFS4_OK)))
			stateId := ret[1].Open.Result.StateId
			Expect(ret[2].GetFH.Status).To(Equal(int32(NFS4_OK)))
			newFH := ret[2].GetFH.FH

			cli40_another := NewNFSv40Client()
			cli40_another.Null()

			ret = cli40.Compound(GetPutFH(newFH),
					GetSetAttr(stateId, GetBitmap(FATTR4_MODE), GetPermAttrList(0700)))

			cli40_another.Compound(GetPutFH(newFH),
				GetGetAttr(FATTR4_MODE, FATTR4_SIZE))

		})
	})
})
