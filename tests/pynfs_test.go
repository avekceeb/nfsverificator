package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/nfs40"
)

var _ = Describe("PyNFS cases v4.0", func() {

	Context("Basic", func() {

		It("NULL Call", func() {
			client.Null()
		})

		It("Get Same FH", func() {
			ret := client.Compound(PutFH(rootFH), GetFH())
			ExpectOK(ret)
			Expect(ret.ResArray[1].GetFH.FH).To(Equal(rootFH))
			ret = client.Compound(PutFH([]byte("bad")), GetFH())
			ExpectErr(ret, NFS4ERR_BADHANDLE)
		})

		It("Lookup empty", func() {
			ExpectErr(client.Compound(PutFH(rootFH), Lookup("")), NFS4ERR_INVAL)
		})

		It("No fh", func() {
			ExpectErr(client.Compound(GetFH()), NFS4ERR_NOFILEHANDLE)
		})

		It("LOOK9", func() {
			dir := (RandString(16))
			ret := client.Compound(PutFH(rootFH), CreateDir(dir), GetFH())
			ExpectOK(ret)
			dirFH := ret.ResArray[2].GetFH.FH
			ret = client.Compound(PutFH(dirFH), CreateDir(dir))
			ExpectOK(ret)
			ret = client.Compound(PutFH(dirFH),
				SetAttr(StateId4{}, GetBitmap(FATTR4_MODE), GetPermAttrList(0000)))
			ret = client.Compound(PutFH(rootFH), Lookup(dir), Lookup(dir))
			if Uid == 0 {
				ExpectOK(ret)
			} else {
				ExpectErr(ret, NFS4ERR_ACCESS)
			}
		})

		It("LOCK1", func() {
			fileName := RandString(8)
			ret := client.Compound(
				PutFH(rootFH),
				Open(client.Seq, client.ClientId, client.Id, fileName),
				GetFH())
			client.Seq += 1
			ExpectOK(ret)
			stateId := ret.ResArray[1].Open.Result.StateId
			ExpectOK(ret)
			newFH := ret.ResArray[2].GetFH.FH
			ret = client.Compound(PutFH(newFH), OpenConfirm(stateId, client.Seq))
			client.Seq += 1
			newStateId := ret.ResArray[1].OpenConfirm.Result.State
			ret = client.Compound(PutFH(newFH),
				Lock(WRITE_LT,0,10,client.Seq, newStateId, client.ClientId, client.Id))
			ExpectOK(ret)
			ret = client.Compound(PutFH(newFH), LockT(WRITE_LT, 0, 10, client.ClientId, "Other Client"))
			ExpectErr(ret, NFS4ERR_DENIED)
		})

	})
})