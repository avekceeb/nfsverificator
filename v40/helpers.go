package v40

func FhFromString(h string) (NfsFh4) {
    return NfsFh4([]byte(h))
    //copy(fh[:], s[:NFS4_FHSIZE])
    //return fh
}

func OpNameNfs40(code int32) string {
	switch code {
	case OP_ACCESS: return "ACCESS"
	case OP_CLOSE: return "CLOSE"
	case OP_COMMIT: return "COMMIT"
	case OP_CREATE: return "CREATE"
	case OP_DELEGPURGE: return "DELEGPURGE"
	case OP_DELEGRETURN: return "DELEGRETURN"
	case OP_GETATTR: return "GETATTR"
	case OP_GETFH: return "GETFH"
	case OP_LINK: return "LINK"
	case OP_LOCK: return "LOCK"
	case OP_LOCKT: return "LOCKT"
	case OP_LOCKU: return "LOCKU"
	case OP_LOOKUP: return "LOOKUP"
	case OP_LOOKUPP: return "LOOKUPP"
	case OP_NVERIFY: return "NVERIFY"
	case OP_OPEN: return "OPEN"
	case OP_OPENATTR: return "OPENATTR"
	case OP_OPEN_CONFIRM: return "OPEN_CONFIRM"
	case OP_OPEN_DOWNGRADE: return "OPEN_DOWNGRADE"
	case OP_PUTFH: return "PUTFH"
	case OP_PUTPUBFH: return "PUTPUBFH"
	case OP_PUTROOTFH: return "PUTROOTFH"
	case OP_READ: return "READ"
	case OP_READDIR: return "READDIR"
	case OP_READLINK: return "READLINK"
	case OP_REMOVE: return "REMOVE"
	case OP_RENAME: return "RENAME"
	case OP_RENEW: return "RENEW"
	case OP_RESTOREFH: return "RESTOREFH"
	case OP_SAVEFH: return "SAVEFH"
	case OP_SECINFO: return "SECINFO"
	case OP_SETATTR: return "SETATTR"
	case OP_SETCLIENTID: return "SETCLIENTID"
	case OP_SETCLIENTID_CONFIRM: return "SETCLIENTID_CONFIRM"
	case OP_VERIFY: return "VERIFY"
	case OP_WRITE: return "WRITE"
	case OP_RELEASE_LOCKOWNER: return "RELEASE_LOCKOWNER"
	case OP_ILLEGAL: return "ILLEGAL"
	//case OP_CB_GETATTR: return "CB_GETATTR"
	//case OP_CB_RECALL: return "CB_RECALL"
	//case OP_CB_ILLEGAL: return "CB_ILLEGAL"
	}
	return "UNKNOWN"
}

func ErrorNameNfs40(code int32) string {
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
	case NFS4ERR_MINOR_VERS_MISMATCH:
		return "NFS4ERR_MINOR_VERS_MISMATCH"
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
	default:
		return "UNKNOWN"
	}
}

func LastRes(res *([]NfsResop4)) (*NfsResop4) {
	return &((*res)[len(*res)-1])
}

func GrabFh(res *([]NfsResop4)) (NfsFh4) {
	return LastRes(res).Opgetfh.Resok4.Object
}

func AreFhEqual(a, b NfsFh4) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
