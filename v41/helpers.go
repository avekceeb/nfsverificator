package v41

func ErrorName(code int32) string {
	switch code {
	case NFS4_OK: return "NFS4_OK"
	case NFS4ERR_PERM: return "NFS4ERR_PERM"
	case NFS4ERR_NOENT: return "NFS4ERR_NOENT"
	case NFS4ERR_IO: return "NFS4ERR_IO"
	case NFS4ERR_NXIO: return "NFS4ERR_NXIO"
	case NFS4ERR_ACCESS: return "NFS4ERR_ACCESS"
	case NFS4ERR_EXIST: return "NFS4ERR_EXIST"
	case NFS4ERR_XDEV: return "NFS4ERR_XDEV"
	case NFS4ERR_NOTDIR: return "NFS4ERR_NOTDIR"
	case NFS4ERR_ISDIR: return "NFS4ERR_ISDIR"
	case NFS4ERR_INVAL: return "NFS4ERR_INVAL"
	case NFS4ERR_FBIG: return "NFS4ERR_FBIG"
	case NFS4ERR_NOSPC: return "NFS4ERR_NOSPC"
	case NFS4ERR_ROFS: return "NFS4ERR_ROFS"
	case NFS4ERR_MLINK: return "NFS4ERR_MLINK"
	case NFS4ERR_NAMETOOLONG: return "NFS4ERR_NAMETOOLONG"
	case NFS4ERR_NOTEMPTY: return "NFS4ERR_NOTEMPTY"
	case NFS4ERR_DQUOT: return "NFS4ERR_DQUOT"
	case NFS4ERR_STALE: return "NFS4ERR_STALE"
	case NFS4ERR_BADHANDLE: return "NFS4ERR_BADHANDLE"
	case NFS4ERR_BAD_COOKIE: return "NFS4ERR_BAD_COOKIE"
	case NFS4ERR_NOTSUPP: return "NFS4ERR_NOTSUPP"
	case NFS4ERR_TOOSMALL: return "NFS4ERR_TOOSMALL"
	case NFS4ERR_SERVERFAULT: return "NFS4ERR_SERVERFAULT"
	case NFS4ERR_BADTYPE: return "NFS4ERR_BADTYPE"
	case NFS4ERR_DELAY: return "NFS4ERR_DELAY"
	case NFS4ERR_SAME: return "NFS4ERR_SAME"
	case NFS4ERR_DENIED: return "NFS4ERR_DENIED"
	case NFS4ERR_EXPIRED: return "NFS4ERR_EXPIRED"
	case NFS4ERR_LOCKED: return "NFS4ERR_LOCKED"
	case NFS4ERR_GRACE: return "NFS4ERR_GRACE"
	case NFS4ERR_FHEXPIRED: return "NFS4ERR_FHEXPIRED"
	case NFS4ERR_SHARE_DENIED: return "NFS4ERR_SHARE_DENIED"
	case NFS4ERR_WRONGSEC: return "NFS4ERR_WRONGSEC"
	case NFS4ERR_CLID_INUSE: return "NFS4ERR_CLID_INUSE"
	case NFS4ERR_RESOURCE: return "NFS4ERR_RESOURCE"
	case NFS4ERR_MOVED: return "NFS4ERR_MOVED"
	case NFS4ERR_NOFILEHANDLE: return "NFS4ERR_NOFILEHANDLE"
	case NFS4ERR_MINOR_VERS_MISMATCH: return "NFS4ERR_MINOR_VERS_MISMATCH"
	case NFS4ERR_STALE_CLIENTID: return "NFS4ERR_STALE_CLIENTID"
	case NFS4ERR_STALE_STATEID: return "NFS4ERR_STALE_STATEID"
	case NFS4ERR_OLD_STATEID: return "NFS4ERR_OLD_STATEID"
	case NFS4ERR_BAD_STATEID: return "NFS4ERR_BAD_STATEID"
	case NFS4ERR_BAD_SEQID: return "NFS4ERR_BAD_SEQID"
	case NFS4ERR_NOT_SAME: return "NFS4ERR_NOT_SAME"
	case NFS4ERR_LOCK_RANGE: return "NFS4ERR_LOCK_RANGE"
	case NFS4ERR_SYMLINK: return "NFS4ERR_SYMLINK"
	case NFS4ERR_RESTOREFH: return "NFS4ERR_RESTOREFH"
	case NFS4ERR_LEASE_MOVED: return "NFS4ERR_LEASE_MOVED"
	case NFS4ERR_ATTRNOTSUPP: return "NFS4ERR_ATTRNOTSUPP"
	case NFS4ERR_NO_GRACE: return "NFS4ERR_NO_GRACE"
	case NFS4ERR_RECLAIM_BAD: return "NFS4ERR_RECLAIM_BAD"
	case NFS4ERR_RECLAIM_CONFLICT: return "NFS4ERR_RECLAIM_CONFLICT"
	case NFS4ERR_BADXDR: return "NFS4ERR_BADXDR"
	case NFS4ERR_LOCKS_HELD: return "NFS4ERR_LOCKS_HELD"
	case NFS4ERR_OPENMODE: return "NFS4ERR_OPENMODE"
	case NFS4ERR_BADOWNER: return "NFS4ERR_BADOWNER"
	case NFS4ERR_BADCHAR: return "NFS4ERR_BADCHAR"
	case NFS4ERR_BADNAME: return "NFS4ERR_BADNAME"
	case NFS4ERR_BAD_RANGE: return "NFS4ERR_BAD_RANGE"
	case NFS4ERR_LOCK_NOTSUPP: return "NFS4ERR_LOCK_NOTSUPP"
	case NFS4ERR_OP_ILLEGAL: return "NFS4ERR_OP_ILLEGAL"
	case NFS4ERR_DEADLOCK: return "NFS4ERR_DEADLOCK"
	case NFS4ERR_FILE_OPEN: return "NFS4ERR_FILE_OPEN"
	case NFS4ERR_ADMIN_REVOKED: return "NFS4ERR_ADMIN_REVOKED"
	case NFS4ERR_CB_PATH_DOWN: return "NFS4ERR_CB_PATH_DOWN"
	case NFS4ERR_BADIOMODE: return "NFS4ERR_BADIOMODE"
	case NFS4ERR_BADLAYOUT: return "NFS4ERR_BADLAYOUT"
	case NFS4ERR_BAD_SESSION_DIGEST: return "NFS4ERR_BAD_SESSION_DIGEST"
	case NFS4ERR_BADSESSION: return "NFS4ERR_BADSESSION"
	case NFS4ERR_BADSLOT: return "NFS4ERR_BADSLOT"
	case NFS4ERR_COMPLETE_ALREADY: return "NFS4ERR_COMPLETE_ALREADY"
	case NFS4ERR_CONN_NOT_BOUND_TO_SESSION:
		return "NFS4ERR_CONN_NOT_BOUND_TO_SESSION"
	case NFS4ERR_DELEG_ALREADY_WANTED: return "NFS4ERR_DELEG_ALREADY_WANTED"
	case NFS4ERR_BACK_CHAN_BUSY: return "NFS4ERR_BACK_CHAN_BUSY"
	case NFS4ERR_LAYOUTTRYLATER: return "NFS4ERR_LAYOUTTRYLATER"
	case NFS4ERR_LAYOUTUNAVAILABLE: return "NFS4ERR_LAYOUTUNAVAILABLE"
	case NFS4ERR_NOMATCHING_LAYOUT: return "NFS4ERR_NOMATCHING_LAYOUT"
	case NFS4ERR_RECALLCONFLICT: return "NFS4ERR_RECALLCONFLICT"
	case NFS4ERR_UNKNOWN_LAYOUTTYPE: return "NFS4ERR_UNKNOWN_LAYOUTTYPE"
	case NFS4ERR_SEQ_MISORDERED: return "NFS4ERR_SEQ_MISORDERED"
	case NFS4ERR_SEQUENCE_POS: return "NFS4ERR_SEQUENCE_POS"
	case NFS4ERR_REQ_TOO_BIG: return "NFS4ERR_REQ_TOO_BIG"
	case NFS4ERR_REP_TOO_BIG: return "NFS4ERR_REP_TOO_BIG"
	case NFS4ERR_REP_TOO_BIG_TO_CACHE: return "NFS4ERR_REP_TOO_BIG_TO_CACHE"
	case NFS4ERR_RETRY_UNCACHED_REP: return "NFS4ERR_RETRY_UNCACHED_REP"
	case NFS4ERR_UNSAFE_COMPOUND: return "NFS4ERR_UNSAFE_COMPOUND"
	case NFS4ERR_TOO_MANY_OPS: return "NFS4ERR_TOO_MANY_OPS"
	case NFS4ERR_OP_NOT_IN_SESSION: return "NFS4ERR_OP_NOT_IN_SESSION"
	case NFS4ERR_HASH_ALG_UNSUPP: return "NFS4ERR_HASH_ALG_UNSUPP"
	case NFS4ERR_CLIENTID_BUSY: return "NFS4ERR_CLIENTID_BUSY"
	case NFS4ERR_PNFS_IO_HOLE: return "NFS4ERR_PNFS_IO_HOLE"
	case NFS4ERR_SEQ_FALSE_RETRY: return "NFS4ERR_SEQ_FALSE_RETRY"
	case NFS4ERR_BAD_HIGH_SLOT: return "NFS4ERR_BAD_HIGH_SLOT"
	case NFS4ERR_DEADSESSION: return "NFS4ERR_DEADSESSION"
	case NFS4ERR_ENCR_ALG_UNSUPP: return "NFS4ERR_ENCR_ALG_UNSUPP"
	case NFS4ERR_PNFS_NO_LAYOUT: return "NFS4ERR_PNFS_NO_LAYOUT"
	case NFS4ERR_NOT_ONLY_OP: return "NFS4ERR_NOT_ONLY_OP"
	case NFS4ERR_WRONG_CRED: return "NFS4ERR_WRONG_CRED"
	case NFS4ERR_WRONG_TYPE: return "NFS4ERR_WRONG_TYPE"
	case NFS4ERR_DIRDELEG_UNAVAIL: return "NFS4ERR_DIRDELEG_UNAVAIL"
	case NFS4ERR_REJECT_DELEG: return "NFS4ERR_REJECT_DELEG"
	case NFS4ERR_RETURNCONFLICT: return "NFS4ERR_RETURNCONFLICT"
	case NFS4ERR_DELEG_REVOKED: return "NFS4ERR_DELEG_REVOKED"
	default:
		return "UNKNOWN"
	}
}