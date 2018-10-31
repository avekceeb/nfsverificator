package v40

const (
	// Enums:
	NF4REG = 1 // NF4REG
	NF4DIR = 2 // NF4DIR
	NF4BLK = 3 // NF4BLK
	NF4CHR = 4 // NF4CHR
	NF4LNK = 5 // NF4LNK
	NF4SOCK = 6 // NF4SOCK
	NF4FIFO = 7 // NF4FIFO
	NF4ATTRDIR = 8 // NF4ATTRDIR
	NF4NAMEDATTR = 9 // NF4NAMEDATTR
	NFS4_OK = 0 // NFS4_OK
	NFS4ERR_PERM = 1 // NFS4ERR_PERM
	NFS4ERR_NOENT = 2 // NFS4ERR_NOENT
	NFS4ERR_IO = 5 // NFS4ERR_IO
	NFS4ERR_NXIO = 6 // NFS4ERR_NXIO
	NFS4ERR_ACCESS = 13 // NFS4ERR_ACCESS
	NFS4ERR_EXIST = 17 // NFS4ERR_EXIST
	NFS4ERR_XDEV = 18 // NFS4ERR_XDEV
	NFS4ERR_NOTDIR = 20 // NFS4ERR_NOTDIR
	NFS4ERR_ISDIR = 21 // NFS4ERR_ISDIR
	NFS4ERR_INVAL = 22 // NFS4ERR_INVAL
	NFS4ERR_FBIG = 27 // NFS4ERR_FBIG
	NFS4ERR_NOSPC = 28 // NFS4ERR_NOSPC
	NFS4ERR_ROFS = 30 // NFS4ERR_ROFS
	NFS4ERR_MLINK = 31 // NFS4ERR_MLINK
	NFS4ERR_NAMETOOLONG = 63 // NFS4ERR_NAMETOOLONG
	NFS4ERR_NOTEMPTY = 66 // NFS4ERR_NOTEMPTY
	NFS4ERR_DQUOT = 69 // NFS4ERR_DQUOT
	NFS4ERR_STALE = 70 // NFS4ERR_STALE
	NFS4ERR_BADHANDLE = 10001 // NFS4ERR_BADHANDLE
	NFS4ERR_BAD_COOKIE = 10003 // NFS4ERR_BAD_COOKIE
	NFS4ERR_NOTSUPP = 10004 // NFS4ERR_NOTSUPP
	NFS4ERR_TOOSMALL = 10005 // NFS4ERR_TOOSMALL
	NFS4ERR_SERVERFAULT = 10006 // NFS4ERR_SERVERFAULT
	NFS4ERR_BADTYPE = 10007 // NFS4ERR_BADTYPE
	NFS4ERR_DELAY = 10008 // NFS4ERR_DELAY
	NFS4ERR_SAME = 10009 // NFS4ERR_SAME
	NFS4ERR_DENIED = 10010 // NFS4ERR_DENIED
	NFS4ERR_EXPIRED = 10011 // NFS4ERR_EXPIRED
	NFS4ERR_LOCKED = 10012 // NFS4ERR_LOCKED
	NFS4ERR_GRACE = 10013 // NFS4ERR_GRACE
	NFS4ERR_FHEXPIRED = 10014 // NFS4ERR_FHEXPIRED
	NFS4ERR_SHARE_DENIED = 10015 // NFS4ERR_SHARE_DENIED
	NFS4ERR_WRONGSEC = 10016 // NFS4ERR_WRONGSEC
	NFS4ERR_CLID_INUSE = 10017 // NFS4ERR_CLID_INUSE
	NFS4ERR_RESOURCE = 10018 // NFS4ERR_RESOURCE
	NFS4ERR_MOVED = 10019 // NFS4ERR_MOVED
	NFS4ERR_NOFILEHANDLE = 10020 // NFS4ERR_NOFILEHANDLE
	NFS4ERR_MINOR_VERS_MISMATCH = 10021 // NFS4ERR_MINOR_VERS_MISMATCH
	NFS4ERR_STALE_CLIENTID = 10022 // NFS4ERR_STALE_CLIENTID
	NFS4ERR_STALE_STATEID = 10023 // NFS4ERR_STALE_STATEID
	NFS4ERR_OLD_STATEID = 10024 // NFS4ERR_OLD_STATEID
	NFS4ERR_BAD_STATEID = 10025 // NFS4ERR_BAD_STATEID
	NFS4ERR_BAD_SEQID = 10026 // NFS4ERR_BAD_SEQID
	NFS4ERR_NOT_SAME = 10027 // NFS4ERR_NOT_SAME
	NFS4ERR_LOCK_RANGE = 10028 // NFS4ERR_LOCK_RANGE
	NFS4ERR_SYMLINK = 10029 // NFS4ERR_SYMLINK
	NFS4ERR_RESTOREFH = 10030 // NFS4ERR_RESTOREFH
	NFS4ERR_LEASE_MOVED = 10031 // NFS4ERR_LEASE_MOVED
	NFS4ERR_ATTRNOTSUPP = 10032 // NFS4ERR_ATTRNOTSUPP
	NFS4ERR_NO_GRACE = 10033 // NFS4ERR_NO_GRACE
	NFS4ERR_RECLAIM_BAD = 10034 // NFS4ERR_RECLAIM_BAD
	NFS4ERR_RECLAIM_CONFLICT = 10035 // NFS4ERR_RECLAIM_CONFLICT
	NFS4ERR_BADXDR = 10036 // NFS4ERR_BADXDR
	NFS4ERR_LOCKS_HELD = 10037 // NFS4ERR_LOCKS_HELD
	NFS4ERR_OPENMODE = 10038 // NFS4ERR_OPENMODE
	NFS4ERR_BADOWNER = 10039 // NFS4ERR_BADOWNER
	NFS4ERR_BADCHAR = 10040 // NFS4ERR_BADCHAR
	NFS4ERR_BADNAME = 10041 // NFS4ERR_BADNAME
	NFS4ERR_BAD_RANGE = 10042 // NFS4ERR_BAD_RANGE
	NFS4ERR_LOCK_NOTSUPP = 10043 // NFS4ERR_LOCK_NOTSUPP
	NFS4ERR_OP_ILLEGAL = 10044 // NFS4ERR_OP_ILLEGAL
	NFS4ERR_DEADLOCK = 10045 // NFS4ERR_DEADLOCK
	NFS4ERR_FILE_OPEN = 10046 // NFS4ERR_FILE_OPEN
	NFS4ERR_ADMIN_REVOKED = 10047 // NFS4ERR_ADMIN_REVOKED
	NFS4ERR_CB_PATH_DOWN = 10048 // NFS4ERR_CB_PATH_DOWN
	SET_TO_SERVER_TIME4 = 0 // SET_TO_SERVER_TIME4
	SET_TO_CLIENT_TIME4 = 1 // SET_TO_CLIENT_TIME4
	READ_LT = 1 // READ_LT
	WRITE_LT = 2 // WRITE_LT
	READW_LT = 3 // READW_LT
	WRITEW_LT = 4 // WRITEW_LT
	UNCHECKED4 = 0 // UNCHECKED4
	GUARDED4 = 1 // GUARDED4
	EXCLUSIVE4 = 2 // EXCLUSIVE4
	OPEN4_NOCREATE = 0 // OPEN4_NOCREATE
	OPEN4_CREATE = 1 // OPEN4_CREATE
	NFS_LIMIT_SIZE = 1 // NFS_LIMIT_SIZE
	NFS_LIMIT_BLOCKS = 2 // NFS_LIMIT_BLOCKS
	OPEN_DELEGATE_NONE = 0 // OPEN_DELEGATE_NONE
	OPEN_DELEGATE_READ = 1 // OPEN_DELEGATE_READ
	OPEN_DELEGATE_WRITE = 2 // OPEN_DELEGATE_WRITE
	CLAIM_NULL = 0 // CLAIM_NULL
	CLAIM_PREVIOUS = 1 // CLAIM_PREVIOUS
	CLAIM_DELEGATE_CUR = 2 // CLAIM_DELEGATE_CUR
	CLAIM_DELEGATE_PREV = 3 // CLAIM_DELEGATE_PREV
	RPC_GSS_SVC_NONE = 1 // RPC_GSS_SVC_NONE
	RPC_GSS_SVC_INTEGRITY = 2 // RPC_GSS_SVC_INTEGRITY
	RPC_GSS_SVC_PRIVACY = 3 // RPC_GSS_SVC_PRIVACY
	UNSTABLE4 = 0 // UNSTABLE4
	DATA_SYNC4 = 1 // DATA_SYNC4
	FILE_SYNC4 = 2 // FILE_SYNC4
	OP_ACCESS = 3 // OP_ACCESS
	OP_CLOSE = 4 // OP_CLOSE
	OP_COMMIT = 5 // OP_COMMIT
	OP_CREATE = 6 // OP_CREATE
	OP_DELEGPURGE = 7 // OP_DELEGPURGE
	OP_DELEGRETURN = 8 // OP_DELEGRETURN
	OP_GETATTR = 9 // OP_GETATTR
	OP_GETFH = 10 // OP_GETFH
	OP_LINK = 11 // OP_LINK
	OP_LOCK = 12 // OP_LOCK
	OP_LOCKT = 13 // OP_LOCKT
	OP_LOCKU = 14 // OP_LOCKU
	OP_LOOKUP = 15 // OP_LOOKUP
	OP_LOOKUPP = 16 // OP_LOOKUPP
	OP_NVERIFY = 17 // OP_NVERIFY
	OP_OPEN = 18 // OP_OPEN
	OP_OPENATTR = 19 // OP_OPENATTR
	OP_OPEN_CONFIRM = 20 // OP_OPEN_CONFIRM
	OP_OPEN_DOWNGRADE = 21 // OP_OPEN_DOWNGRADE
	OP_PUTFH = 22 // OP_PUTFH
	OP_PUTPUBFH = 23 // OP_PUTPUBFH
	OP_PUTROOTFH = 24 // OP_PUTROOTFH
	OP_READ = 25 // OP_READ
	OP_READDIR = 26 // OP_READDIR
	OP_READLINK = 27 // OP_READLINK
	OP_REMOVE = 28 // OP_REMOVE
	OP_RENAME = 29 // OP_RENAME
	OP_RENEW = 30 // OP_RENEW
	OP_RESTOREFH = 31 // OP_RESTOREFH
	OP_SAVEFH = 32 // OP_SAVEFH
	OP_SECINFO = 33 // OP_SECINFO
	OP_SETATTR = 34 // OP_SETATTR
	OP_SETCLIENTID = 35 // OP_SETCLIENTID
	OP_SETCLIENTID_CONFIRM = 36 // OP_SETCLIENTID_CONFIRM
	OP_VERIFY = 37 // OP_VERIFY
	OP_WRITE = 38 // OP_WRITE
	OP_RELEASE_LOCKOWNER = 39 // OP_RELEASE_LOCKOWNER
	OP_ILLEGAL = 10044 // OP_ILLEGAL
	OP_CB_GETATTR = 3 // OP_CB_GETATTR
	OP_CB_RECALL = 4 // OP_CB_RECALL
	OP_CB_ILLEGAL = 10044 // OP_CB_ILLEGAL

	// Consts:
	RPCSEC_GSS = 6 // RPCSEC_GSS
	FALSE = 0 // FALSE
	TRUE = 1 // TRUE
	NFS4_PROGRAM = 100003 // NFS4_PROGRAM
	NFS_V4 = 4 // NFS_V4
	NFSPROC4_NULL = 0 // NFSPROC4_NULL
	NFSPROC4_COMPOUND = 1 // NFSPROC4_COMPOUND
	NFS4_CALLBACK = 0x40000000 // NFS4_CALLBACK
	NFS_CB = 1 // NFS_CB
	NFS4_FHSIZE = 128 // NFS4_FHSIZE
	NFS4_VERIFIER_SIZE = 8 // NFS4_VERIFIER_SIZE
	NFS4_OTHER_SIZE = 12 // NFS4_OTHER_SIZE
	NFS4_OPAQUE_LIMIT = 1024 // NFS4_OPAQUE_LIMIT
	NFS4_INT64_MAX = 0x7fffffffffffffff // NFS4_INT64_MAX
	NFS4_UINT64_MAX = 0xffffffffffffffff // NFS4_UINT64_MAX
	NFS4_INT32_MAX = 0x7fffffff // NFS4_INT32_MAX
	NFS4_UINT32_MAX = 0xffffffff // NFS4_UINT32_MAX
	ACL4_SUPPORT_ALLOW_ACL = 0x00000001 // ACL4_SUPPORT_ALLOW_ACL
	ACL4_SUPPORT_DENY_ACL = 0x00000002 // ACL4_SUPPORT_DENY_ACL
	ACL4_SUPPORT_AUDIT_ACL = 0x00000004 // ACL4_SUPPORT_AUDIT_ACL
	ACL4_SUPPORT_ALARM_ACL = 0x00000008 // ACL4_SUPPORT_ALARM_ACL
	ACE4_ACCESS_ALLOWED_ACE_TYPE = 0x00000000 // ACE4_ACCESS_ALLOWED_ACE_TYPE
	ACE4_ACCESS_DENIED_ACE_TYPE = 0x00000001 // ACE4_ACCESS_DENIED_ACE_TYPE
	ACE4_SYSTEM_AUDIT_ACE_TYPE = 0x00000002 // ACE4_SYSTEM_AUDIT_ACE_TYPE
	ACE4_SYSTEM_ALARM_ACE_TYPE = 0x00000003 // ACE4_SYSTEM_ALARM_ACE_TYPE
	ACE4_FILE_INHERIT_ACE = 0x00000001 // ACE4_FILE_INHERIT_ACE
	ACE4_DIRECTORY_INHERIT_ACE = 0x00000002 // ACE4_DIRECTORY_INHERIT_ACE
	ACE4_NO_PROPAGATE_INHERIT_ACE = 0x00000004 // ACE4_NO_PROPAGATE_INHERIT_ACE
	ACE4_INHERIT_ONLY_ACE = 0x00000008 // ACE4_INHERIT_ONLY_ACE
	ACE4_SUCCESSFUL_ACCESS_ACE_FLAG = 0x00000010 // ACE4_SUCCESSFUL_ACCESS_ACE_FLAG
	ACE4_FAILED_ACCESS_ACE_FLAG = 0x00000020 // ACE4_FAILED_ACCESS_ACE_FLAG
	ACE4_IDENTIFIER_GROUP = 0x00000040 // ACE4_IDENTIFIER_GROUP
	ACE4_READ_DATA = 0x00000001 // ACE4_READ_DATA
	ACE4_LIST_DIRECTORY = 0x00000001 // ACE4_LIST_DIRECTORY
	ACE4_WRITE_DATA = 0x00000002 // ACE4_WRITE_DATA
	ACE4_ADD_FILE = 0x00000002 // ACE4_ADD_FILE
	ACE4_APPEND_DATA = 0x00000004 // ACE4_APPEND_DATA
	ACE4_ADD_SUBDIRECTORY = 0x00000004 // ACE4_ADD_SUBDIRECTORY
	ACE4_READ_NAMED_ATTRS = 0x00000008 // ACE4_READ_NAMED_ATTRS
	ACE4_WRITE_NAMED_ATTRS = 0x00000010 // ACE4_WRITE_NAMED_ATTRS
	ACE4_EXECUTE = 0x00000020 // ACE4_EXECUTE
	ACE4_DELETE_CHILD = 0x00000040 // ACE4_DELETE_CHILD
	ACE4_READ_ATTRIBUTES = 0x00000080 // ACE4_READ_ATTRIBUTES
	ACE4_WRITE_ATTRIBUTES = 0x00000100 // ACE4_WRITE_ATTRIBUTES
	ACE4_DELETE = 0x00010000 // ACE4_DELETE
	ACE4_READ_ACL = 0x00020000 // ACE4_READ_ACL
	ACE4_WRITE_ACL = 0x00040000 // ACE4_WRITE_ACL
	ACE4_WRITE_OWNER = 0x00080000 // ACE4_WRITE_OWNER
	ACE4_SYNCHRONIZE = 0x00100000 // ACE4_SYNCHRONIZE
	ACE4_GENERIC_READ = 0x00120081 // ACE4_GENERIC_READ
	ACE4_GENERIC_WRITE = 0x00160106 // ACE4_GENERIC_WRITE
	ACE4_GENERIC_EXECUTE = 0x001200A0 // ACE4_GENERIC_EXECUTE
	MODE4_SUID = 0x800 // MODE4_SUID
	MODE4_SGID = 0x400 // MODE4_SGID
	MODE4_SVTX = 0x200 // MODE4_SVTX
	MODE4_RUSR = 0x100 // MODE4_RUSR
	MODE4_WUSR = 0x080 // MODE4_WUSR
	MODE4_XUSR = 0x040 // MODE4_XUSR
	MODE4_RGRP = 0x020 // MODE4_RGRP
	MODE4_WGRP = 0x010 // MODE4_WGRP
	MODE4_XGRP = 0x008 // MODE4_XGRP
	MODE4_ROTH = 0x004 // MODE4_ROTH
	MODE4_WOTH = 0x002 // MODE4_WOTH
	MODE4_XOTH = 0x001 // MODE4_XOTH
	FH4_PERSISTENT = 0x00000000 // FH4_PERSISTENT
	FH4_NOEXPIRE_WITH_OPEN = 0x00000001 // FH4_NOEXPIRE_WITH_OPEN
	FH4_VOLATILE_ANY = 0x00000002 // FH4_VOLATILE_ANY
	FH4_VOL_MIGRATION = 0x00000004 // FH4_VOL_MIGRATION
	FH4_VOL_RENAME = 0x00000008 // FH4_VOL_RENAME
	FATTR4_SUPPORTED_ATTRS = 0 // FATTR4_SUPPORTED_ATTRS
	FATTR4_TYPE = 1 // FATTR4_TYPE
	FATTR4_FH_EXPIRE_TYPE = 2 // FATTR4_FH_EXPIRE_TYPE
	FATTR4_CHANGE = 3 // FATTR4_CHANGE
	FATTR4_SIZE = 4 // FATTR4_SIZE
	FATTR4_LINK_SUPPORT = 5 // FATTR4_LINK_SUPPORT
	FATTR4_SYMLINK_SUPPORT = 6 // FATTR4_SYMLINK_SUPPORT
	FATTR4_NAMED_ATTR = 7 // FATTR4_NAMED_ATTR
	FATTR4_FSID = 8 // FATTR4_FSID
	FATTR4_UNIQUE_HANDLES = 9 // FATTR4_UNIQUE_HANDLES
	FATTR4_LEASE_TIME = 10 // FATTR4_LEASE_TIME
	FATTR4_RDATTR_ERROR = 11 // FATTR4_RDATTR_ERROR
	FATTR4_FILEHANDLE = 19 // FATTR4_FILEHANDLE
	FATTR4_ACL = 12 // FATTR4_ACL
	FATTR4_ACLSUPPORT = 13 // FATTR4_ACLSUPPORT
	FATTR4_ARCHIVE = 14 // FATTR4_ARCHIVE
	FATTR4_CANSETTIME = 15 // FATTR4_CANSETTIME
	FATTR4_CASE_INSENSITIVE = 16 // FATTR4_CASE_INSENSITIVE
	FATTR4_CASE_PRESERVING = 17 // FATTR4_CASE_PRESERVING
	FATTR4_CHOWN_RESTRICTED = 18 // FATTR4_CHOWN_RESTRICTED
	FATTR4_FILEID = 20 // FATTR4_FILEID
	FATTR4_FILES_AVAIL = 21 // FATTR4_FILES_AVAIL
	FATTR4_FILES_FREE = 22 // FATTR4_FILES_FREE
	FATTR4_FILES_TOTAL = 23 // FATTR4_FILES_TOTAL
	FATTR4_FS_LOCATIONS = 24 // FATTR4_FS_LOCATIONS
	FATTR4_HIDDEN = 25 // FATTR4_HIDDEN
	FATTR4_HOMOGENEOUS = 26 // FATTR4_HOMOGENEOUS
	FATTR4_MAXFILESIZE = 27 // FATTR4_MAXFILESIZE
	FATTR4_MAXLINK = 28 // FATTR4_MAXLINK
	FATTR4_MAXNAME = 29 // FATTR4_MAXNAME
	FATTR4_MAXREAD = 30 // FATTR4_MAXREAD
	FATTR4_MAXWRITE = 31 // FATTR4_MAXWRITE
	FATTR4_MIMETYPE = 32 // FATTR4_MIMETYPE
	FATTR4_MODE = 33 // FATTR4_MODE
	FATTR4_NO_TRUNC = 34 // FATTR4_NO_TRUNC
	FATTR4_NUMLINKS = 35 // FATTR4_NUMLINKS
	FATTR4_OWNER = 36 // FATTR4_OWNER
	FATTR4_OWNER_GROUP = 37 // FATTR4_OWNER_GROUP
	FATTR4_QUOTA_AVAIL_HARD = 38 // FATTR4_QUOTA_AVAIL_HARD
	FATTR4_QUOTA_AVAIL_SOFT = 39 // FATTR4_QUOTA_AVAIL_SOFT
	FATTR4_QUOTA_USED = 40 // FATTR4_QUOTA_USED
	FATTR4_RAWDEV = 41 // FATTR4_RAWDEV
	FATTR4_SPACE_AVAIL = 42 // FATTR4_SPACE_AVAIL
	FATTR4_SPACE_FREE = 43 // FATTR4_SPACE_FREE
	FATTR4_SPACE_TOTAL = 44 // FATTR4_SPACE_TOTAL
	FATTR4_SPACE_USED = 45 // FATTR4_SPACE_USED
	FATTR4_SYSTEM = 46 // FATTR4_SYSTEM
	FATTR4_TIME_ACCESS = 47 // FATTR4_TIME_ACCESS
	FATTR4_TIME_ACCESS_SET = 48 // FATTR4_TIME_ACCESS_SET
	FATTR4_TIME_BACKUP = 49 // FATTR4_TIME_BACKUP
	FATTR4_TIME_CREATE = 50 // FATTR4_TIME_CREATE
	FATTR4_TIME_DELTA = 51 // FATTR4_TIME_DELTA
	FATTR4_TIME_METADATA = 52 // FATTR4_TIME_METADATA
	FATTR4_TIME_MODIFY = 53 // FATTR4_TIME_MODIFY
	FATTR4_TIME_MODIFY_SET = 54 // FATTR4_TIME_MODIFY_SET
	FATTR4_MOUNTED_ON_FILEID = 55 // FATTR4_MOUNTED_ON_FILEID
	ACCESS4_READ = 0x00000001 // ACCESS4_READ
	ACCESS4_LOOKUP = 0x00000002 // ACCESS4_LOOKUP
	ACCESS4_MODIFY = 0x00000004 // ACCESS4_MODIFY
	ACCESS4_EXTEND = 0x00000008 // ACCESS4_EXTEND
	ACCESS4_DELETE = 0x00000010 // ACCESS4_DELETE
	ACCESS4_EXECUTE = 0x00000020 // ACCESS4_EXECUTE
	OPEN4_SHARE_ACCESS_READ = 0x00000001 // OPEN4_SHARE_ACCESS_READ
	OPEN4_SHARE_ACCESS_WRITE = 0x00000002 // OPEN4_SHARE_ACCESS_WRITE
	OPEN4_SHARE_ACCESS_BOTH = 0x00000003 // OPEN4_SHARE_ACCESS_BOTH
	OPEN4_SHARE_DENY_NONE = 0x00000000 // OPEN4_SHARE_DENY_NONE
	OPEN4_SHARE_DENY_READ = 0x00000001 // OPEN4_SHARE_DENY_READ
	OPEN4_SHARE_DENY_WRITE = 0x00000002 // OPEN4_SHARE_DENY_WRITE
	OPEN4_SHARE_DENY_BOTH = 0x00000003 // OPEN4_SHARE_DENY_BOTH
	OPEN4_RESULT_CONFIRM = 0x00000002 // OPEN4_RESULT_CONFIRM
	OPEN4_RESULT_LOCKTYPE_POSIX = 0x00000004 // OPEN4_RESULT_LOCKTYPE_POSIX

)

