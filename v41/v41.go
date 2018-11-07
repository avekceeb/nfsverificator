package v41

/*
TODO: ???
in CallbackSecParms4:
	AUTH_NONE = 1
	AUTH_SYS = 2
	RPCSEC_GSS = 3
 */
// TODO: to GSS or smth
// Added manually (Dmitry A.)
type AuthsysParms struct {
	Stamp       uint32
	Machinename string
	Uid         uint32
	Gid         uint32
	GidLen      uint32
	// this was producing extra 4-byte field
	//Gids        uint32
}




const (
	/* Enums: */
	NF4REG = 1
	NF4DIR = 2
	NF4BLK = 3
	NF4CHR = 4
	NF4LNK = 5
	NF4SOCK = 6
	NF4FIFO = 7
	NF4ATTRDIR = 8
	NF4NAMEDATTR = 9
	NFS4_OK = 0
	NFS4ERR_PERM = 1
	NFS4ERR_NOENT = 2
	NFS4ERR_IO = 5
	NFS4ERR_NXIO = 6
	NFS4ERR_ACCESS = 13
	NFS4ERR_EXIST = 17
	NFS4ERR_XDEV = 18
	NFS4ERR_NOTDIR = 20
	NFS4ERR_ISDIR = 21
	NFS4ERR_INVAL = 22
	NFS4ERR_FBIG = 27
	NFS4ERR_NOSPC = 28
	NFS4ERR_ROFS = 30
	NFS4ERR_MLINK = 31
	NFS4ERR_NAMETOOLONG = 63
	NFS4ERR_NOTEMPTY = 66
	NFS4ERR_DQUOT = 69
	NFS4ERR_STALE = 70
	NFS4ERR_BADHANDLE = 10001
	NFS4ERR_BAD_COOKIE = 10003
	NFS4ERR_NOTSUPP = 10004
	NFS4ERR_TOOSMALL = 10005
	NFS4ERR_SERVERFAULT = 10006
	NFS4ERR_BADTYPE = 10007
	NFS4ERR_DELAY = 10008
	NFS4ERR_SAME = 10009
	NFS4ERR_DENIED = 10010
	NFS4ERR_EXPIRED = 10011
	NFS4ERR_LOCKED = 10012
	NFS4ERR_GRACE = 10013
	NFS4ERR_FHEXPIRED = 10014
	NFS4ERR_SHARE_DENIED = 10015
	NFS4ERR_WRONGSEC = 10016
	NFS4ERR_CLID_INUSE = 10017
	NFS4ERR_RESOURCE = 10018
	NFS4ERR_MOVED = 10019
	NFS4ERR_NOFILEHANDLE = 10020
	NFS4ERR_MINOR_VERS_MISMATCH = 10021
	NFS4ERR_STALE_CLIENTID = 10022
	NFS4ERR_STALE_STATEID = 10023
	NFS4ERR_OLD_STATEID = 10024
	NFS4ERR_BAD_STATEID = 10025
	NFS4ERR_BAD_SEQID = 10026
	NFS4ERR_NOT_SAME = 10027
	NFS4ERR_LOCK_RANGE = 10028
	NFS4ERR_SYMLINK = 10029
	NFS4ERR_RESTOREFH = 10030
	NFS4ERR_LEASE_MOVED = 10031
	NFS4ERR_ATTRNOTSUPP = 10032
	NFS4ERR_NO_GRACE = 10033
	NFS4ERR_RECLAIM_BAD = 10034
	NFS4ERR_RECLAIM_CONFLICT = 10035
	NFS4ERR_BADXDR = 10036
	NFS4ERR_LOCKS_HELD = 10037
	NFS4ERR_OPENMODE = 10038
	NFS4ERR_BADOWNER = 10039
	NFS4ERR_BADCHAR = 10040
	NFS4ERR_BADNAME = 10041
	NFS4ERR_BAD_RANGE = 10042
	NFS4ERR_LOCK_NOTSUPP = 10043
	NFS4ERR_OP_ILLEGAL = 10044
	NFS4ERR_DEADLOCK = 10045
	NFS4ERR_FILE_OPEN = 10046
	NFS4ERR_ADMIN_REVOKED = 10047
	NFS4ERR_CB_PATH_DOWN = 10048
	NFS4ERR_BADIOMODE = 10049
	NFS4ERR_BADLAYOUT = 10050
	NFS4ERR_BAD_SESSION_DIGEST = 10051
	NFS4ERR_BADSESSION = 10052
	NFS4ERR_BADSLOT = 10053
	NFS4ERR_COMPLETE_ALREADY = 10054
	NFS4ERR_CONN_NOT_BOUND_TO_SESSION = 10055
	NFS4ERR_DELEG_ALREADY_WANTED = 10056
	NFS4ERR_BACK_CHAN_BUSY = 10057
	NFS4ERR_LAYOUTTRYLATER = 10058
	NFS4ERR_LAYOUTUNAVAILABLE = 10059
	NFS4ERR_NOMATCHING_LAYOUT = 10060
	NFS4ERR_RECALLCONFLICT = 10061
	NFS4ERR_UNKNOWN_LAYOUTTYPE = 10062
	NFS4ERR_SEQ_MISORDERED = 10063
	NFS4ERR_SEQUENCE_POS = 10064
	NFS4ERR_REQ_TOO_BIG = 10065
	NFS4ERR_REP_TOO_BIG = 10066
	NFS4ERR_REP_TOO_BIG_TO_CACHE = 10067
	NFS4ERR_RETRY_UNCACHED_REP = 10068
	NFS4ERR_UNSAFE_COMPOUND = 10069
	NFS4ERR_TOO_MANY_OPS = 10070
	NFS4ERR_OP_NOT_IN_SESSION = 10071
	NFS4ERR_HASH_ALG_UNSUPP = 10072
	NFS4ERR_CLIENTID_BUSY = 10074
	NFS4ERR_PNFS_IO_HOLE = 10075
	NFS4ERR_SEQ_FALSE_RETRY = 10076
	NFS4ERR_BAD_HIGH_SLOT = 10077
	NFS4ERR_DEADSESSION = 10078
	NFS4ERR_ENCR_ALG_UNSUPP = 10079
	NFS4ERR_PNFS_NO_LAYOUT = 10080
	NFS4ERR_NOT_ONLY_OP = 10081
	NFS4ERR_WRONG_CRED = 10082
	NFS4ERR_WRONG_TYPE = 10083
	NFS4ERR_DIRDELEG_UNAVAIL = 10084
	NFS4ERR_REJECT_DELEG = 10085
	NFS4ERR_RETURNCONFLICT = 10086
	NFS4ERR_DELEG_REVOKED = 10087
	SET_TO_SERVER_TIME4 = 0
	SET_TO_CLIENT_TIME4 = 1
	LAYOUT4_NFSV4_1_FILES = 1
	LAYOUT4_OSD2_OBJECTS = 2
	LAYOUT4_BLOCK_VOLUME = 3
	LAYOUTIOMODE4_READ = 1
	LAYOUTIOMODE4_RW = 2
	LAYOUTIOMODE4_ANY = 3
	LAYOUTRETURN4_FILE = 1
	LAYOUTRETURN4_FSID = 2
	LAYOUTRETURN4_ALL = 3
	STATUS4_FIXED = 1
	STATUS4_UPDATED = 2
	STATUS4_VERSIONED = 3
	STATUS4_WRITABLE = 4
	STATUS4_REFERRAL = 5
	READ_LT = 1
	WRITE_LT = 2
	READW_LT = 3
	WRITEW_LT = 4
	SSV4_SUBKEY_MIC_I2T = 1
	SSV4_SUBKEY_MIC_T2I = 2
	SSV4_SUBKEY_SEAL_I2T = 3
	SSV4_SUBKEY_SEAL_T2I = 4
	NFLH4_CARE_DENSE = 1
	NFLH4_CARE_COMMIT_THRU_MDS = 2
	NFLH4_CARE_STRIPE_UNIT_SIZE = 64
	NFLH4_CARE_STRIPE_COUNT = 128
	UNCHECKED4 = 0
	GUARDED4 = 1
	EXCLUSIVE4 = 2
	EXCLUSIVE4_1 = 3
	OPEN4_NOCREATE = 0
	OPEN4_CREATE = 1
	NFS_LIMIT_SIZE = 1
	NFS_LIMIT_BLOCKS = 2
	OPEN_DELEGATE_NONE = 0
	OPEN_DELEGATE_READ = 1
	OPEN_DELEGATE_WRITE = 2
	OPEN_DELEGATE_NONE_EXT = 3
	CLAIM_NULL = 0
	CLAIM_PREVIOUS = 1
	CLAIM_DELEGATE_CUR = 2
	CLAIM_DELEGATE_PREV = 3
	CLAIM_FH = 4
	CLAIM_DELEG_CUR_FH = 5
	CLAIM_DELEG_PREV_FH = 6
	WND4_NOT_WANTED = 0
	WND4_CONTENTION = 1
	WND4_RESOURCE = 2
	WND4_NOT_SUPP_FTYPE = 3
	WND4_WRITE_DELEG_NOT_SUPP_FTYPE = 4
	WND4_NOT_SUPP_UPGRADE = 5
	WND4_NOT_SUPP_DOWNGRADE = 6
	WND4_CANCELLED = 7
	WND4_IS_DIR = 8
	RPC_GSS_SVC_NONE = 1
	RPC_GSS_SVC_INTEGRITY = 2
	RPC_GSS_SVC_PRIVACY = 3
	UNSTABLE4 = 0
	DATA_SYNC4 = 1
	FILE_SYNC4 = 2
	CDFC4_FORE = 1
	CDFC4_BACK = 2
	CDFC4_FORE_OR_BOTH = 3
	CDFC4_BACK_OR_BOTH = 7
	CDFS4_FORE = 1
	CDFS4_BACK = 2
	CDFS4_BOTH = 3
	SP4_NONE = 0
	SP4_MACH_CRED = 1
	SP4_SSV = 2
	GDD4_OK = 0
	GDD4_UNAVAIL = 1
	SECINFO_STYLE4_CURRENT_FH = 0
	SECINFO_STYLE4_PARENT = 1
	OP_ACCESS = 3
	OP_CLOSE = 4
	OP_COMMIT = 5
	OP_CREATE = 6
	OP_DELEGPURGE = 7
	OP_DELEGRETURN = 8
	OP_GETATTR = 9
	OP_GETFH = 10
	OP_LINK = 11
	OP_LOCK = 12
	OP_LOCKT = 13
	OP_LOCKU = 14
	OP_LOOKUP = 15
	OP_LOOKUPP = 16
	OP_NVERIFY = 17
	OP_OPEN = 18
	OP_OPENATTR = 19
	OP_OPEN_CONFIRM = 20
	OP_OPEN_DOWNGRADE = 21
	OP_PUTFH = 22
	OP_PUTPUBFH = 23
	OP_PUTROOTFH = 24
	OP_READ = 25
	OP_READDIR = 26
	OP_READLINK = 27
	OP_REMOVE = 28
	OP_RENAME = 29
	OP_RENEW = 30
	OP_RESTOREFH = 31
	OP_SAVEFH = 32
	OP_SECINFO = 33
	OP_SETATTR = 34
	OP_SETCLIENTID = 35
	OP_SETCLIENTID_CONFIRM = 36
	OP_VERIFY = 37
	OP_WRITE = 38
	OP_RELEASE_LOCKOWNER = 39
	OP_BACKCHANNEL_CTL = 40
	OP_BIND_CONN_TO_SESSION = 41
	OP_EXCHANGE_ID = 42
	OP_CREATE_SESSION = 43
	OP_DESTROY_SESSION = 44
	OP_FREE_STATEID = 45
	OP_GET_DIR_DELEGATION = 46
	OP_GETDEVICEINFO = 47
	OP_GETDEVICELIST = 48
	OP_LAYOUTCOMMIT = 49
	OP_LAYOUTGET = 50
	OP_LAYOUTRETURN = 51
	OP_SECINFO_NO_NAME = 52
	OP_SEQUENCE = 53
	OP_SET_SSV = 54
	OP_TEST_STATEID = 55
	OP_WANT_DELEGATION = 56
	OP_DESTROY_CLIENTID = 57
	OP_RECLAIM_COMPLETE = 58
	OP_ILLEGAL = 10044
	LAYOUTRECALL4_FILE = 1
	LAYOUTRECALL4_FSID = 2
	LAYOUTRECALL4_ALL = 3
	NOTIFY4_CHANGE_CHILD_ATTRS = 0
	NOTIFY4_CHANGE_DIR_ATTRS = 1
	NOTIFY4_REMOVE_ENTRY = 2
	NOTIFY4_ADD_ENTRY = 3
	NOTIFY4_RENAME_ENTRY = 4
	NOTIFY4_CHANGE_COOKIE_VERIFIER = 5
	NOTIFY_DEVICEID4_CHANGE = 1
	NOTIFY_DEVICEID4_DELETE = 2
	OP_CB_GETATTR = 3
	OP_CB_RECALL = 4
	OP_CB_LAYOUTRECALL = 5
	OP_CB_NOTIFY = 6
	OP_CB_PUSH_DELEG = 7
	OP_CB_RECALL_ANY = 8
	OP_CB_RECALLABLE_OBJ_AVAIL = 9
	OP_CB_RECALL_SLOT = 10
	OP_CB_SEQUENCE = 11
	OP_CB_WANTS_CANCELLED = 12
	OP_CB_NOTIFY_LOCK = 13
	OP_CB_NOTIFY_DEVICEID = 14
	OP_CB_ILLEGAL = 10044

	/* Consts: */
	NFS4_FHSIZE = 128
	NFS4_VERIFIER_SIZE = 8
	NFS4_OPAQUE_LIMIT = 1024
	NFS4_SESSIONID_SIZE = 16
	NFS4_INT64_MAX = 0x7fffffffffffffff
	NFS4_UINT64_MAX = 0xffffffffffffffff
	NFS4_INT32_MAX = 0x7fffffff
	NFS4_UINT32_MAX = 0xffffffff
	NFS4_MAXFILELEN = 0xffffffffffffffff
	NFS4_MAXFILEOFF = 0xfffffffffffffffe
	ACL4_SUPPORT_ALLOW_ACL = 0x00000001
	ACL4_SUPPORT_DENY_ACL = 0x00000002
	ACL4_SUPPORT_AUDIT_ACL = 0x00000004
	ACL4_SUPPORT_ALARM_ACL = 0x00000008
	ACE4_ACCESS_ALLOWED_ACE_TYPE = 0x00000000
	ACE4_ACCESS_DENIED_ACE_TYPE = 0x00000001
	ACE4_SYSTEM_AUDIT_ACE_TYPE = 0x00000002
	ACE4_SYSTEM_ALARM_ACE_TYPE = 0x00000003
	ACE4_FILE_INHERIT_ACE = 0x00000001
	ACE4_DIRECTORY_INHERIT_ACE = 0x00000002
	ACE4_NO_PROPAGATE_INHERIT_ACE = 0x00000004
	ACE4_INHERIT_ONLY_ACE = 0x00000008
	ACE4_SUCCESSFUL_ACCESS_ACE_FLAG = 0x00000010
	ACE4_FAILED_ACCESS_ACE_FLAG = 0x00000020
	ACE4_IDENTIFIER_GROUP = 0x00000040
	ACE4_INHERITED_ACE = 0x00000080
	ACE4_READ_DATA = 0x00000001
	ACE4_LIST_DIRECTORY = 0x00000001
	ACE4_WRITE_DATA = 0x00000002
	ACE4_ADD_FILE = 0x00000002
	ACE4_APPEND_DATA = 0x00000004
	ACE4_ADD_SUBDIRECTORY = 0x00000004
	ACE4_READ_NAMED_ATTRS = 0x00000008
	ACE4_WRITE_NAMED_ATTRS = 0x00000010
	ACE4_EXECUTE = 0x00000020
	ACE4_DELETE_CHILD = 0x00000040
	ACE4_READ_ATTRIBUTES = 0x00000080
	ACE4_WRITE_ATTRIBUTES = 0x00000100
	ACE4_WRITE_RETENTION = 0x00000200
	ACE4_WRITE_RETENTION_HOLD = 0x00000400
	ACE4_DELETE = 0x00010000
	ACE4_READ_ACL = 0x00020000
	ACE4_WRITE_ACL = 0x00040000
	ACE4_WRITE_OWNER = 0x00080000
	ACE4_SYNCHRONIZE = 0x00100000
	ACE4_GENERIC_READ = 0x00120081
	ACE4_GENERIC_WRITE = 0x00160106
	ACE4_GENERIC_EXECUTE = 0x001200A0
	ACL4_AUTO_INHERIT = 0x00000001
	ACL4_PROTECTED = 0x00000002
	ACL4_DEFAULTED = 0x00000004
	MODE4_SUID = 0x800
	MODE4_SGID = 0x400
	MODE4_SVTX = 0x200
	MODE4_RUSR = 0x100
	MODE4_WUSR = 0x080
	MODE4_XUSR = 0x040
	MODE4_RGRP = 0x020
	MODE4_WGRP = 0x010
	MODE4_XGRP = 0x008
	MODE4_ROTH = 0x004
	MODE4_WOTH = 0x002
	MODE4_XOTH = 0x001
	FH4_PERSISTENT = 0x00000000
	FH4_NOEXPIRE_WITH_OPEN = 0x00000001
	FH4_VOLATILE_ANY = 0x00000002
	FH4_VOL_MIGRATION = 0x00000004
	FH4_VOL_RENAME = 0x00000008
	NFS4_DEVICEID4_SIZE = 16
	LAYOUT4_RET_REC_FILE = 1
	LAYOUT4_RET_REC_FSID = 2
	LAYOUT4_RET_REC_ALL = 3
	TH4_READ_SIZE = 0
	TH4_WRITE_SIZE = 1
	TH4_READ_IOSIZE = 2
	TH4_WRITE_IOSIZE = 3
	RET4_DURATION_INFINITE = 0xffffffffffffffff
	FSCHARSET_CAP4_CONTAINS_NON_UTF8 = 0x1
	FSCHARSET_CAP4_ALLOWS_ONLY_UTF8 = 0x2
	FATTR4_SUPPORTED_ATTRS = 0
	FATTR4_TYPE = 1
	FATTR4_FH_EXPIRE_TYPE = 2
	FATTR4_CHANGE = 3
	FATTR4_SIZE = 4
	FATTR4_LINK_SUPPORT = 5
	FATTR4_SYMLINK_SUPPORT = 6
	FATTR4_NAMED_ATTR = 7
	FATTR4_FSID = 8
	FATTR4_UNIQUE_HANDLES = 9
	FATTR4_LEASE_TIME = 10
	FATTR4_RDATTR_ERROR = 11
	FATTR4_FILEHANDLE = 19
	FATTR4_SUPPATTR_EXCLCREAT = 75
	FATTR4_ACL = 12
	FATTR4_ACLSUPPORT = 13
	FATTR4_ARCHIVE = 14
	FATTR4_CANSETTIME = 15
	FATTR4_CASE_INSENSITIVE = 16
	FATTR4_CASE_PRESERVING = 17
	FATTR4_CHOWN_RESTRICTED = 18
	FATTR4_FILEID = 20
	FATTR4_FILES_AVAIL = 21
	FATTR4_FILES_FREE = 22
	FATTR4_FILES_TOTAL = 23
	FATTR4_FS_LOCATIONS = 24
	FATTR4_HIDDEN = 25
	FATTR4_HOMOGENEOUS = 26
	FATTR4_MAXFILESIZE = 27
	FATTR4_MAXLINK = 28
	FATTR4_MAXNAME = 29
	FATTR4_MAXREAD = 30
	FATTR4_MAXWRITE = 31
	FATTR4_MIMETYPE = 32
	FATTR4_MODE = 33
	FATTR4_NO_TRUNC = 34
	FATTR4_NUMLINKS = 35
	FATTR4_OWNER = 36
	FATTR4_OWNER_GROUP = 37
	FATTR4_QUOTA_AVAIL_HARD = 38
	FATTR4_QUOTA_AVAIL_SOFT = 39
	FATTR4_QUOTA_USED = 40
	FATTR4_RAWDEV = 41
	FATTR4_SPACE_AVAIL = 42
	FATTR4_SPACE_FREE = 43
	FATTR4_SPACE_TOTAL = 44
	FATTR4_SPACE_USED = 45
	FATTR4_SYSTEM = 46
	FATTR4_TIME_ACCESS = 47
	FATTR4_TIME_ACCESS_SET = 48
	FATTR4_TIME_BACKUP = 49
	FATTR4_TIME_CREATE = 50
	FATTR4_TIME_DELTA = 51
	FATTR4_TIME_METADATA = 52
	FATTR4_TIME_MODIFY = 53
	FATTR4_TIME_MODIFY_SET = 54
	FATTR4_MOUNTED_ON_FILEID = 55
	FATTR4_DIR_NOTIF_DELAY = 56
	FATTR4_DIRENT_NOTIF_DELAY = 57
	FATTR4_DACL = 58
	FATTR4_SACL = 59
	FATTR4_CHANGE_POLICY = 60
	FATTR4_FS_STATUS = 61
	FATTR4_FS_LAYOUT_TYPES = 62
	FATTR4_LAYOUT_HINT = 63
	FATTR4_LAYOUT_TYPES = 64
	FATTR4_LAYOUT_BLKSIZE = 65
	FATTR4_LAYOUT_ALIGNMENT = 66
	FATTR4_FS_LOCATIONS_INFO = 67
	FATTR4_MDSTHRESHOLD = 68
	FATTR4_RETENTION_GET = 69
	FATTR4_RETENTION_SET = 70
	FATTR4_RETENTEVT_GET = 71
	FATTR4_RETENTEVT_SET = 72
	FATTR4_RETENTION_HOLD = 73
	FATTR4_MODE_SET_MASKED = 74
	FATTR4_FS_CHARSET_CAP = 76
	FSLI4BX_GFLAGS = 0
	FSLI4BX_TFLAGS = 1
	FSLI4BX_CLSIMUL = 2
	FSLI4BX_CLHANDLE = 3
	FSLI4BX_CLFILEID = 4
	FSLI4BX_CLWRITEVER = 5
	FSLI4BX_CLCHANGE = 6
	FSLI4BX_CLREADDIR = 7
	FSLI4BX_READRANK = 8
	FSLI4BX_WRITERANK = 9
	FSLI4BX_READORDER = 10
	FSLI4BX_WRITEORDER = 11
	FSLI4GF_WRITABLE = 0x01
	FSLI4GF_CUR_REQ = 0x02
	FSLI4GF_ABSENT = 0x04
	FSLI4GF_GOING = 0x08
	FSLI4GF_SPLIT = 0x10
	FSLI4TF_RDMA = 0x01
	FSLI4IF_VAR_SUB = 0x00000001
	NFL4_UFLG_MASK = 0x0000003F
	NFL4_UFLG_DENSE = 0x00000001
	NFL4_UFLG_COMMIT_THRU_MDS = 0x00000002
	NFL4_UFLG_STRIPE_UNIT_SIZE_MASK = 0xFFFFFFC0
	ACCESS4_READ = 0x00000001
	ACCESS4_LOOKUP = 0x00000002
	ACCESS4_MODIFY = 0x00000004
	ACCESS4_EXTEND = 0x00000008
	ACCESS4_DELETE = 0x00000010
	ACCESS4_EXECUTE = 0x00000020
	OPEN4_SHARE_ACCESS_READ = 0x00000001
	OPEN4_SHARE_ACCESS_WRITE = 0x00000002
	OPEN4_SHARE_ACCESS_BOTH = 0x00000003
	OPEN4_SHARE_DENY_NONE = 0x00000000
	OPEN4_SHARE_DENY_READ = 0x00000001
	OPEN4_SHARE_DENY_WRITE = 0x00000002
	OPEN4_SHARE_DENY_BOTH = 0x00000003
	OPEN4_SHARE_ACCESS_WANT_DELEG_MASK = 0xFF00
	OPEN4_SHARE_ACCESS_WANT_NO_PREFERENCE = 0x0000
	OPEN4_SHARE_ACCESS_WANT_READ_DELEG = 0x0100
	OPEN4_SHARE_ACCESS_WANT_WRITE_DELEG = 0x0200
	OPEN4_SHARE_ACCESS_WANT_ANY_DELEG = 0x0300
	OPEN4_SHARE_ACCESS_WANT_NO_DELEG = 0x0400
	OPEN4_SHARE_ACCESS_WANT_CANCEL = 0x0500
	OPEN4_SHARE_ACCESS_WANT_SIGNAL_DELEG_WHEN_RESRC_AVAIL = 0x10000
	OPEN4_SHARE_ACCESS_WANT_PUSH_DELEG_WHEN_UNCONTENDED = 0x20000
	OPEN4_RESULT_CONFIRM = 0x00000002
	OPEN4_RESULT_LOCKTYPE_POSIX = 0x00000004
	OPEN4_RESULT_PRESERVE_UNLINKED = 0x00000008
	OPEN4_RESULT_MAY_NOTIFY_LOCK = 0x00000020
	EXCHGID4_FLAG_SUPP_MOVED_REFER = 0x00000001
	EXCHGID4_FLAG_SUPP_MOVED_MIGR = 0x00000002
	EXCHGID4_FLAG_BIND_PRINC_STATEID = 0x00000100
	EXCHGID4_FLAG_USE_NON_PNFS = 0x00010000
	EXCHGID4_FLAG_USE_PNFS_MDS = 0x00020000
	EXCHGID4_FLAG_USE_PNFS_DS = 0x00040000
	EXCHGID4_FLAG_MASK_PNFS = 0x00070000
	EXCHGID4_FLAG_UPD_CONFIRMED_REC_A = 0x40000000
	EXCHGID4_FLAG_CONFIRMED_R = 0x80000000
	CREATE_SESSION4_FLAG_PERSIST = 0x00000001
	CREATE_SESSION4_FLAG_CONN_BACK_CHAN = 0x00000002
	CREATE_SESSION4_FLAG_CONN_RDMA = 0x00000004
	SEQ4_STATUS_CB_PATH_DOWN = 0x00000001
	SEQ4_STATUS_CB_GSS_CONTEXTS_EXPIRING = 0x00000002
	SEQ4_STATUS_CB_GSS_CONTEXTS_EXPIRED = 0x00000004
	SEQ4_STATUS_EXPIRED_ALL_STATE_REVOKED = 0x00000008
	SEQ4_STATUS_EXPIRED_SOME_STATE_REVOKED = 0x00000010
	SEQ4_STATUS_ADMIN_STATE_REVOKED = 0x00000020
	SEQ4_STATUS_RECALLABLE_STATE_REVOKED = 0x00000040
	SEQ4_STATUS_LEASE_MOVED = 0x00000080
	SEQ4_STATUS_RESTART_RECLAIM_NEEDED = 0x00000100
	SEQ4_STATUS_CB_PATH_DOWN_SESSION = 0x00000200
	SEQ4_STATUS_BACKCHANNEL_FAULT = 0x00000400
	SEQ4_STATUS_DEVID_CHANGED = 0x00000800
	SEQ4_STATUS_DEVID_DELETED = 0x00001000
	RCA4_TYPE_MASK_RDATA_DLG = 0
	RCA4_TYPE_MASK_WDATA_DLG = 1
	RCA4_TYPE_MASK_DIR_DLG = 2
	RCA4_TYPE_MASK_FILE_LAYOUT = 3
	RCA4_TYPE_MASK_BLK_LAYOUT = 4
	RCA4_TYPE_MASK_OBJ_LAYOUT_MIN = 8
	RCA4_TYPE_MASK_OBJ_LAYOUT_MAX = 9
	RCA4_TYPE_MASK_OTHER_LAYOUT_MIN = 12
	RCA4_TYPE_MASK_OTHER_LAYOUT_MAX = 15

)


