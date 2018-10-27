package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/nfs40"
)

var _ = Describe("PyNFS cases v4.0", func() {

	Context("Basic", func() {

		It("NULL Call", func() {
			Expect(client.Null()).To(BeNil())
		})

		It("Get Same FH", func() {
			r := client.Compound(PutFH(rootFH), GetFH())
			ExpectOK(r)
			Expect(r.ResArray[1].GetFH.FH).To(Equal(rootFH))
			r = client.Compound(PutFH([]byte("bad")), GetFH())
			ExpectErr(r, NFS4ERR_BADHANDLE)
		})

		It("Lookup empty", func() {
			ExpectErr(client.Compound(PutFH(rootFH), Lookup("")), NFS4ERR_INVAL)
		})

		It("No fh", func() {
			ExpectErr(client.Compound(GetFH()), NFS4ERR_NOFILEHANDLE)
		})

		It("LOOK9", func() {
			dir := (RandString(16))
			r := client.Compound(PutFH(rootFH), CreateDir(dir), GetFH())
			ExpectOK(r)
			dirFH := r.ResArray[2].GetFH.FH
			r = client.Compound(PutFH(dirFH), CreateDir(dir))
			ExpectOK(r)
			r = client.Compound(PutFH(dirFH),
				SetAttr(StateId4{}, GetBitmap(FATTR4_MODE), GetPermAttrList(0000)))
			r = client.Compound(PutFH(rootFH), Lookup(dir), Lookup(dir))
			if Uid == 0 {
				ExpectOK(r)
			} else {
				ExpectErr(r, NFS4ERR_ACCESS)
			}
		})

		It("LOCK1", func() {
			fileName := RandString(8)
			r := client.Compound(
				PutFH(rootFH),
				Open(client.Seq, client.ClientId, client.Id, fileName),
				GetFH())
			client.Seq += 1
			ExpectOK(r)
			stateId := r.ResArray[1].Open.Result.StateId
			ExpectOK(r)
			newFH := r.ResArray[2].GetFH.FH
			r = client.Compound(PutFH(newFH), OpenConfirm(stateId, client.Seq))
			client.Seq += 1
			newStateId := r.ResArray[1].OpenConfirm.Result.State
			r = client.Compound(PutFH(newFH),
				Lock(WRITE_LT,0,10,client.Seq, newStateId, client.ClientId, client.Id))
			ExpectOK(r)
			r = client.Compound(PutFH(newFH), LockT(WRITE_LT, 0, 10, client.ClientId, "Other Client"))
			ExpectErr(r, NFS4ERR_DENIED)
		})

	})
})