// Most of the typedefs are replaced by basic go types

type NfsFh4 [NFS4_FHSIZE]byte
type Verifier4 [NFS4_VERIFIER_SIZE]byte // verifier4
type Fattr4Fsid Fsid4 // fattr4_fsid
type Fattr4Acl []Nfsace4 // fattr4_acl
type Fattr4FsLocations FsLocations4 // fattr4_fs_locations
type Fattr4Rawdev Specdata4 // fattr4_rawdev
type Fattr4TimeAccess Nfstime4 // fattr4_time_access
type Fattr4TimeAccessSet Settime4 // fattr4_time_access_set
type Fattr4TimeBackup Nfstime4 // fattr4_time_backup
type Fattr4TimeCreate Nfstime4 // fattr4_time_create
type Fattr4TimeDelta Nfstime4 // fattr4_time_delta
type Fattr4TimeMetadata Nfstime4 // fattr4_time_metadata
type Fattr4TimeModify Nfstime4 // fattr4_time_modify
type Fattr4TimeModifySet Settime4 // fattr4_time_modify_set
type SECINFO4resok []Secinfo4 // SECINFO4resok


type Nfstime4 struct { // nfstime4
	Seconds int64
	Nseconds uint32
}


type Fsid4 struct { // fsid4
	Major uint64
	Minor uint64
}


