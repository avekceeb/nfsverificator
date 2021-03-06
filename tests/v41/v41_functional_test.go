package v41tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/common"
	"time"
	"math/rand"
	"fmt"
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

		It("Refer Option", func() {
			// TODO !!!!
			//if "" == Config.GetRefPath() {
			Skip("No ref path exposed by server")
			//}
			if ! CheckFlag(c.EidFlags, EXCHGID4_FLAG_SUPP_MOVED_REFER) {
				Skip("Server does not support EXCHGID4_FLAG_SUPP_MOVED_REFER")
			}
			// TODO: put refer path to config
			c.Fail(NFS4ERR_MOVED,
				c.SequenceArgs(),
				Putfh(rootFH),
				Lookup("ref"),
				Getfh())
			//fh := GrabFh(&r)
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Lookup("ref"),
				Getattr(MakeGetAttrFlags(
					FATTR4_FS_LOCATIONS)))
			// ??? FATTR4_FS_LOCATIONS_INFO,
		})

		It("ExchangeId flags combinations RFC 5661 13.1", func() {
			validCombinations := []uint32{
				uint32(EXCHGID4_FLAG_USE_PNFS_MDS),
				uint32(EXCHGID4_FLAG_USE_PNFS_DS),
				uint32(EXCHGID4_FLAG_USE_NON_PNFS),
				MakeUint32Flags(EXCHGID4_FLAG_USE_PNFS_MDS,
					EXCHGID4_FLAG_USE_PNFS_DS),
				MakeUint32Flags(EXCHGID4_FLAG_USE_PNFS_DS,
					EXCHGID4_FLAG_USE_NON_PNFS)}
			pnfsFlags := EXCHGID4_FLAG_MASK_PNFS & c.EidFlags
			Assert41.Assert(InSliceUint32(pnfsFlags, validCombinations),
				"Invalid conbination of pnfs flags")
			// TODO: request invalid combinations
		})

		It("Getattr", func(){
			// TODO: why these commented attr dont work ?
			allAttrs := []int{
				FATTR4_TYPE, FATTR4_FH_EXPIRE_TYPE,
				FATTR4_CHANGE, FATTR4_SIZE, FATTR4_LINK_SUPPORT,
				FATTR4_SYMLINK_SUPPORT, FATTR4_NAMED_ATTR, FATTR4_FSID,
				FATTR4_UNIQUE_HANDLES, FATTR4_LEASE_TIME, FATTR4_RDATTR_ERROR,
				FATTR4_FILEHANDLE, FATTR4_SUPPATTR_EXCLCREAT, FATTR4_ACL,
				FATTR4_ACLSUPPORT, FATTR4_ARCHIVE, FATTR4_CANSETTIME,
				FATTR4_CASE_INSENSITIVE, FATTR4_CASE_PRESERVING,
				FATTR4_CHOWN_RESTRICTED,
				FATTR4_FILEID, FATTR4_FILES_AVAIL, FATTR4_FILES_FREE,
				FATTR4_FILES_TOTAL, FATTR4_FS_LOCATIONS, FATTR4_HIDDEN,
				FATTR4_HOMOGENEOUS, FATTR4_MAXFILESIZE, FATTR4_MAXLINK,
				FATTR4_MAXNAME, FATTR4_MAXREAD, FATTR4_MAXWRITE, FATTR4_MIMETYPE,
				FATTR4_MODE, FATTR4_NO_TRUNC, FATTR4_NUMLINKS,
				FATTR4_OWNER, FATTR4_OWNER_GROUP, FATTR4_QUOTA_AVAIL_HARD,
				FATTR4_QUOTA_AVAIL_SOFT,
				FATTR4_QUOTA_USED, FATTR4_RAWDEV, FATTR4_SPACE_AVAIL,
				FATTR4_SPACE_FREE, FATTR4_SPACE_TOTAL,
				FATTR4_SPACE_USED, FATTR4_SYSTEM, FATTR4_TIME_ACCESS,
				/*FATTR4_TIME_ACCESS_SET, */FATTR4_TIME_BACKUP, FATTR4_TIME_CREATE,
				FATTR4_TIME_DELTA, FATTR4_TIME_METADATA, FATTR4_TIME_MODIFY,
				/*FATTR4_TIME_MODIFY_SET,*/ FATTR4_MOUNTED_ON_FILEID,
				FATTR4_DIR_NOTIF_DELAY, FATTR4_DIRENT_NOTIF_DELAY,
				FATTR4_DACL, FATTR4_SACL, FATTR4_CHANGE_POLICY,
				FATTR4_FS_STATUS, FATTR4_FS_LAYOUT_TYPES,
				/*FATTR4_LAYOUT_HINT,*/ FATTR4_LAYOUT_TYPES,
				FATTR4_LAYOUT_BLKSIZE, FATTR4_LAYOUT_ALIGNMENT,
				FATTR4_FS_LOCATIONS_INFO, FATTR4_MDSTHRESHOLD,
				FATTR4_RETENTION_GET, /*FATTR4_RETENTION_SET,*/
				FATTR4_RETENTEVT_GET,
				/*FATTR4_RETENTEVT_SET,*/ FATTR4_RETENTION_HOLD,
				/*FATTR4_MODE_SET_MASKED,*/ FATTR4_FS_CHARSET_CAP,
			}
			c.Compound(c.SequenceArgs(),
				Putfh(globalFileFH),
				Getattr([]uint32{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}))
			c.Pass(c.SequenceArgs(),
				Putfh(globalDirFH),
				Getattr(MakeGetAttrFlags(allAttrs...)))
			for _,v := range allAttrs {
				By(fmt.Sprintf("Trying Getattr with single %d attr", v))
				c.Pass(c.SequenceArgs(),
					Putfh(globalDirFH),
					Getattr(MakeGetAttrFlags(v)))
			}
		})

		It("NVerify (PyNFS::NVF1r,NVF2r)", func(){
			r := c.Pass(c.SequenceArgs(),
				Putfh(globalFileFH),
				Getattr(MakeGetAttrFlags(
					FATTR4_SUPPORTED_ATTRS,
					FATTR4_TYPE, FATTR4_FH_EXPIRE_TYPE,
					FATTR4_CHANGE, FATTR4_SIZE, FATTR4_LINK_SUPPORT,
					FATTR4_NAMED_ATTR, FATTR4_FSID, FATTR4_UNIQUE_HANDLES,
					FATTR4_LEASE_TIME, FATTR4_FILEHANDLE)))
			attrs := LastRes(&r).Opgetattr.Resok4.ObjAttributes
			By("Nverify should return `Same` error for pre-read attrs")
			c.Fail(NFS4ERR_SAME,
				c.SequenceArgs(),
				Putfh(globalFileFH),
				Nverify(attrs),
				Lookup(RandString(5)), Lookup(RandString(5)))
			attrFileType := Fattr4{Attrmask:MakeGetAttrFlags(FATTR4_TYPE),
				AttrVals:[]byte{0,0,0,byte(NF4REG)}}
			By("Nverify should return `Same` error for one attr type")
			c.Fail(NFS4ERR_SAME,
				c.SequenceArgs(),
				Putfh(globalFileFH),
				Nverify(attrFileType),
				Lookup(RandString(5)), Lookup(RandString(5)))
			By("Nverify should return OK because attr dir != attr file")
			c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Nverify(attrFileType),
				Lookup(globalFile))
		})

		It("Lookupp", func(){
			r := c.Pass(c.SequenceArgs(),
				Putfh(globalDirFH),	Lookupp(), Getfh())
			Assert41.Assert(AreFhEqual(rootFH, GrabFh(&r)),
				"Parent fh is not root")
		})

		It("Save/Restore FH", func(){
			args := []NfsArgop4{c.SequenceArgs(), Putfh(rootFH), Putfh(rootFH)}
			maxOps := int(c.ForeChAttr.CaMaxoperations)
			for i:=0; i<(maxOps-4)/2; i++ {
				args = append(args, Savefh(), Restorefh())
			}
			args = append(args, Getfh())
			By("Many save/restore")
			r := c.Pass(args...)
			Assert41.Assert(AreFhEqual(rootFH, LastRes(&r).Opgetfh.Resok4.Object),
				"Filehandle should be the same")
			By("Several ops between save and restore")
			args = []NfsArgop4{c.SequenceArgs(), Putfh(rootFH), Savefh(), Putfh(globalDirFH)}
			for i:=0; i<(maxOps-6)/2; i++ {
				args = append(args, Putfh(rootFH), Lookup(globalFile))
			}
			args = append(args, Restorefh(), Getfh())
			r = c.Pass(args...)
			Assert41.Assert(AreFhEqual(rootFH, LastRes(&r).Opgetfh.Resok4.Object),
				"Filehandle should be the same")
			By("Repeat with too many ops. TODO: RESOURCE ERR???")
			args[0] = c.SequenceArgs()
			args = append(args, Putrootfh(), Savefh(), Restorefh())
			c.FailOneOf([]int32{NFS4ERR_TOO_MANY_OPS, NFS4ERR_RESOURCE},
				args...)
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
					MakeGetAttrFlags(FATTR4_SIZE)))
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(name))
		})

		It("Readdir", func(){
			_ = c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				Readdir(0, Verifier4{}, 4096, 8192,
					MakeGetAttrFlags(FATTR4_SIZE)))
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
			Assert41.Assert(fileName == LastRes(&r).Opreadlink.Resok4.Link,
				"Readlink should read the same string")
			c.Pass(c.SequenceArgs(), Putfh(rootFH), Remove(linkName))
		})

		It("Hard Link RFC 5661 4.2.1", func() {
			openArgs := c.OpenArgs()
			fileName := openArgs.Opopen.Claim.File
			linkName := RandString(12)
			r := c.Pass(c.SequenceArgs(),
				Putfh(rootFH),
				openArgs,
				Getfh())
			fh := GrabFh(&r)
			By("Create link to file")
			c.Pass(c.SequenceArgs(),
				Putfh(fh),
				Savefh(),
				Putfh(rootFH),
				Link(linkName))
			By("Check link is present")
			r = c.Pass(c.SequenceArgs(), Putfh(rootFH), Lookup(linkName), Getfh())
			fhL := GrabFh(&r)
			By("Check file is present")
			r = c.Pass(c.SequenceArgs(), Putfh(rootFH), Lookup(fileName), Getfh())
			fhF := GrabFh(&r)
			By("Check that fh is the same")
			Assert41.Assert(AreFhEqual(fhL, fhF), "fh are different!")
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
			Assert41.Assert(0 == (res.Supported & absentMask),
				"Wrong Supported mask")
			Assert41.Assert(0 == (res.Access & absentMask),
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
			fh := GrabFh(&r)
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
			// TODO: get filesize
			r = c.Pass(c.SequenceArgs(), Putfh(fh), Read(stateId,0,128))
			Assert41.Assert(content == string(LastRes(&r).Opread.Resok4.Data),
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
			r = c.Pass(
				c.SequenceArgs(),
				Putfh(fh),
				c.LockArgs(sid))
			// ganesha does not require new sid to unlock
			// but linux does
			sid = LastRes(&r).Oplock.Resok4.LockStateid
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
			Skip("TODO: diffrent results on linux vs ganesha: NFS4ERR_BAD_STATEID vs pass")
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

		It("Close Expired", func() {
			// TODO: got NFS4ERR_BAD_STATEID
			cli2 := NewNFSv41DefaultClient()
			cli2.ExchangeId()
			cli2.CreateSession()
			openArgs := cli2.OpenNoCreateArgs()
			openArgs.Opopen.Claim.File = globalFile
			r := cli2.Pass(cli2.SequenceArgs(), Putfh(rootFH), openArgs)
			stateId := LastRes(&r).Opopen.Resok4.Stateid
			interval := time.Second * time.Duration(c.LeaseTime / 6)
			for i:=0;i<7;i++ {
				time.Sleep(interval)
				c.Pass(c.SequenceArgs())
			}
			c.Fail(NFS4ERR_EXPIRED,
				c.SequenceArgs(), Putfh(globalFileFH), Close(c.Seq, stateId))
		})

	})

})
