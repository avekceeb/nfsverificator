package v41tests

import (
	. "github.com/onsi/ginkgo"
 	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/common"
	"time"
	"math/rand"
)

/*
	TODO: refer | migration
   When the EXCHGID4_FLAG_SUPP_MOVED_REFER flag bit is set, the client
   indicates that it is capable of dealing with an NFS4ERR_MOVED error
   as part of a referral sequence.
   If the server will potentially perform a referral, it MUST set
   EXCHGID4_FLAG_SUPP_MOVED_REFER in eir_flags.

   When the EXCHGID4_FLAG_SUPP_MOVED_MIGR is set, the client indicates
   that it is capable of dealing with an NFS4ERR_MOVED error as part of
   a file system migration sequence.
   If the server
   will potentially perform a migration, it MUST set
   EXCHGID4_FLAG_SUPP_MOVED_MIGR in eir_flags.

https://www.ietf.org/mail-archive/web/nfsv4/current/msg00919.html
https://wiki.linux-nfs.org/wiki/index.php/NFS_Recovery_and_Client_Migration
 */

var _ = Describe("Functional", func() {

	Context("Basic", func() {

/* TODO:
RFC 5661 Section 13.  NFSv4.1 as a Storage Protocol in pNFS: the File Layout Type

   The client MAY request zero or more of EXCHGID4_FLAG_USE_NON_PNFS,
   EXCHGID4_FLAG_USE_PNFS_DS, or EXCHGID4_FLAG_USE_PNFS_MDS, even though
   some combinations (e.g., EXCHGID4_FLAG_USE_NON_PNFS |
   EXCHGID4_FLAG_USE_PNFS_MDS) are contradictory.  However, the server
   MUST only return the following acceptable combinations:

        +--------------------------------------------------------+
        | Acceptable Results from EXCHANGE_ID                    |
        +--------------------------------------------------------+
        | EXCHGID4_FLAG_USE_PNFS_MDS                             |
        | EXCHGID4_FLAG_USE_PNFS_MDS | EXCHGID4_FLAG_USE_PNFS_DS |
        | EXCHGID4_FLAG_USE_PNFS_DS                              |
        | EXCHGID4_FLAG_USE_NON_PNFS                             |
        | EXCHGID4_FLAG_USE_PNFS_DS | EXCHGID4_FLAG_USE_NON_PNFS |
        +--------------------------------------------------------+

   As the above table implies, a server can have one or two roles.  A
   server can be both a metadata server and a data server, or it can be
   both a data server and non-metadata server.  In addition to returning
   two roles in the EXCHANGE_ID's results, and thus serving both roles
   via a common client ID, a server can serve two roles by returning a
   unique client ID and server owner for each role in each of two
   EXCHANGE_ID results, with each result indicating each role.

 */

		It("ExchangeId flags combinations", func() {
			validCombinations := []uint32{
				uint32(EXCHGID4_FLAG_USE_PNFS_MDS),
				uint32(EXCHGID4_FLAG_USE_PNFS_DS),
				uint32(EXCHGID4_FLAG_USE_NON_PNFS),
				MakeUint32Flags(EXCHGID4_FLAG_USE_PNFS_MDS,
					EXCHGID4_FLAG_USE_PNFS_DS),
				MakeUint32Flags(EXCHGID4_FLAG_USE_PNFS_DS,
					EXCHGID4_FLAG_USE_NON_PNFS)}
			pnfsFlags := EXCHGID4_FLAG_MASK_PNFS & c.EidFlags
			Assert(InSliceUint32(pnfsFlags, validCombinations),
				"Invalid conbination of pnfs flags")
			// TODO: request invalid combinations
		})

		It("Lokupp", func(){
			r := c.Pass(c.SequenceArgs(),
				Putfh(globalDirFH),	Lookupp(), Getfh())
			Assert(AreFhEqual(rootFH, LastRes(&r).Opgetfh.Resok4.Object),
				"Parent fh is not root")
		})

		It("Save/Restore FH", func(){
			args := []NfsArgop4{c.SequenceArgs(), Putfh(rootFH), Putfh(rootFH)}
			maxOps := int(c.ForeChAttr.CaMaxoperations)
			for i:=0; i<(maxOps-4)/2; i++ {
				args = append(args, Savefh(), Restorefh())
			}
			args = append(args, Getfh())
			r := c.Pass(args...)
			Assert(AreFhEqual(rootFH, LastRes(&r).Opgetfh.Resok4.Object), "Filehandle should be the same")
			By("Repeat with too many ops")
			args[0] = c.SequenceArgs()
			args = append(args, Putrootfh(), Savefh(), Restorefh())
			c.Fail(NFS4ERR_TOO_MANY_OPS, args...)
		})

		It("BUG: deadlock secinfo+readdir compound", func(){
			openArgs := c.OpenArgs()
			name := openArgs.Opopen.Claim.File
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				openArgs,
			)
			c.Fail(NFS4ERR_NOFILEHANDLE,
				c.SequenceArgs(),
				Putfh(rootFH),
				Secinfo(name),
				Readdir(0, Verifier4{}, 4096, 8192,
					[]uint32{MakeGetAttrFlags(FATTR4_SIZE)}))
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(name))
		})

		It("Readdir", func(){
			_ = c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Readdir(0, Verifier4{}, 4096, 8192,
					[]uint32{MakeGetAttrFlags(FATTR4_SIZE)}))
		})

		It("Rename file (PyNFS::RNM1r)", func(){
			openArgs := c.OpenArgs()
			oldName := openArgs.Opopen.Claim.File
			newName := RandString(12)
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				openArgs,
			)
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Savefh(),
				Putfh(rootFH),
				Rename(oldName, newName))
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Lookup(newName))
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(newName))
		})

		It("Rename dir (PyNFS::RNM1d,RM1d)", func(){
			// TODO: other dir
			oldName := RandString(12)
			newName := RandString(12)
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				c.CreateArgs(oldName))
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Savefh(),
				Putfh(rootFH),
				Rename(oldName, newName))
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Lookup(newName))
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(newName))
		})

		It("Soft Link (PyNFS::???,RM1a)", func(){
			By("Target file has not been created")
			fileName := RandString(12)
			linkName := RandString(12)
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Create(Createtype4{Type:NF4LNK, Linkdata:fileName}, linkName,
					Fattr4{Attrmask:GetBitmap(FATTR4_MODE),
                   		AttrVals: GetPermAttrList(0644)}))
			By("Ensure the link is present")
			r := c.Pass(
				c.SequenceArgs(), Putfh(rootFH), Lookup(linkName), Getfh())
			fh := LastRes(&r).Opgetfh.Resok4.Object
			r = c.Pass(c.SequenceArgs(), Putfh(fh), Readlink())
			Assert(fileName == LastRes(&r).Opreadlink.Resok4.Link,
				"Readlink should read the same string")
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(linkName))
		})

		It("Hard Link", func() {
			openArgs := c.OpenArgs()
			fileName := openArgs.Opopen.Claim.File
			linkName := RandString(12)
			r := c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				openArgs,
				Getfh())
			fh := r[3].Opgetfh.Resok4.Object
			By("Create link to file")
			c.Pass(c.SequenceArgs(),
				Putfh(fh),
				Savefh(),
				Putfh(rootFH),
				Link(linkName))
			By("Check link is present")
			c.Pass(c.SequenceArgs(), Putfh(rootFH),	Lookup(linkName))
			By("Check file is present")
			c.Pass(c.SequenceArgs(), Putfh(rootFH),	Lookup(fileName))
			By("Clean up")
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(linkName))
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(fileName))
		})

		It("TODO: Access", func(){
			absentMask := MakeUint32Flags(ACCESS4_DELETE,
				ACCESS4_EXTEND, ACCESS4_LOOKUP, ACCESS4_MODIFY)
			r := c.Pass(
				c.SequenceArgs(),
				Putrootfh(),
				Access(MakeUint32Flags(ACCESS4_READ, ACCESS4_EXECUTE)),
			)
			res := LastRes(&r).Opaccess.Resok4
			By("ensure not asked bits are not set")
			Assert(0 == (res.Supported & absentMask),
				"Wrong Supported mask")
			Assert(0 == (res.Access & absentMask),
				"Wrong Access mask")
			//if 1 == res.Supported && uint32(ACCESS4_READ) {
			//}
		})

		It("TODO: bad session error", func(){
			badSid := [NFS4_SESSIONID_SIZE]byte{}
			for i:=range badSid {
				badSid[i] = byte(rand.Uint32())
			}
			c.Fail(
				NFS4ERR_BADSESSION,
				Sequence(badSid, c.Seq, 0, 0, false),
				Putrootfh(), Getfh())
		})

		It("Open/Write/Close/Open/Read", func(){
			content := RandString(128)
			openArgs := c.OpenArgs()
			fileName := openArgs.Opopen.Claim.File
			r := c.Pass(c.SequenceArgs(), Putfh(rootFH), openArgs, Getfh())
			resok := r[2].Opopen.Resok4
			stateId := resok.Stateid
			// TODO: CB_NOTIFY_LOCK
			//if ! CheckFlag(resok.Rflags, OPEN4_RESULT_MAY_NOTIFY_LOCK) {
			//	fmt.Println("TODO: Server does not Notify Lock Call")
			//}
			fh := LastRes(&r).Opgetfh.Resok4.Object
			By("Write to file")
			r = c.Pass(
				c.SequenceArgs(), Putfh(fh),
				Write(stateId, 0, UNSTABLE4, []byte(content)))
			By("Close file")
			c.Pass(c.SequenceArgs(), Putfh(fh), Close(c.Seq, stateId))
			By("Open again")
			openArgs = c.OpenNoCreateArgs()
			openArgs.Opopen.Claim.File = fileName
			r = c.Pass(c.SequenceArgs(), Putfh(rootFH), openArgs)
			stateId = LastRes(&r).Opopen.Resok4.Stateid
			By("Read file")
			r = c.Pass(c.SequenceArgs(), Putfh(fh), Read(stateId,0,128))
			Assert(content == string(LastRes(&r).Opread.Resok4.Data),
				"Data read is not the same")
			c.Pass(c.SequenceArgs(), Putfh(fh), Close(c.Seq, stateId))
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(fileName))
		})

        It("Lock/Test/Unlock (PyNFS::LOCK1,LKU1)", func() {
			openArgs := c.OpenArgs()
			fileName := openArgs.Opopen.Claim.File
            r := c.Pass(
				c.SequenceArgs(),
				Putfh(rootFH),
				openArgs, Getfh())
			resok := LastRes(&r).Opgetfh.Resok4
			fh := resok.Object
			sid := r[2].Opopen.Resok4.Stateid
            c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				c.LockArgs(sid))
            c.Fail(
				NFS4ERR_DENIED,
				c.SequenceArgs(),
				Putfh(fh),
				c.LocktArgs("Other owner"))
            c.Pass(c.SequenceArgs(),
				Putfh(fh),
				c.LocktArgs(c.Id))
			c.Pass(c.SequenceArgs(),
				Putfh(fh), c.LockuArgs(sid))
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(fileName))
        })

        It("TODO: Lock/Unlock in one compound", func() {
			openArgs := c.OpenArgs()
			fileName := openArgs.Opopen.Claim.File
            r := c.Pass(
				c.SequenceArgs(),
				Putfh(rootFH),
				openArgs, Getfh())
			resok := LastRes(&r).Opgetfh.Resok4
			fh := resok.Object
			sid := r[2].Opopen.Resok4.Stateid
            c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				c.LockArgs(sid),
				Putfh(fh),
				c.LocktArgs(c.Id),
				Putfh(fh),
				c.LockuArgs(sid))
			// TODO:
			//sid = LastRes(&r).Oplocku.LockStateid
            //c.Fail(NFS4ERR_OLD_STATEID,
				//c.SequenceArgs(),
				//Putfh(fh),
				//c.LockArgs(sid),
				//Putfh(fh),
				//c.LocktArgs(c.Id),
				//Putfh(fh),
				//c.LockuArgs(sid),
				//Putfh(fh),
				//c.LockArgs(sid),
				//Putfh(fh),
				//c.LocktArgs(c.Id),
				//Putfh(fh),
				//c.LockuArgs(sid))
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(fileName))
        })

		It("DestroySession", func(){
			newClient := NewNFSv41DefaultClient()
			newClient.ExchangeId()
			newClient.CreateSession()
			newClient.GetSomeAttr()
			By("Destroing session...")
			newClient.Pass(newClient.SequenceArgs(),
				DestroySession(newClient.Sid))
			By("Trying to use destroyed...")
			newClient.Fail(NFS4ERR_BADSESSION, newClient.SequenceArgs())
			newClient.Pass(DestroyClientid(newClient.ClientId))
		})

	})

	Context("Slow", func() {

        It("Session Expiration, Lock Release", func() {
			cliExpiring := NewNFSv41DefaultClient()
			cliExpiring.ExchangeId()
			cliExpiring.CreateSession()
			cliExpiring.GetSomeAttr()

            r := cliExpiring.Pass(
				cliExpiring.SequenceArgs(),
				Putfh(rootFH),
				cliExpiring.OpenArgs(), Getfh())
			resok := LastRes(&r).Opgetfh.Resok4
			fh := resok.Object
			sid := r[2].Opopen.Resok4.Stateid
            cliExpiring.Pass(
				cliExpiring.SequenceArgs(),
				Putfh(fh),
				cliExpiring.LockArgs(sid))

			By("pinging server in default client and abandon in cliExpiring")
			// supposing LeaseTime is the same
			interval := time.Second * time.Duration(c.LeaseTime / 6)
			for i:=0;i<7;i++ {
				time.Sleep(interval)
				c.Pass(c.SequenceArgs())
			}

            c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				c.LocktArgs("Other Owner"))

            cliExpiring.Fail(
				NFS4ERR_BADSESSION,
				cliExpiring.SequenceArgs(),
			)
			cliExpiring.Close()
        })

		It("CreateSession Timeout (PyNFS::EID9)", func() {
			cliStale := NewNFSv41DefaultClient()

			time.Sleep(time.Second * time.Duration(c.LeaseTime + 5))

			cliStale.Fail(
				NFS4ERR_STALE_CLIENTID,
				CreateSession(
					cliStale.ClientId,
					cliStale.Seq,
					DefCsFlags,
					DefChannelAttrs,
					DefChannelAttrs,
					0x40000000,
					[]CallbackSecParms4{{
						CbSecflavor:1,
						CbspSysCred:cliStale.AuthSys}}))
			cliStale.Close()
		})

	})

})