type FsLocation4 struct { // fs_location4
	Server []string
	Rootpath []string
}


type FsLocations4 struct { // fs_locations4
	FsRoot []string
	Locations []FsLocation4
}


type Nfsace4 struct { // nfsace4
	Type uint32
	Flag uint32
	AccessMask uint32
	Who string
}


type Specdata4 struct { // specdata4
	Specdata1 uint32
	Specdata2 uint32
}


type Fattr4 struct { // fattr4
	Attrmask []uint32
	AttrVals []byte
}


type ChangeInfo4 struct { // change_info4
	Atomic bool
	Before uint64
	After uint64
}


type Clientaddr4 struct { // clientaddr4
	RNetid string
	RAddr string
}


type CbClient4 struct { // cb_client4
	CbProgram uint32
	CbLocation Clientaddr4
}


type Stateid4 struct { // stateid4
	Seqid uint32
	Other [NFS4_OTHER_SIZE]byte
}


type NfsClientID4 struct { // nfs_client_id4
	Verifier Verifier4
	ID string //[]byte
}


type OpenOwner4 struct { // open_owner4
	Clientid uint64
	Owner string //[]byte
}


type LockOwner4 struct { // lock_owner4
	Clientid uint64
	Owner string //[]byte
}


type ACCESS4args struct { // ACCESS4args
	Access uint32
}