type NfsFh4 []byte // nfs_fh4
type Sessionid4 [NFS4_SESSIONID_SIZE]byte // sessionid4
type Verifier4 [NFS4_VERIFIER_SIZE]byte // verifier4
type Deviceid4 [NFS4_DEVICEID4_SIZE]byte // deviceid4
type Fattr4Fsid Fsid4 // fattr4_fsid
type Fattr4Acl []Nfsace4 // fattr4_acl
type Fattr4FsLocations FsLocations4 // fattr4_fs_locations
type Fattr4ModeSetMasked ModeMasked4 // fattr4_mode_set_masked
type Fattr4Rawdev Specdata4 // fattr4_rawdev
type Fattr4TimeAccess Nfstime4 // fattr4_time_access
type Fattr4TimeAccessSet Settime4 // fattr4_time_access_set
type Fattr4TimeBackup Nfstime4 // fattr4_time_backup
type Fattr4TimeCreate Nfstime4 // fattr4_time_create
type Fattr4TimeDelta Nfstime4 // fattr4_time_delta
type Fattr4TimeMetadata Nfstime4 // fattr4_time_metadata
type Fattr4TimeModify Nfstime4 // fattr4_time_modify
type Fattr4TimeModifySet Settime4 // fattr4_time_modify_set
type Fattr4DirNotifDelay Nfstime4 // fattr4_dir_notif_delay
type Fattr4DirentNotifDelay Nfstime4 // fattr4_dirent_notif_delay
type Fattr4FsStatus Fs4Status // fattr4_fs_status
type Fattr4LayoutHint Layouthint4 // fattr4_layout_hint
type Fattr4Mdsthreshold Mdsthreshold4 // fattr4_mdsthreshold
type Fattr4RetentionGet RetentionGet4 // fattr4_retention_get
type Fattr4RetentionSet RetentionSet4 // fattr4_retention_set
type Fattr4RetentevtGet RetentionGet4 // fattr4_retentevt_get
type Fattr4RetentevtSet RetentionSet4 // fattr4_retentevt_set
type Fattr4Dacl Nfsacl41 // fattr4_dacl
type Fattr4Sacl Nfsacl41 // fattr4_sacl
type Clientaddr4 Netaddr4 // clientaddr4
type OpenOwner4 StateOwner4 // open_owner4
type LockOwner4 StateOwner4 // lock_owner4
type Fattr4FsLocationsInfo FsLocationsInfo4 // fattr4_fs_locations_info
type MultipathList4 []Netaddr4 // multipath_list4
type SECINFO4resok []Secinfo4 // SECINFO4resok
type AttrNotice4 Nfstime4 // attr_notice4
type SecinfoNoName4res SECINFO4res // SECINFO_NO_NAME4res
type CbRecallableObjAvail4args CbRecallAny4args // CB_RECALLABLE_OBJ_AVAIL4args



