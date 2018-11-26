package v41tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/common"
)

var _ = Describe("Problematic", func() {

	Context("TODO", func() {

		It("Verify RFC 7530 16.35.5", func(){
			/* TODO: fattr4 with fh not parsed
			probably the problem is fh are of variable size ???
			One possible use of the VERIFY operation is the following COMPOUND
			sequence.  With this, the client is attempting to verify that the
			file being removed will match what the client expects to be removed.
			This sequence can help prevent the unintended deletion of a file.
			PUTFH (directory filehandle)
			LOOKUP (filename)
			VERIFY (filehandle == fh)
			PUTFH (directory filehandle)
			REMOVE (filename)
			*/
			By("Verify filehandle is the same")
			opArgs := c.OpenArgs()
			name := opArgs.Opopen.Claim.File
			r := c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				opArgs,
				Getfh())
			fh := GrabFh(&r)
			By("...this works:")
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Lookup(name),
				Verify(Fattr4{
					Attrmask:MakeGetAttrFlags(FATTR4_TYPE),
					AttrVals:[]byte{0,0,0,byte(NF4REG)}}),
				Putfh(rootFH),
				Lookup(name))
			By("...and this works:")
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Lookup(name),
				Verify(Fattr4{
					Attrmask:MakeGetAttrFlags(FATTR4_SIZE),
					AttrVals:[]byte{0,0,0,0,0,0,0,0}}),
				Putfh(rootFH),
				Lookup(name))
			By("...this doesnt:")
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Lookup(name),
				Verify(Fattr4{
					Attrmask:MakeGetAttrFlags(FATTR4_FILEHANDLE),
					AttrVals:[]byte(fh)}),
				Putfh(rootFH),
				Remove(name))
		})

		It("TODO: NFS4ERR_SEQ_FALSE_RETRY", func(){
			Skip("TODO: toxic test")
			c.Pass(
				c.SequenceArgs(),
				Putrootfh(), Getfh())
			c.Fail(
				NFS4ERR_SEQ_FALSE_RETRY,
				Sequence(c.Sid, c.Seq - 1, 0, 0, false),
				Putrootfh(), Getfh())
		})

	})
})