type ACCESS4resok struct { // ACCESS4resok
	Supported uint32
	Access uint32
}


type CLOSE4args struct { // CLOSE4args
	Seqid uint32
	OpenStateid Stateid4
}


type COMMIT4args struct { // COMMIT4args
	Offset uint32
	Count uint32
}


type COMMIT4resok struct { // COMMIT4resok
	Writeverf Verifier4
}


type CREATE4args struct { // CREATE4args
	Objtype Createtype4
	Objname string
	Createattrs Fattr4
}


type CREATE4resok struct { // CREATE4resok
	Cinfo ChangeInfo4
	Attrset []uint32
}


type DELEGPURGE4args struct { // DELEGPURGE4args
	Clientid uint64
}


type DELEGPURGE4res struct { // DELEGPURGE4res
	Status int32
}


type DELEGRETURN4args struct { // DELEGRETURN4args
	DelegStateid Stateid4
}


type DELEGRETURN4res struct { // DELEGRETURN4res
	Status int32
}


type GETATTR4args struct { // GETATTR4args
	AttrRequest []uint32
}


type GETATTR4resok struct { // GETATTR4resok
	ObjAttributes Fattr4
}


type GETFH4resok struct { // GETFH4resok
	Object NfsFh4
}


type LINK4args struct { // LINK4args
	Newname string
}


type LINK4resok struct { // LINK4resok
	Cinfo ChangeInfo4
}