const NFS4_PROGRAM = 100003
const NFS_V4 = 4
const NFSPROC4_COMPOUND = 1
const NFSPROC4_NULL = 0



type Nfstime4 struct { // nfstime4
	Seconds int64
	Nseconds uint32
}


type Fsid4 struct { // fsid4
	Major uint64
	Minor uint64
}


type ChangePolicy4 struct { // change_policy4
	CpMajor uint64
	CpMinor uint64
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


type Nfsacl41 struct { // nfsacl41
	Na41Flag uint32
	Na41Aces []Nfsace4
}


type ModeMasked4 struct { // mode_masked4
	MmValueToSet uint32
	MmMaskBits uint32
}


type Specdata4 struct { // specdata4
	Specdata1 uint32
	Specdata2 uint32
}


type Netaddr4 struct { // netaddr4
	NaRNetid string
	NaRAddr string
}


type NfsImplID4 struct { // nfs_impl_id4
	NiiDomain string
	NiiName string
	NiiDate Nfstime4
}


type Stateid4 struct { // stateid4
	Seqid uint32
	Other [12]byte
}


type LayoutContent4 struct { // layout_content4
	LocType int32
	LocBody []byte
}


type Layouthint4 struct { // layouthint4
	LohType int32
	LohBody []byte
}


type Layout4 struct { // layout4
	LoOffset uint64
	LoLength uint64
	LoIomode int32
	LoContent LayoutContent4
}


type DeviceAddr4 struct { // device_addr4
	DaLayoutType int32
	DaAddrBody []byte
}


type Layoutupdate4 struct { // layoutupdate4
	LouType int32
	LouBody []byte
}


type LayoutreturnFile4 struct { // layoutreturn_file4
	LrfOffset uint64
	LrfLength uint64
	LrfStateid Stateid4
	LrfBody []byte
}


type Fs4Status struct { // fs4_status
	FssAbsent bool
	FssType int32
	FssSource string
	FssCurrent string
	FssAge int32
	FssVersion Nfstime4
}


type ThresholdItem4 struct { // threshold_item4
	ThiLayoutType int32
	ThiHintset []uint32
	ThiHintlist []byte
}


type Mdsthreshold4 struct { // mdsthreshold4
	MthHints []ThresholdItem4
}


type RetentionGet4 struct { // retention_get4
	RgDuration uint64
	RgBeginTime []Nfstime4
}


type RetentionSet4 struct { // retention_set4
	RsEnable bool
	RsDuration []uint64
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


type CbClient4 struct { // cb_client4
	CbProgram uint32
	CbLocation Netaddr4
}


type NfsClientID4 struct { // nfs_client_id4
	Verifier Verifier4
	ID string
}


type ClientOwner4 struct { // client_owner4
	CoVerifier Verifier4
	CoOwnerid string
}


type ServerOwner4 struct { // server_owner4
	SoMinorID uint64
	SoMajorID string
}


type StateOwner4 struct { // state_owner4
	Clientid uint64
	Owner string
}


type SsvMicPlainTkn4 struct { // ssv_mic_plain_tkn4
	SmptSsvSeq uint32
	SmptOrigPlain string
}


type SsvMicTkn4 struct { // ssv_mic_tkn4
	SmtSsvSeq uint32
	SmtHmac []byte
}


type SsvSealPlainTkn4 struct { // ssv_seal_plain_tkn4
	SsptConfounder []byte
	SsptSsvSeq uint32
	SsptOrigPlain []byte
	SsptPad []byte
}


type SsvSealCipherTkn4 struct { // ssv_seal_cipher_tkn4
	SsctSsvSeq uint32
	SsctIv []byte
	SsctEncrData []byte
	SsctHmac []byte
}


type FsLocationsServer4 struct { // fs_locations_server4
	FlsCurrency int32
	FlsInfo []byte
	FlsServer string
}


type FsLocationsItem4 struct { // fs_locations_item4
	FliEntries []FsLocationsServer4
	FliRootpath []string
}


type FsLocationsInfo4 struct { // fs_locations_info4
	FliFlags uint32
	FliValidFor int32
	FliFsRoot []string
	FliItems []FsLocationsItem4
}


type Nfsv41FileLayouthint4 struct { // nfsv4_1_file_layouthint4
	NflhCare uint32
	NflhUtil uint32
	NflhStripeCount uint32
}


type Nfsv41FileLayoutDsAddr4 struct { // nfsv4_1_file_layout_ds_addr4
	NfldaStripeIndices []uint32
	NfldaMultipathDsList []MultipathList4
}


type Nfsv41FileLayout4 struct { // nfsv4_1_file_layout4
	NflDeviceid Deviceid4
	NflUtil uint32
	NflFirstStripeIndex uint32
	NflPatternOffset uint64
	NflFhList []NfsFh4
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
	Offset uint64
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
	Offset uint64
	Length uint64
	Locker Locker4
}


type LOCK4denied struct { // LOCK4denied
	Offset uint64
	Length uint64
	Locktype int32
	Owner LockOwner4
}


type LOCK4resok struct { // LOCK4resok
	LockStateid Stateid4
}


type LOCKT4args struct { // LOCKT4args
	Locktype int32
	Offset uint64
	Length uint64
	Owner LockOwner4
}


type LOCKU4args struct { // LOCKU4args
	Locktype int32
	Seqid uint32
	LockStateid Stateid4
	Offset uint64
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


type Creatverfattr struct { // creatverfattr
	CvaVerf Verifier4
	CvaAttrs Fattr4
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
	Offset uint64
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
	Offset uint64
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


type GssCbHandles4 struct { // gss_cb_handles4
	GcbpService int32
	GcbpHandleFromServer []byte
	GcbpHandleFromClient []byte
}


type BackchannelCtl4args struct { // BACKCHANNEL_CTL4args
	BcaCbProgram uint32
	BcaSecParms []CallbackSecParms4
}


type BackchannelCtl4res struct { // BACKCHANNEL_CTL4res
	BcrStatus int32
}


type BindConnToSession4args struct { // BIND_CONN_TO_SESSION4args
	BctsaSessid Sessionid4
	BctsaDir int32
	BctsaUseConnInRdmaMode bool
}


type BindConnToSession4resok struct { // BIND_CONN_TO_SESSION4resok
	BctsrSessid Sessionid4
	BctsrDir int32
	BctsrUseConnInRdmaMode bool
}


type StateProtectOps4 struct { // state_protect_ops4
	SpoMustEnforce []uint32
	SpoMustAllow []uint32
}


type SsvSpParms4 struct { // ssv_sp_parms4
	SspOps StateProtectOps4
	SspHashAlgs [][]byte
	SspEncrAlgs [][]byte
	SspWindow uint32
	SspNumGssHandles uint32
}


type ExchangeID4args struct { // EXCHANGE_ID4args
	EiaClientowner ClientOwner4
	EiaFlags uint32
	EiaStateProtect StateProtect4A
	EiaClientImplID []NfsImplID4
}


type SsvProtInfo4 struct { // ssv_prot_info4
	SpiOps StateProtectOps4
	SpiHashAlg uint32
	SpiEncrAlg uint32
	SpiSsvLen uint32
	SpiWindow uint32
	SpiHandles [][]byte
}


type ExchangeID4resok struct { // EXCHANGE_ID4resok
	EirClientid uint64
	EirSequenceid uint32
	EirFlags uint32
	EirStateProtect StateProtect4R
	EirServerOwner ServerOwner4
	EirServerScope []byte
	EirServerImplID []NfsImplID4
}

/*
ca_maxrequestsize:

The maximum size of a COMPOUND or CB_COMPOUND request that will
be sent.  This size represents the XDR encoded size of the
request, including the RPC headers (including security flavor
credentials and verifiers) but excludes any RPC transport
framing headers.
If a requester sends a
request that exceeds ca_maxrequestsize, the error
NFS4ERR_REQ_TOO_BIG will be returned

ca_maxoperations:

The maximum number of operations the replier will accept in a
COMPOUND or CB_COMPOUND.  For the backchannel, the server MUST
NOT change the value the client offers.  For the fore channel,
the server MAY change the requested value.  After the session
is created, if a requester sends a COMPOUND or CB_COMPOUND with
more operations than ca_maxoperations, the replier MUST return
NFS4ERR_TOO_MANY_OPS.

*/

type ChannelAttrs4 struct { // channel_attrs4
	CaHeaderpadsize uint32
	CaMaxrequestsize uint32
	CaMaxresponsesize uint32
	CaMaxresponsesizeCached uint32
	CaMaxoperations uint32
	CaMaxrequests uint32
	CaRdmaIrd []uint32
}


type CreateSession4args struct { // CREATE_SESSION4args
	CsaClientid uint64
	CsaSequence uint32
	CsaFlags uint32
	CsaForeChanAttrs ChannelAttrs4
	CsaBackChanAttrs ChannelAttrs4
	CsaCbProgram uint32
	CsaSecParms []CallbackSecParms4
}


type CreateSession4resok struct { // CREATE_SESSION4resok
	CsrSessionid Sessionid4
	CsrSequence uint32
	CsrFlags uint32
	CsrForeChanAttrs ChannelAttrs4
	CsrBackChanAttrs ChannelAttrs4
}


type DestroySession4args struct { // DESTROY_SESSION4args
	DsaSessionid Sessionid4
}


type DestroySession4res struct { // DESTROY_SESSION4res
	DsrStatus int32
}


type FreeStateid4args struct { // FREE_STATEID4args
	FsaStateid Stateid4
}


type FreeStateid4res struct { // FREE_STATEID4res
	FsrStatus int32
}


type GetDirDelegation4args struct { // GET_DIR_DELEGATION4args
	GddaSignalDelegAvail bool
	GddaNotificationTypes []uint32
	GddaChildAttrDelay AttrNotice4
	GddaDirAttrDelay AttrNotice4
	GddaChildAttributes []uint32
	GddaDirAttributes []uint32
}


type GetDirDelegation4resok struct { // GET_DIR_DELEGATION4resok
	GddrCookieverf Verifier4
	GddrStateid Stateid4
	GddrNotification []uint32
	GddrChildAttributes []uint32
	GddrDirAttributes []uint32
}


type GETDEVICEINFO4args struct { // GETDEVICEINFO4args
	GdiaDeviceID Deviceid4
	GdiaLayoutType int32
	GdiaMaxcount uint32
	GdiaNotifyTypes []uint32
}


type GETDEVICEINFO4resok struct { // GETDEVICEINFO4resok
	GdirDeviceAddr DeviceAddr4
	GdirNotification []uint32
}


type GETDEVICELIST4args struct { // GETDEVICELIST4args
	GdlaLayoutType int32
	GdlaMaxdevices uint32
	GdlaCookie uint64
	GdlaCookieverf Verifier4
}


type GETDEVICELIST4resok struct { // GETDEVICELIST4resok
	GdlrCookie uint64
	GdlrCookieverf Verifier4
	GdlrDeviceidList []Deviceid4
	GdlrEof bool
}


type LAYOUTCOMMIT4args struct { // LAYOUTCOMMIT4args
	LocaOffset uint64
	LocaLength uint64
	LocaReclaim bool
	LocaStateid Stateid4
	LocaLastWriteOffset Newoffset4
	LocaTimeModify Newtime4
	LocaLayoutupdate Layoutupdate4
}


type LAYOUTCOMMIT4resok struct { // LAYOUTCOMMIT4resok
	LocrNewsize Newsize4
}


type LAYOUTGET4args struct { // LAYOUTGET4args
	LogaSignalLayoutAvail bool
	LogaLayoutType int32
	LogaIomode int32
	LogaOffset uint64
	LogaLength uint64
	LogaMinlength uint64
	LogaStateid Stateid4
	LogaMaxcount uint32
}


type LAYOUTGET4resok struct { // LAYOUTGET4resok
	LogrReturnOnClose bool
	LogrStateid Stateid4
	LogrLayout []Layout4
}


type LAYOUTRETURN4args struct { // LAYOUTRETURN4args
	LoraReclaim bool
	LoraLayoutType int32
	LoraIomode int32
	LoraLayoutreturn Layoutreturn4
}


type SEQUENCE4args struct { // SEQUENCE4args
	SaSessionid Sessionid4
	SaSequenceid uint32
	SaSlotid uint32
	SaHighestSlotid uint32
	SaCachethis bool
}


type SEQUENCE4resok struct { // SEQUENCE4resok
	SrSessionid Sessionid4
	SrSequenceid uint32
	SrSlotid uint32
	SrHighestSlotid uint32
	SrTargetHighestSlotid uint32
	SrStatusFlags uint32
}


type SsaDigestInput4 struct { // ssa_digest_input4
	SdiSeqargs SEQUENCE4args
}


type SetSsv4args struct { // SET_SSV4args
	SsaSsv []byte
	SsaDigest []byte
}


type SsrDigestInput4 struct { // ssr_digest_input4
	SdiSeqres SEQUENCE4res
}


type SetSsv4resok struct { // SET_SSV4resok
	SsrDigest []byte
}


type TestStateid4args struct { // TEST_STATEID4args
	TsStateids []Stateid4
}


type TestStateid4resok struct { // TEST_STATEID4resok
	TsrStatusCodes []int32
}


type WantDelegation4args struct { // WANT_DELEGATION4args
	WdaWant uint32
	WdaClaim DelegClaim4
}


type Destroyuint64args struct { // DESTROY_CLIENTID4args
	DcaClientid uint64
}


type Destroyuint64res struct { // DESTROY_CLIENTID4res
	DcrStatus int32
}


type ReclaimComplete4args struct { // RECLAIM_COMPLETE4args
	RcaOneFs bool
}


type ReclaimComplete4res struct { // RECLAIM_COMPLETE4res
	RcrStatus int32
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


type LayoutrecallFile4 struct { // layoutrecall_file4
	LorFh NfsFh4
	LorOffset uint64
	LorLength uint64
	LorStateid Stateid4
}


type CbLayoutrecall4args struct { // CB_LAYOUTRECALL4args
	CloraType int32
	CloraIomode int32
	CloraChanged bool
	CloraRecall Layoutrecall4
}


type CbLayoutrecall4res struct { // CB_LAYOUTRECALL4res
	ClorrStatus int32
}


type NotifyEntry4 struct { // notify_entry4
	NeFile string
	NeAttrs Fattr4
}


type PrevEntry4 struct { // prev_entry4
	PePrevEntry NotifyEntry4
	PePrevEntryCookie uint64
}


type NotifyRemove4 struct { // notify_remove4
	NrmOldEntry NotifyEntry4
	NrmOldEntryCookie uint64
}


type NotifyAdd4 struct { // notify_add4
	NadOldEntry []NotifyRemove4
	NadNewEntry NotifyEntry4
	NadNewEntryCookie []uint64
	NadPrevEntry []PrevEntry4
	NadLastEntry bool
}


type NotifyAttr4 struct { // notify_attr4
	NaChangedEntry NotifyEntry4
}


type NotifyRename4 struct { // notify_rename4
	NrnOldEntry NotifyRemove4
	NrnNewEntry NotifyAdd4
}


type NotifyVerifier4 struct { // notify_verifier4
	NvOldCookieverf Verifier4
	NvNewCookieverf Verifier4
}


type Notify4 struct { // notify4
	NotifyMask []uint32
	NotifyVals []byte
}


type CbNotify4args struct { // CB_NOTIFY4args
	CnaStateid Stateid4
	CnaFh NfsFh4
	CnaChanges []Notify4
}


type CbNotify4res struct { // CB_NOTIFY4res
	CnrStatus int32
}


type CbPushDeleg4args struct { // CB_PUSH_DELEG4args
	CpdaFh NfsFh4
	CpdaDelegation OpenDelegation4
}


type CbPushDeleg4res struct { // CB_PUSH_DELEG4res
	CpdrStatus int32
}


type CbRecallAny4args struct { // CB_RECALL_ANY4args
	CraaObjectsToKeep uint32
	CraaTypeMask []uint32
}


type CbRecallAny4res struct { // CB_RECALL_ANY4res
	CrarStatus int32
}


type CbRecallableObjAvail4res struct { // CB_RECALLABLE_OBJ_AVAIL4res
	CroaStatus int32
}


type CbRecallSlot4args struct { // CB_RECALL_SLOT4args
	RsaTargetHighestSlotid uint32
}


type CbRecallSlot4res struct { // CB_RECALL_SLOT4res
	RsrStatus int32
}


type ReferringCall4 struct { // referring_call4
	RcSequenceid uint32
	RcSlotid uint32
}


type ReferringCallList4 struct { // referring_call_list4
	RclSessionid Sessionid4
	RclReferringCalls []ReferringCall4
}


type CbSequence4args struct { // CB_SEQUENCE4args
	CsaSessionid Sessionid4
	CsaSequenceid uint32
	CsaSlotid uint32
	CsaHighestSlotid uint32
	CsaCachethis bool
	CsaReferringCallLists []ReferringCallList4
}


type CbSequence4resok struct { // CB_SEQUENCE4resok
	CsrSessionid Sessionid4
	CsrSequenceid uint32
	CsrSlotid uint32
	CsrHighestSlotid uint32
	CsrTargetHighestSlotid uint32
}


type CbWantsCancelled4args struct { // CB_WANTS_CANCELLED4args
	CwcaContendedWantsCancelled bool
	CwcaResourcedWantsCancelled bool
}


type CbWantsCancelled4res struct { // CB_WANTS_CANCELLED4res
	CwcrStatus int32
}


type CbNotifyLock4args struct { // CB_NOTIFY_LOCK4args
	CnlaFh NfsFh4
	CnlaLockOwner LockOwner4
}


type CbNotifyLock4res struct { // CB_NOTIFY_LOCK4res
	CnlrStatus int32
}


type NotifyDeviceidDelete4 struct { // notify_deviceid_delete4
	NddLayouttype int32
	NddDeviceid Deviceid4
}


type NotifyDeviceidChange4 struct { // notify_deviceid_change4
	NdcLayouttype int32
	NdcDeviceid Deviceid4
	NdcImmediate bool
}


type CbNotifyDeviceid4args struct { // CB_NOTIFY_DEVICEID4args
	CndaChanges []Notify4
}


type CbNotifyDeviceid4res struct { // CB_NOTIFY_DEVICEID4res
	CndrStatus int32
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


type Layoutreturn4 struct {
	LrReturntype int32 `xdr:"union"`
	LrLayout  LayoutreturnFile4  `xdr:"unioncase=1"` // File
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
	DevdataBlk  Specdata4  `xdr:"unioncase=3"` // Nf4blk
	DevdataChr  Specdata4  `xdr:"unioncase=4"` // Nf4chr
	//    `xdr:"unioncase=6"` // Nf4sock
	//    `xdr:"unioncase=7"` // Nf4fifo
	//    `xdr:"unioncase=2"` // Nf4dir
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
	OpenOwner  OpenToLockOwner4  `xdr:"unioncase=TRUE"` // True
	LockOwner  ExistLockOwner4  `xdr:"unioncase=FALSE"` // False
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
	CreateattrsUnchecked  Fattr4  `xdr:"unioncase=0"` // Unchecked4
	CreateattrsGuarded  Fattr4  `xdr:"unioncase=1"` // Guarded4
	Createverf  Verifier4  `xdr:"unioncase=2"` // Exclusive4
	ChCreateboth  Creatverfattr  `xdr:"unioncase=3"` // 1
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
	//    `xdr:"unioncase=4"` // Fh
	//    `xdr:"unioncase=6"` // Fh
	OcDelegateStateid  Stateid4  `xdr:"unioncase=5"` // Fh
}


type OpenNoneDelegation4 struct {
	OndWhy int32 `xdr:"union"`
	OndServerWillPushDeleg  bool  `xdr:"unioncase=1"` // Contention
	OndServerWillSignalAvail  bool  `xdr:"unioncase=2"` // Resource
}


type OpenDelegation4 struct {
	DelegationType int32 `xdr:"union"`
	//    `xdr:"unioncase=0"` // None
	Read  OpenReadDelegation4  `xdr:"unioncase=1"` // Read
	Write  OpenWriteDelegation4  `xdr:"unioncase=2"` // Write
	OdWhynone  OpenNoneDelegation4  `xdr:"unioncase=3"` // Ext
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
	FlavorInfo  RpcsecGssInfo  `xdr:"unioncase=RPCSEC_GSS"` // Gss
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


type CallbackSecParms4 struct {
	CbSecflavor uint32 `xdr:"union"`
	//    `xdr:"unioncase=AUTH_NONE"` // None
	CbspSysCred  AuthsysParms  `xdr:"unioncase=1"` // Sys
	CbspGssHandles  GssCbHandles4  `xdr:"unioncase=2"` // Gss
}


type BindConnToSession4res struct {
	BctsrStatus int32 `xdr:"union"`
	BctsrResok4  BindConnToSession4resok  `xdr:"unioncase=0"` // Ok
}


type StateProtect4A struct {
	SpaHow int32 `xdr:"union"`
	//    `xdr:"unioncase=0"` // None
	SpaMachOps  StateProtectOps4  `xdr:"unioncase=1"` // Cred
	SpaSsvParms  SsvSpParms4  `xdr:"unioncase=2"` // Ssv
}


type StateProtect4R struct {
	SprHow int32 `xdr:"union"`
	//    `xdr:"unioncase=0"` // None
	SprMachOps  StateProtectOps4  `xdr:"unioncase=1"` // Cred
	SprSsvInfo  SsvProtInfo4  `xdr:"unioncase=2"` // Ssv
}


type ExchangeID4res struct {
	EirStatus int32 `xdr:"union"`
	EirResok4  ExchangeID4resok  `xdr:"unioncase=0"` // Ok
}


type CreateSession4res struct {
	CsrStatus int32 `xdr:"union"`
	CsrResok4  CreateSession4resok  `xdr:"unioncase=0"` // Ok
}


type GetDirDelegation4resNonFatal struct {
	GddrnfStatus int32 `xdr:"union"`
	GddrnfResok4  GetDirDelegation4resok  `xdr:"unioncase=0"` // Ok
	GddrnfWillSignalDelegAvail  bool  `xdr:"unioncase=1"` // Unavail
}


type GetDirDelegation4res struct {
	GddrStatus int32 `xdr:"union"`
	GddrResNonFatal4  GetDirDelegation4resNonFatal  `xdr:"unioncase=0"` // Ok
}


type GETDEVICEINFO4res struct {
	GdirStatus int32 `xdr:"union"`
	GdirResok4  GETDEVICEINFO4resok  `xdr:"unioncase=0"` // Ok
	GdirMincount  uint32  `xdr:"unioncase=10005"` // Toosmall
}


type GETDEVICELIST4res struct {
	GdlrStatus int32 `xdr:"union"`
	GdlrResok4  GETDEVICELIST4resok  `xdr:"unioncase=0"` // Ok
}


type Newtime4 struct {
	NtTimechanged bool `xdr:"union"`
	NtTime  Nfstime4  `xdr:"unioncase=TRUE"` // True
	//    `xdr:"unioncase=FALSE"` // False
}


type Newoffset4 struct {
	NoNewoffset bool `xdr:"union"`
	NoOffset  uint64  `xdr:"unioncase=TRUE"` // True
	//    `xdr:"unioncase=FALSE"` // False
}


type Newsize4 struct {
	NsSizechanged bool `xdr:"union"`
	NsSize  uint64  `xdr:"unioncase=TRUE"` // True
	//    `xdr:"unioncase=FALSE"` // False
}


type LAYOUTCOMMIT4res struct {
	LocrStatus int32 `xdr:"union"`
	LocrResok4  LAYOUTCOMMIT4resok  `xdr:"unioncase=0"` // Ok
}


type LAYOUTGET4res struct {
	LogrStatus int32 `xdr:"union"`
	LogrResok4  LAYOUTGET4resok  `xdr:"unioncase=0"` // Ok
	LogrWillSignalLayoutAvail  bool  `xdr:"unioncase=10058"` // Layouttrylater
}


type LayoutreturnStateid struct {
	LrsPresent bool `xdr:"union"`
	LrsStateid  Stateid4  `xdr:"unioncase=TRUE"` // True
	//    `xdr:"unioncase=FALSE"` // False
}


type LAYOUTRETURN4res struct {
	LorrStatus int32 `xdr:"union"`
	LorrStateid  LayoutreturnStateid  `xdr:"unioncase=0"` // Ok
}


type SEQUENCE4res struct {
	SrStatus int32 `xdr:"union"`
	SrResok4  SEQUENCE4resok  `xdr:"unioncase=0"` // Ok
}


type SetSsv4res struct {
	SsrStatus int32 `xdr:"union"`
	SsrResok4  SetSsv4resok  `xdr:"unioncase=0"` // Ok
}


type TestStateid4res struct {
	TsrStatus int32 `xdr:"union"`
	TsrResok4  TestStateid4resok  `xdr:"unioncase=0"` // Ok
}


type DelegClaim4 struct {
	DcClaim int32 `xdr:"union"`
	//    `xdr:"unioncase=4"` // Fh
	//    `xdr:"unioncase=6"` // Fh
	DcDelegateType  int32  `xdr:"unioncase=1"` // Previous
}


type WantDelegation4res struct {
	WdrStatus int32 `xdr:"union"`
	WdrResok4  OpenDelegation4  `xdr:"unioncase=0"` // Ok
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
	OpbackchannelCtl  BackchannelCtl4args  `xdr:"unioncase=40"` // Ctl
	OpbindConnToSession  BindConnToSession4args  `xdr:"unioncase=41"` // Session
	OpexchangeID  ExchangeID4args  `xdr:"unioncase=42"` // Id
	OpcreateSession  CreateSession4args  `xdr:"unioncase=43"` // Session
	OpdestroySession  DestroySession4args  `xdr:"unioncase=44"` // Session
	OpfreeStateid  FreeStateid4args  `xdr:"unioncase=45"` // Stateid
	OpgetDirDelegation  GetDirDelegation4args  `xdr:"unioncase=46"` // Delegation
	Opgetdeviceinfo  GETDEVICEINFO4args  `xdr:"unioncase=47"` // Getdeviceinfo
	Opgetdevicelist  GETDEVICELIST4args  `xdr:"unioncase=48"` // Getdevicelist
	Oplayoutcommit  LAYOUTCOMMIT4args  `xdr:"unioncase=49"` // Layoutcommit
	Oplayoutget  LAYOUTGET4args  `xdr:"unioncase=50"` // Layoutget
	Oplayoutreturn  LAYOUTRETURN4args  `xdr:"unioncase=51"` // Layoutreturn
	OpsecinfoNoName  int32  `xdr:"unioncase=52"` // Name
	Opsequence  SEQUENCE4args  `xdr:"unioncase=53"` // Sequence
	OpsetSsv  SetSsv4args  `xdr:"unioncase=54"` // Ssv
	OptestStateid  TestStateid4args  `xdr:"unioncase=55"` // Stateid
	OpwantDelegation  WantDelegation4args  `xdr:"unioncase=56"` // Delegation
	OpdestroyClientid  Destroyuint64args  `xdr:"unioncase=57"` // Clientid
	OpreclaimComplete  ReclaimComplete4args  `xdr:"unioncase=58"` // Complete
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
	OpbackchannelCtl  BackchannelCtl4res  `xdr:"unioncase=40"` // Ctl
	OpbindConnToSession  BindConnToSession4res  `xdr:"unioncase=41"` // Session
	OpexchangeID  ExchangeID4res  `xdr:"unioncase=42"` // Id
	OpcreateSession  CreateSession4res  `xdr:"unioncase=43"` // Session
	OpdestroySession  DestroySession4res  `xdr:"unioncase=44"` // Session
	OpfreeStateid  FreeStateid4res  `xdr:"unioncase=45"` // Stateid
	OpgetDirDelegation  GetDirDelegation4res  `xdr:"unioncase=46"` // Delegation
	Opgetdeviceinfo  GETDEVICEINFO4res  `xdr:"unioncase=47"` // Getdeviceinfo
	Opgetdevicelist  GETDEVICELIST4res  `xdr:"unioncase=48"` // Getdevicelist
	Oplayoutcommit  LAYOUTCOMMIT4res  `xdr:"unioncase=49"` // Layoutcommit
	Oplayoutget  LAYOUTGET4res  `xdr:"unioncase=50"` // Layoutget
	Oplayoutreturn  LAYOUTRETURN4res  `xdr:"unioncase=51"` // Layoutreturn
	OpsecinfoNoName  SecinfoNoName4res  `xdr:"unioncase=52"` // Name
	Opsequence  SEQUENCE4res  `xdr:"unioncase=53"` // Sequence
	OpsetSsv  SetSsv4res  `xdr:"unioncase=54"` // Ssv
	OptestStateid  TestStateid4res  `xdr:"unioncase=55"` // Stateid
	OpwantDelegation  WantDelegation4res  `xdr:"unioncase=56"` // Delegation
	OpdestroyClientid  Destroyuint64res  `xdr:"unioncase=57"` // Clientid
	OpreclaimComplete  ReclaimComplete4res  `xdr:"unioncase=58"` // Complete
	Opillegal  ILLEGAL4res  `xdr:"unioncase=10044"` // Illegal
}


type CbGetattr4res struct {
	Status int32 `xdr:"union"`
	Resok4  CbGetattr4resok  `xdr:"unioncase=0"` // Ok
}


type Layoutrecall4 struct {
	LorRecalltype int32 `xdr:"union"`
	LorLayout  LayoutrecallFile4  `xdr:"unioncase=1"` // File
	LorFsid  Fsid4  `xdr:"unioncase=2"` // Fsid
	//    `xdr:"unioncase=3"` // All
}


type CbSequence4res struct {
	CsrStatus int32 `xdr:"union"`
	CsrResok4  CbSequence4resok  `xdr:"unioncase=0"` // Ok
}


type NfsCbArgop4 struct {
	Argop uint32 `xdr:"union"`
	Opcbgetattr  CbGetattr4args  `xdr:"unioncase=3"` // Getattr
	Opcbrecall  CbRecall4args  `xdr:"unioncase=4"` // Recall
	Opcblayoutrecall  CbLayoutrecall4args  `xdr:"unioncase=5"` // Layoutrecall
	Opcbnotify  CbNotify4args  `xdr:"unioncase=6"` // Notify
	OpcbpushDeleg  CbPushDeleg4args  `xdr:"unioncase=7"` // Deleg
	OpcbrecallAny  CbRecallAny4args  `xdr:"unioncase=8"` // Any
	OpcbrecallableObjAvail  CbRecallableObjAvail4args  `xdr:"unioncase=9"` // Avail
	OpcbrecallSlot  CbRecallSlot4args  `xdr:"unioncase=10"` // Slot
	Opcbsequence  CbSequence4args  `xdr:"unioncase=11"` // Sequence
	OpcbwantsCancelled  CbWantsCancelled4args  `xdr:"unioncase=12"` // Cancelled
	OpcbnotifyLock  CbNotifyLock4args  `xdr:"unioncase=13"` // Lock
	OpcbnotifyDeviceid  CbNotifyDeviceid4args  `xdr:"unioncase=14"` // Deviceid
	//    `xdr:"unioncase=10044"` // Illegal
}


type NfsCbResop4 struct {
	Resop uint32 `xdr:"union"`
	Opcbgetattr  CbGetattr4res  `xdr:"unioncase=3"` // Getattr
	Opcbrecall  CbRecall4res  `xdr:"unioncase=4"` // Recall
	Opcblayoutrecall  CbLayoutrecall4res  `xdr:"unioncase=5"` // Layoutrecall
	Opcbnotify  CbNotify4res  `xdr:"unioncase=6"` // Notify
	OpcbpushDeleg  CbPushDeleg4res  `xdr:"unioncase=7"` // Deleg
	OpcbrecallAny  CbRecallAny4res  `xdr:"unioncase=8"` // Any
	OpcbrecallableObjAvail  CbRecallableObjAvail4res  `xdr:"unioncase=9"` // Avail
	OpcbrecallSlot  CbRecallSlot4res  `xdr:"unioncase=10"` // Slot
	Opcbsequence  CbSequence4res  `xdr:"unioncase=11"` // Sequence
	OpcbwantsCancelled  CbWantsCancelled4res  `xdr:"unioncase=12"` // Cancelled
	OpcbnotifyLock  CbNotifyLock4res  `xdr:"unioncase=13"` // Lock
	OpcbnotifyDeviceid  CbNotifyDeviceid4res  `xdr:"unioncase=14"` // Deviceid
	Opcbillegal  CbIllegal4res  `xdr:"unioncase=10044"` // Illegal
}





func Access (access uint32) (NfsArgop4) {
	return NfsArgop4{Argop:3, Opaccess:ACCESS4args{ Access:access } }
}


func Close (seqid uint32, openstateid Stateid4) (NfsArgop4) {
	return NfsArgop4{Argop:4, Opclose:CLOSE4args{ Seqid:seqid, OpenStateid:openstateid } }
}


func Commit (offset uint64, count uint32) (NfsArgop4) {
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


func Lock (locktype int32, reclaim bool, offset uint64, length uint64, locker Locker4) (NfsArgop4) {
	return NfsArgop4{Argop:12, Oplock:LOCK4args{ Locktype:locktype, Reclaim:reclaim, Offset:offset, Length:length, Locker:locker } }
}


func Lockt (locktype int32, offset uint64, length uint64, owner LockOwner4) (NfsArgop4) {
	return NfsArgop4{Argop:13, Oplockt:LOCKT4args{ Locktype:locktype, Offset:offset, Length:length, Owner:owner } }
}


func Locku (locktype int32, seqid uint32, lockstateid Stateid4, offset uint64, length uint64) (NfsArgop4) {
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


func OpenConfirm (openstateid Stateid4, seqid uint32) (NfsArgop4) {
	return NfsArgop4{Argop:20, OpopenConfirm:OpenConfirm4args{ OpenStateid:openstateid, Seqid:seqid } }
}


func Downgrade (openstateid Stateid4, seqid uint32, shareaccess uint32, sharedeny uint32) (NfsArgop4) {
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


func Read (stateid Stateid4, offset uint64, count uint32) (NfsArgop4) {
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


func Write (stateid Stateid4, offset uint64, stable int32, data []byte) (NfsArgop4) {
	return NfsArgop4{Argop:38, Opwrite:WRITE4args{ Stateid:stateid, Offset:offset, Stable:stable, Data:data } }
}


func Lockowner (lockowner LockOwner4) (NfsArgop4) {
	return NfsArgop4{Argop:39, OpreleaseLockowner:ReleaseLockowner4args{ LockOwner:lockowner } }
}


func Ctl (bcacbprogram uint32, bcasecparms []CallbackSecParms4) (NfsArgop4) {
	return NfsArgop4{Argop:40, OpbackchannelCtl:BackchannelCtl4args{ BcaCbProgram:bcacbprogram, BcaSecParms:bcasecparms } }
}


func BindConnToSession (bctsasessid Sessionid4, bctsadir int32, bctsauseconninrdmamode bool) (NfsArgop4) {
	return NfsArgop4{Argop:41, OpbindConnToSession:BindConnToSession4args{ BctsaSessid:bctsasessid, BctsaDir:bctsadir, BctsaUseConnInRdmaMode:bctsauseconninrdmamode } }
}


func ExchangeId (eiaclientowner ClientOwner4, eiaflags uint32, eiastateprotect StateProtect4A, eiaclientimplid []NfsImplID4) (NfsArgop4) {
	return NfsArgop4{Argop:42, OpexchangeID:ExchangeID4args{ EiaClientowner:eiaclientowner, EiaFlags:eiaflags, EiaStateProtect:eiastateprotect, EiaClientImplID:eiaclientimplid } }
}


func CreateSession (csaclientid uint64, csasequence uint32, csaflags uint32, csaforechanattrs ChannelAttrs4, csabackchanattrs ChannelAttrs4, csacbprogram uint32, csasecparms []CallbackSecParms4) (NfsArgop4) {
	return NfsArgop4{Argop:43, OpcreateSession:CreateSession4args{ CsaClientid:csaclientid, CsaSequence:csasequence, CsaFlags:csaflags, CsaForeChanAttrs:csaforechanattrs, CsaBackChanAttrs:csabackchanattrs, CsaCbProgram:csacbprogram, CsaSecParms:csasecparms } }
}


func DestroySession (dsasessionid Sessionid4) (NfsArgop4) {
	return NfsArgop4{Argop:44, OpdestroySession:DestroySession4args{ DsaSessionid:dsasessionid } }
}


func FreeStateid (fsastateid Stateid4) (NfsArgop4) {
	return NfsArgop4{Argop:45, OpfreeStateid:FreeStateid4args{ FsaStateid:fsastateid } }
}


func GetDirDelegation (gddasignaldelegavail bool, gddanotificationtypes []uint32, gddachildattrdelay AttrNotice4, gddadirattrdelay AttrNotice4, gddachildattributes []uint32, gddadirattributes []uint32) (NfsArgop4) {
	return NfsArgop4{Argop:46, OpgetDirDelegation:GetDirDelegation4args{ GddaSignalDelegAvail:gddasignaldelegavail, GddaNotificationTypes:gddanotificationtypes, GddaChildAttrDelay:gddachildattrdelay, GddaDirAttrDelay:gddadirattrdelay, GddaChildAttributes:gddachildattributes, GddaDirAttributes:gddadirattributes } }
}


func Getdeviceinfo (gdiadeviceid Deviceid4, gdialayouttype int32, gdiamaxcount uint32, gdianotifytypes []uint32) (NfsArgop4) {
	return NfsArgop4{Argop:47, Opgetdeviceinfo:GETDEVICEINFO4args{ GdiaDeviceID:gdiadeviceid, GdiaLayoutType:gdialayouttype, GdiaMaxcount:gdiamaxcount, GdiaNotifyTypes:gdianotifytypes } }
}


func Getdevicelist (gdlalayouttype int32, gdlamaxdevices uint32, gdlacookie uint64, gdlacookieverf Verifier4) (NfsArgop4) {
	return NfsArgop4{Argop:48, Opgetdevicelist:GETDEVICELIST4args{ GdlaLayoutType:gdlalayouttype, GdlaMaxdevices:gdlamaxdevices, GdlaCookie:gdlacookie, GdlaCookieverf:gdlacookieverf } }
}


func Layoutcommit (locaoffset uint64, localength uint64, locareclaim bool, locastateid Stateid4, localastwriteoffset Newoffset4, locatimemodify Newtime4, localayoutupdate Layoutupdate4) (NfsArgop4) {
	return NfsArgop4{Argop:49, Oplayoutcommit:LAYOUTCOMMIT4args{ LocaOffset:locaoffset, LocaLength:localength, LocaReclaim:locareclaim, LocaStateid:locastateid, LocaLastWriteOffset:localastwriteoffset, LocaTimeModify:locatimemodify, LocaLayoutupdate:localayoutupdate } }
}


func Layoutget (logasignallayoutavail bool, logalayouttype int32, logaiomode int32, logaoffset uint64, logalength uint64, logaminlength uint64, logastateid Stateid4, logamaxcount uint32) (NfsArgop4) {
	return NfsArgop4{Argop:50, Oplayoutget:LAYOUTGET4args{ LogaSignalLayoutAvail:logasignallayoutavail, LogaLayoutType:logalayouttype, LogaIomode:logaiomode, LogaOffset:logaoffset, LogaLength:logalength, LogaMinlength:logaminlength, LogaStateid:logastateid, LogaMaxcount:logamaxcount } }
}


func Layoutreturn (lorareclaim bool, loralayouttype int32, loraiomode int32, loralayoutreturn Layoutreturn4) (NfsArgop4) {
	return NfsArgop4{Argop:51, Oplayoutreturn:LAYOUTRETURN4args{ LoraReclaim:lorareclaim, LoraLayoutType:loralayouttype, LoraIomode:loraiomode, LoraLayoutreturn:loralayoutreturn } }
}


func SecinfoNoName (arg int32) (NfsArgop4) {
	return NfsArgop4{Argop:52, OpsecinfoNoName:arg }
}


func Sequence (sasessionid Sessionid4, sasequenceid uint32, saslotid uint32, sahighestslotid uint32, sacachethis bool) (NfsArgop4) {
	return NfsArgop4{Argop:53, Opsequence:SEQUENCE4args{ SaSessionid:sasessionid, SaSequenceid:sasequenceid, SaSlotid:saslotid, SaHighestSlotid:sahighestslotid, SaCachethis:sacachethis } }
}


func Ssv (ssassv []byte, ssadigest []byte) (NfsArgop4) {
	return NfsArgop4{Argop:54, OpsetSsv:SetSsv4args{ SsaSsv:ssassv, SsaDigest:ssadigest } }
}


func TestStateid (tsstateids []Stateid4) (NfsArgop4) {
	return NfsArgop4{Argop:55, OptestStateid:TestStateid4args{ TsStateids:tsstateids } }
}


func WantDelegation (wdawant uint32, wdaclaim DelegClaim4) (NfsArgop4) {
	return NfsArgop4{Argop:56, OpwantDelegation:WantDelegation4args{ WdaWant:wdawant, WdaClaim:wdaclaim } }
}


func DestroyClientid (dcaclientid uint64) (NfsArgop4) {
	return NfsArgop4{Argop:57, OpdestroyClientid:Destroyuint64args{ DcaClientid:dcaclientid } }
}


func ReclaimComplete (rcaonefs bool) (NfsArgop4) {
	return NfsArgop4{Argop:58, OpreclaimComplete:ReclaimComplete4args{ RcaOneFs:rcaonefs } }
}


func Illegal () (NfsArgop4) {
	return NfsArgop4{Argop:10044 }
}