type OpenToLockOwner4 struct { // open_to_lock_owner4
	OpenSeqid uint32
	OpenStateid Stateid4
	LockSeqid uint32
	LockOwner LockOwner4
}


type ExistLockOwner4 struct { // exist_lock_owner4
	LockStateid Stateid4
	LockSeqid uint32
}


type LOCK4args struct { // LOCK4args
	Locktype int32
	Reclaim bool
	Offset uint32
	Length uint64
	Locker Locker4
}


type LOCK4denied struct { // LOCK4denied
	Offset uint32
	Length uint64
	Locktype int32
	Owner LockOwner4
}


type LOCK4resok struct { // LOCK4resok
	LockStateid Stateid4
}


type LOCKT4args struct { // LOCKT4args
	Locktype int32
	Offset uint32
	Length uint64
	Owner LockOwner4
}


type LOCKU4args struct { // LOCKU4args
	Locktype int32
	Seqid uint32
	LockStateid Stateid4
	Offset uint32
	Length uint64
}


type LOOKUP4args struct { // LOOKUP4args
	Objname string
}


type LOOKUP4res struct { // LOOKUP4res
	Status int32
}


type LOOKUPP4res struct { // LOOKUPP4res
	Status int32
}


type NVERIFY4args struct { // NVERIFY4args
	ObjAttributes Fattr4
}


type NVERIFY4res struct { // NVERIFY4res
	Status int32
}


type NfsModifiedLimit4 struct { // nfs_modified_limit4
	NumBlocks uint32
	BytesPerBlock uint32
}


type OpenClaimDelegateCur4 struct { // open_claim_delegate_cur4
	DelegateStateid Stateid4
	File string
}


type OPEN4args struct { // OPEN4args
	Seqid uint32
	ShareAccess uint32
	ShareDeny uint32
	Owner OpenOwner4
	Openhow Openflag4
	Claim OpenClaim4
}


type OpenReadDelegation4 struct { // open_read_delegation4
	Stateid Stateid4
	Recall bool
	Permissions Nfsace4
}


type OpenWriteDelegation4 struct { // open_write_delegation4
	Stateid Stateid4
	Recall bool
	SpaceLimit NfsSpaceLimit4
	Permissions Nfsace4
}


type OPEN4resok struct { // OPEN4resok
	Stateid Stateid4
	Cinfo ChangeInfo4
	Rflags uint32
	Attrset []uint32
	Delegation OpenDelegation4
}


type OPENATTR4args struct { // OPENATTR4args
	Createdir bool
}


type OPENATTR4res struct { // OPENATTR4res
	Status int32
}


type OpenConfirm4args struct { // OPEN_CONFIRM4args
	OpenStateid Stateid4
	Seqid uint32
}


type OpenConfirm4resok struct { // OPEN_CONFIRM4resok
	OpenStateid Stateid4
}


type OpenDowngrade4args struct { // OPEN_DOWNGRADE4args
	OpenStateid Stateid4
	Seqid uint32
	ShareAccess uint32
	ShareDeny uint32
}


type OpenDowngrade4resok struct { // OPEN_DOWNGRADE4resok
	OpenStateid Stateid4
}


type PUTFH4args struct { // PUTFH4args
	Object NfsFh4
}


type PUTFH4res struct { // PUTFH4res
	Status int32
}


type PUTPUBFH4res struct { // PUTPUBFH4res
	Status int32
}


type PUTROOTFH4res struct { // PUTROOTFH4res
	Status int32
}


type READ4args struct { // READ4args
	Stateid Stateid4
	Offset uint32
	Count uint32
}


type READ4resok struct { // READ4resok
	Eof bool
	Data []byte
}


type READDIR4args struct { // READDIR4args
	Cookie uint64
	Cookieverf Verifier4
	Dircount uint32
	Maxcount uint32
	AttrRequest []uint32
}


type Entry4 struct { // entry4
	Cookie uint64
	Name string
	Attrs Fattr4
	OptNextentry []Entry4
}


type Dirlist4 struct { // dirlist4
	OptEntries []Entry4
	Eof bool
}


type READDIR4resok struct { // READDIR4resok
	Cookieverf Verifier4
	Reply Dirlist4
}


type READLINK4resok struct { // READLINK4resok
	Link string
}


type REMOVE4args struct { // REMOVE4args
	Target string
}


type REMOVE4resok struct { // REMOVE4resok
	Cinfo ChangeInfo4
}


type RENAME4args struct { // RENAME4args
	Oldname string
	Newname string
}


type RENAME4resok struct { // RENAME4resok
	SourceCinfo ChangeInfo4
	TargetCinfo ChangeInfo4
}


type RENEW4args struct { // RENEW4args
	Clientid uint64
}


type RENEW4res struct { // RENEW4res
	Status int32
}


type RESTOREFH4res struct { // RESTOREFH4res
	Status int32
}


type SAVEFH4res struct { // SAVEFH4res
	Status int32
}


type SECINFO4args struct { // SECINFO4args
	Name string
}


type RpcsecGssInfo struct { // rpcsec_gss_info
	Oid []byte
	Qop uint32
	Service int32
}


type SETATTR4args struct { // SETATTR4args
	Stateid Stateid4
	ObjAttributes Fattr4
}


type SETATTR4res struct { // SETATTR4res
	Status int32
	Attrsset []uint32
}


type SETCLIENTID4args struct { // SETCLIENTID4args
	Client NfsClientID4
	Callback CbClient4
	CallbackIdent uint32
}


type SETCLIENTID4resok struct { // SETCLIENTID4resok
	Clientid uint64
	SetclientidConfirm Verifier4
}


type SetclientidConfirm4args struct { // SETCLIENTID_CONFIRM4args
	Clientid uint64
	SetclientidConfirm Verifier4
}


type SetclientidConfirm4res struct { // SETCLIENTID_CONFIRM4res
	Status int32
}


type VERIFY4args struct { // VERIFY4args
	ObjAttributes Fattr4
}


type VERIFY4res struct { // VERIFY4res
	Status int32
}


type WRITE4args struct { // WRITE4args
	Stateid Stateid4
	Offset uint32
	Stable int32
	Data []byte
}


type WRITE4resok struct { // WRITE4resok
	Count uint32
	Committed int32
	Writeverf Verifier4
}


type ReleaseLockowner4args struct { // RELEASE_LOCKOWNER4args
	LockOwner LockOwner4
}


type ReleaseLockowner4res struct { // RELEASE_LOCKOWNER4res
	Status int32
}


type ILLEGAL4res struct { // ILLEGAL4res
	Status int32
}


type COMPOUND4args struct { // COMPOUND4args
	Tag string
	Minorversion uint32
	Argarray []NfsArgop4
}


type COMPOUND4res struct { // COMPOUND4res
	Status int32
	Tag string
	Resarray []NfsResop4
}


type CbGetattr4args struct { // CB_GETATTR4args
	Fh NfsFh4
	AttrRequest []uint32
}


type CbGetattr4resok struct { // CB_GETATTR4resok
	ObjAttributes Fattr4
}


type CbRecall4args struct { // CB_RECALL4args
	Stateid Stateid4
	Truncate bool
	Fh NfsFh4
}


type CbRecall4res struct { // CB_RECALL4res
	Status int32
}


type CbIllegal4res struct { // CB_ILLEGAL4res
	Status int32
}


type CbCompound4args struct { // CB_COMPOUND4args
	Tag string
	Minorversion uint32
	CallbackIdent uint32
	Argarray []NfsCbArgop4
}


type CbCompound4res struct { // CB_COMPOUND4res
	Status int32
	Tag string
	Resarray []NfsCbResop4
}





type Settime4 struct {
	SetIt int32 `xdr:"union"`
	Time  Nfstime4  `xdr:"unioncase=1"` // Time4
}


type ACCESS4res struct {
	Status int32 `xdr:"union"`
	Resok4  ACCESS4resok  `xdr:"unioncase=0"` // Ok
}


type CLOSE4res struct {
	Status int32 `xdr:"union"`
	OpenStateid  Stateid4  `xdr:"unioncase=0"` // Ok
}


type COMMIT4res struct {
	Status int32 `xdr:"union"`
	Resok4  COMMIT4resok  `xdr:"unioncase=0"` // Ok
}


type Createtype4 struct {
	Type int32 `xdr:"union"`
	Linkdata  string  `xdr:"unioncase=5"` // Nf4lnk
	Blkdata  Specdata4  `xdr:"unioncase=3"` // Nf4blk
	Chrdata  Specdata4  `xdr:"unioncase=4"` // Nf4chr
}


type CREATE4res struct {
	Status int32 `xdr:"union"`
	Resok4  CREATE4resok  `xdr:"unioncase=0"` // Ok
}


type GETATTR4res struct {
	Status int32 `xdr:"union"`
	Resok4  GETATTR4resok  `xdr:"unioncase=0"` // Ok
}


type GETFH4res struct {
	Status int32 `xdr:"union"`
	Resok4  GETFH4resok  `xdr:"unioncase=0"` // Ok
}


type LINK4res struct {
	Status int32 `xdr:"union"`
	Resok4  LINK4resok  `xdr:"unioncase=0"` // Ok
}


type Locker4 struct {
	NewLockOwner bool `xdr:"union"`
	OpenOwner  OpenToLockOwner4  `xdr:"unioncase=1"` // True
	LockOwner  ExistLockOwner4  `xdr:"unioncase=0"` // False
}


type LOCK4res struct {
	Status int32 `xdr:"union"`
	Resok4  LOCK4resok  `xdr:"unioncase=0"` // Ok
	Denied  LOCK4denied  `xdr:"unioncase=10010"` // Denied
}


type LOCKT4res struct {
	Status int32 `xdr:"union"`
	Denied  LOCK4denied  `xdr:"unioncase=10010"` // Denied
	//    `xdr:"unioncase=0"` // Ok
}


type LOCKU4res struct {
	Status int32 `xdr:"union"`
	LockStateid  Stateid4  `xdr:"unioncase=0"` // Ok
}


type Createhow4 struct {
	Mode int32 `xdr:"union"`
	// TODO: codegen failed
	CreateattrsUnchecked  Fattr4  `xdr:"unioncase=0"` // Unchecked4
	CreateattrsGuarded  Fattr4  `xdr:"unioncase=1"` // Guarded4
	Createverf  Verifier4  `xdr:"unioncase=2"` // Exclusive4
}


type Openflag4 struct {
	Opentype int32 `xdr:"union"`
	How  Createhow4  `xdr:"unioncase=1"` // Create
}


type NfsSpaceLimit4 struct {
	Limitby int32 `xdr:"union"`
	Filesize  uint64  `xdr:"unioncase=1"` // Size
	ModBlocks  NfsModifiedLimit4  `xdr:"unioncase=2"` // Blocks
}


type OpenClaim4 struct {
	Claim int32 `xdr:"union"`
	File  string  `xdr:"unioncase=0"` // Null
	DelegateType  int32  `xdr:"unioncase=1"` // Previous
	DelegateCurInfo  OpenClaimDelegateCur4  `xdr:"unioncase=2"` // Cur
	FileDelegatePrev  string  `xdr:"unioncase=3"` // Prev
}


type OpenDelegation4 struct {
	DelegationType int32 `xdr:"union"`
	//    `xdr:"unioncase=0"` // None
	Read  OpenReadDelegation4  `xdr:"unioncase=1"` // Read
	Write  OpenWriteDelegation4  `xdr:"unioncase=2"` // Write
}


type OPEN4res struct {
	Status int32 `xdr:"union"`
	Resok4  OPEN4resok  `xdr:"unioncase=0"` // Ok
}


type OpenConfirm4res struct {
	Status int32 `xdr:"union"`
	Resok4  OpenConfirm4resok  `xdr:"unioncase=0"` // Ok
}


type OpenDowngrade4res struct {
	Status int32 `xdr:"union"`
	Resok4  OpenDowngrade4resok  `xdr:"unioncase=0"` // Ok
}


type READ4res struct {
	Status int32 `xdr:"union"`
	Resok4  READ4resok  `xdr:"unioncase=0"` // Ok
}


type READDIR4res struct {
	Status int32 `xdr:"union"`
	Resok4  READDIR4resok  `xdr:"unioncase=0"` // Ok
}


type READLINK4res struct {
	Status int32 `xdr:"union"`
	Resok4  READLINK4resok  `xdr:"unioncase=0"` // Ok
}


type REMOVE4res struct {
	Status int32 `xdr:"union"`
	Resok4  REMOVE4resok  `xdr:"unioncase=0"` // Ok
}


type RENAME4res struct {
	Status int32 `xdr:"union"`
	Resok4  RENAME4resok  `xdr:"unioncase=0"` // Ok
}


type Secinfo4 struct {
	Flavor uint32 `xdr:"union"`
	FlavorInfo  RpcsecGssInfo  `xdr:"unioncase=6"` // Gss
}


type SECINFO4res struct {
	Status int32 `xdr:"union"`
	Resok4  SECINFO4resok  `xdr:"unioncase=0"` // Ok
}


type SETCLIENTID4res struct {
	Status int32 `xdr:"union"`
	Resok4  SETCLIENTID4resok  `xdr:"unioncase=0"` // Ok
	ClientUsing  Clientaddr4  `xdr:"unioncase=10017"` // Inuse
}


type WRITE4res struct {
	Status int32 `xdr:"union"`
	Resok4  WRITE4resok  `xdr:"unioncase=0"` // Ok
}


type NfsArgop4 struct {
	Argop int32 `xdr:"union"`
	Opaccess  ACCESS4args  `xdr:"unioncase=3"` // Access
	Opclose  CLOSE4args  `xdr:"unioncase=4"` // Close
	Opcommit  COMMIT4args  `xdr:"unioncase=5"` // Commit
	Opcreate  CREATE4args  `xdr:"unioncase=6"` // Create
	Opdelegpurge  DELEGPURGE4args  `xdr:"unioncase=7"` // Delegpurge
	Opdelegreturn  DELEGRETURN4args  `xdr:"unioncase=8"` // Delegreturn
	Opgetattr  GETATTR4args  `xdr:"unioncase=9"` // Getattr
	//    `xdr:"unioncase=10"` // Getfh
	Oplink  LINK4args  `xdr:"unioncase=11"` // Link
	Oplock  LOCK4args  `xdr:"unioncase=12"` // Lock
	Oplockt  LOCKT4args  `xdr:"unioncase=13"` // Lockt
	Oplocku  LOCKU4args  `xdr:"unioncase=14"` // Locku
	Oplookup  LOOKUP4args  `xdr:"unioncase=15"` // Lookup
	//    `xdr:"unioncase=16"` // Lookupp
	Opnverify  NVERIFY4args  `xdr:"unioncase=17"` // Nverify
	Opopen  OPEN4args  `xdr:"unioncase=18"` // Open
	Opopenattr  OPENATTR4args  `xdr:"unioncase=19"` // Openattr
	OpopenConfirm  OpenConfirm4args  `xdr:"unioncase=20"` // Confirm
	OpopenDowngrade  OpenDowngrade4args  `xdr:"unioncase=21"` // Downgrade
	Opputfh  PUTFH4args  `xdr:"unioncase=22"` // Putfh
	//    `xdr:"unioncase=23"` // Putpubfh
	//    `xdr:"unioncase=24"` // Putrootfh
	Opread  READ4args  `xdr:"unioncase=25"` // Read
	Opreaddir  READDIR4args  `xdr:"unioncase=26"` // Readdir
	//    `xdr:"unioncase=27"` // Readlink
	Opremove  REMOVE4args  `xdr:"unioncase=28"` // Remove
	Oprename  RENAME4args  `xdr:"unioncase=29"` // Rename
	Oprenew  RENEW4args  `xdr:"unioncase=30"` // Renew
	//    `xdr:"unioncase=31"` // Restorefh
	//    `xdr:"unioncase=32"` // Savefh
	Opsecinfo  SECINFO4args  `xdr:"unioncase=33"` // Secinfo
	Opsetattr  SETATTR4args  `xdr:"unioncase=34"` // Setattr
	Opsetclientid  SETCLIENTID4args  `xdr:"unioncase=35"` // Setclientid
	OpsetclientidConfirm  SetclientidConfirm4args  `xdr:"unioncase=36"` // Confirm
	Opverify  VERIFY4args  `xdr:"unioncase=37"` // Verify
	Opwrite  WRITE4args  `xdr:"unioncase=38"` // Write
	OpreleaseLockowner  ReleaseLockowner4args  `xdr:"unioncase=39"` // Lockowner
	//    `xdr:"unioncase=10044"` // Illegal
}


type NfsResop4 struct {
	Resop int32 `xdr:"union"`
	Opaccess  ACCESS4res  `xdr:"unioncase=3"` // Access
	Opclose  CLOSE4res  `xdr:"unioncase=4"` // Close
	Opcommit  COMMIT4res  `xdr:"unioncase=5"` // Commit
	Opcreate  CREATE4res  `xdr:"unioncase=6"` // Create
	Opdelegpurge  DELEGPURGE4res  `xdr:"unioncase=7"` // Delegpurge
	Opdelegreturn  DELEGRETURN4res  `xdr:"unioncase=8"` // Delegreturn
	Opgetattr  GETATTR4res  `xdr:"unioncase=9"` // Getattr
	Opgetfh  GETFH4res  `xdr:"unioncase=10"` // Getfh
	Oplink  LINK4res  `xdr:"unioncase=11"` // Link
	Oplock  LOCK4res  `xdr:"unioncase=12"` // Lock
	Oplockt  LOCKT4res  `xdr:"unioncase=13"` // Lockt
	Oplocku  LOCKU4res  `xdr:"unioncase=14"` // Locku
	Oplookup  LOOKUP4res  `xdr:"unioncase=15"` // Lookup
	Oplookupp  LOOKUPP4res  `xdr:"unioncase=16"` // Lookupp
	Opnverify  NVERIFY4res  `xdr:"unioncase=17"` // Nverify
	Opopen  OPEN4res  `xdr:"unioncase=18"` // Open
	Opopenattr  OPENATTR4res  `xdr:"unioncase=19"` // Openattr
	OpopenConfirm  OpenConfirm4res  `xdr:"unioncase=20"` // Confirm
	OpopenDowngrade  OpenDowngrade4res  `xdr:"unioncase=21"` // Downgrade
	Opputfh  PUTFH4res  `xdr:"unioncase=22"` // Putfh
	Opputpubfh  PUTPUBFH4res  `xdr:"unioncase=23"` // Putpubfh
	Opputrootfh  PUTROOTFH4res  `xdr:"unioncase=24"` // Putrootfh
	Opread  READ4res  `xdr:"unioncase=25"` // Read
	Opreaddir  READDIR4res  `xdr:"unioncase=26"` // Readdir
	Opreadlink  READLINK4res  `xdr:"unioncase=27"` // Readlink
	Opremove  REMOVE4res  `xdr:"unioncase=28"` // Remove
	Oprename  RENAME4res  `xdr:"unioncase=29"` // Rename
	Oprenew  RENEW4res  `xdr:"unioncase=30"` // Renew
	Oprestorefh  RESTOREFH4res  `xdr:"unioncase=31"` // Restorefh
	Opsavefh  SAVEFH4res  `xdr:"unioncase=32"` // Savefh
	Opsecinfo  SECINFO4res  `xdr:"unioncase=33"` // Secinfo
	Opsetattr  SETATTR4res  `xdr:"unioncase=34"` // Setattr
	Opsetclientid  SETCLIENTID4res  `xdr:"unioncase=35"` // Setclientid
	OpsetclientidConfirm  SetclientidConfirm4res  `xdr:"unioncase=36"` // Confirm
	Opverify  VERIFY4res  `xdr:"unioncase=37"` // Verify
	Opwrite  WRITE4res  `xdr:"unioncase=38"` // Write
	OpreleaseLockowner  ReleaseLockowner4res  `xdr:"unioncase=39"` // Lockowner
	Opillegal  ILLEGAL4res  `xdr:"unioncase=10044"` // Illegal
}


type CbGetattr4res struct {
	Status int32 `xdr:"union"`
	Resok4  CbGetattr4resok  `xdr:"unioncase=0"` // Ok
}


type NfsCbArgop4 struct {
	Argop uint32 `xdr:"union"`
	Opcbgetattr  CbGetattr4args  `xdr:"unioncase=3"` // Getattr
	Opcbrecall  CbRecall4args  `xdr:"unioncase=4"` // Recall
	//    `xdr:"unioncase=10044"` // Illegal
}


type NfsCbResop4 struct {
	Resop uint32 `xdr:"union"`
	Opcbgetattr  CbGetattr4res  `xdr:"unioncase=3"` // Getattr
	Opcbrecall  CbRecall4res  `xdr:"unioncase=4"` // Recall
	Opcbillegal  CbIllegal4res  `xdr:"unioncase=10044"` // Illegal
}





func Access (access uint32) (NfsArgop4) {
	return NfsArgop4{Argop:3, Opaccess:ACCESS4args{ Access:access } }
}


func Close (seqid uint32, openstateid Stateid4) (NfsArgop4) {
	return NfsArgop4{Argop:4, Opclose:CLOSE4args{ Seqid:seqid, OpenStateid:openstateid } }
}


func Commit (offset uint32, count uint32) (NfsArgop4) {
	return NfsArgop4{Argop:5, Opcommit:COMMIT4args{ Offset:offset, Count:count } }
}


func Create (objtype Createtype4, objname string, createattrs Fattr4) (NfsArgop4) {
	return NfsArgop4{Argop:6, Opcreate:CREATE4args{ Objtype:objtype, Objname:objname, Createattrs:createattrs } }
}


func Delegpurge (clientid uint64) (NfsArgop4) {
	return NfsArgop4{Argop:7, Opdelegpurge:DELEGPURGE4args{ Clientid:clientid } }
}


func Delegreturn (delegstateid Stateid4) (NfsArgop4) {
	return NfsArgop4{Argop:8, Opdelegreturn:DELEGRETURN4args{ DelegStateid:delegstateid } }
}


func Getattr (attrrequest []uint32) (NfsArgop4) {
	return NfsArgop4{Argop:9, Opgetattr:GETATTR4args{ AttrRequest:attrrequest } }
}


func Getfh () (NfsArgop4) {
	return NfsArgop4{Argop:10 }
}


func Link (newname string) (NfsArgop4) {
	return NfsArgop4{Argop:11, Oplink:LINK4args{ Newname:newname } }
}


func Lock (locktype int32, reclaim bool, offset uint32, length uint64, locker Locker4) (NfsArgop4) {
	return NfsArgop4{Argop:12, Oplock:LOCK4args{ Locktype:locktype, Reclaim:reclaim, Offset:offset, Length:length, Locker:locker } }
}


func Lockt (locktype int32, offset uint32, length uint64, owner LockOwner4) (NfsArgop4) {
	return NfsArgop4{Argop:13, Oplockt:LOCKT4args{ Locktype:locktype, Offset:offset, Length:length, Owner:owner } }
}


func Locku (locktype int32, seqid uint32, lockstateid Stateid4, offset uint32, length uint64) (NfsArgop4) {
	return NfsArgop4{Argop:14, Oplocku:LOCKU4args{ Locktype:locktype, Seqid:seqid, LockStateid:lockstateid, Offset:offset, Length:length } }
}


func Lookup (objname string) (NfsArgop4) {
	return NfsArgop4{Argop:15, Oplookup:LOOKUP4args{ Objname:objname } }
}


func Lookupp () (NfsArgop4) {
	return NfsArgop4{Argop:16 }
}


func Nverify (objattributes Fattr4) (NfsArgop4) {
	return NfsArgop4{Argop:17, Opnverify:NVERIFY4args{ ObjAttributes:objattributes } }
}


func Open (seqid uint32, shareaccess uint32, sharedeny uint32, owner OpenOwner4, openhow Openflag4, claim OpenClaim4) (NfsArgop4) {
	return NfsArgop4{Argop:18, Opopen:OPEN4args{ Seqid:seqid, ShareAccess:shareaccess, ShareDeny:sharedeny, Owner:owner, Openhow:openhow, Claim:claim } }
}


func Openattr (createdir bool) (NfsArgop4) {
	return NfsArgop4{Argop:19, Opopenattr:OPENATTR4args{ Createdir:createdir } }
}

// TODO: codegen failed
func OpenConfirm (openstateid Stateid4, seqid uint32) (NfsArgop4) {
	return NfsArgop4{Argop:20, OpopenConfirm:OpenConfirm4args{ OpenStateid:openstateid, Seqid:seqid } }
}


func OpenDowngrade (openstateid Stateid4, seqid uint32, shareaccess uint32, sharedeny uint32) (NfsArgop4) {
	return NfsArgop4{Argop:21, OpopenDowngrade:OpenDowngrade4args{ OpenStateid:openstateid, Seqid:seqid, ShareAccess:shareaccess, ShareDeny:sharedeny } }
}


func Putfh (object NfsFh4) (NfsArgop4) {
	return NfsArgop4{Argop:22, Opputfh:PUTFH4args{ Object:object } }
}


func Putpubfh () (NfsArgop4) {
	return NfsArgop4{Argop:23 }
}


func Putrootfh () (NfsArgop4) {
	return NfsArgop4{Argop:24 }
}


func Read (stateid Stateid4, offset uint32, count uint32) (NfsArgop4) {
	return NfsArgop4{Argop:25, Opread:READ4args{ Stateid:stateid, Offset:offset, Count:count } }
}


func Readdir (cookie uint64, cookieverf Verifier4, dircount uint32, maxcount uint32, attrrequest []uint32) (NfsArgop4) {
	return NfsArgop4{Argop:26, Opreaddir:READDIR4args{ Cookie:cookie, Cookieverf:cookieverf, Dircount:dircount, Maxcount:maxcount, AttrRequest:attrrequest } }
}


func Readlink () (NfsArgop4) {
	return NfsArgop4{Argop:27 }
}


func Remove (target string) (NfsArgop4) {
	return NfsArgop4{Argop:28, Opremove:REMOVE4args{ Target:target } }
}


func Rename (oldname string, newname string) (NfsArgop4) {
	return NfsArgop4{Argop:29, Oprename:RENAME4args{ Oldname:oldname, Newname:newname } }
}


func Renew (clientid uint64) (NfsArgop4) {
	return NfsArgop4{Argop:30, Oprenew:RENEW4args{ Clientid:clientid } }
}


func Restorefh () (NfsArgop4) {
	return NfsArgop4{Argop:31 }
}


func Savefh () (NfsArgop4) {
	return NfsArgop4{Argop:32 }
}


func Secinfo (name string) (NfsArgop4) {
	return NfsArgop4{Argop:33, Opsecinfo:SECINFO4args{ Name:name } }
}


func Setattr (stateid Stateid4, objattributes Fattr4) (NfsArgop4) {
	return NfsArgop4{Argop:34, Opsetattr:SETATTR4args{ Stateid:stateid, ObjAttributes:objattributes } }
}


func Setclientid (client NfsClientID4, callback CbClient4, callbackident uint32) (NfsArgop4) {
	return NfsArgop4{Argop:35, Opsetclientid:SETCLIENTID4args{ Client:client, Callback:callback, CallbackIdent:callbackident } }
}


func SetclientidConfirm (clientid uint64, setclientidconfirm Verifier4) (NfsArgop4) {
	return NfsArgop4{Argop:36, OpsetclientidConfirm:SetclientidConfirm4args{ Clientid:clientid, SetclientidConfirm:setclientidconfirm } }
}


func Verify (objattributes Fattr4) (NfsArgop4) {
	return NfsArgop4{Argop:37, Opverify:VERIFY4args{ ObjAttributes:objattributes } }
}


func Write (stateid Stateid4, offset uint32, stable int32, data []byte) (NfsArgop4) {
	return NfsArgop4{Argop:38, Opwrite:WRITE4args{ Stateid:stateid, Offset:offset, Stable:stable, Data:data } }
}


func ReleaseLockowner (lockowner LockOwner4) (NfsArgop4) {
	return NfsArgop4{Argop:39, OpreleaseLockowner:ReleaseLockowner4args{ LockOwner:lockowner } }
}


func Illegal () (NfsArgop4) {
	return NfsArgop4{Argop:10044 }
}


