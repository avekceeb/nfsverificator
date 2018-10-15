package nfs40

import (
    "github.com/avekceeb/nfsverificator/rpc"
)

const (
	NFS4_PROGRAM = 100003
	NFS_V4 = 4
	NFSPROC4_NULL = 0
	NFSPROC4_COMPOUND = 1
	NFS4_CALLBACK = 0x40000000
	NFS_CB = 1
	NFS4_FHSIZE = 128
	NFS4_VERIFIER_SIZE = 8
	NFS4_OTHER_SIZE = 12
	NFS4_OPAQUE_LIMIT = 1024
	NFS4_INT64_MAX = 0x7fffffffffffffff
	NFS4_UINT64_MAX = 0xffffffffffffffff
	NFS4_INT32_MAX = 0x7fffffff
	NFS4_UINT32_MAX = 0xffffffff
)

const (
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
	OP_ILLEGAL = 10044
)

type PUTFH4args struct {
	FH string
}

// stateid4
type StateId4 struct {
	SeqId uint32
	Other [NFS4_OTHER_SIZE]byte
}

// opentype4
const (
	OPEN4_NOCREATE = 0
	OPEN4_CREATE = 1
)

// TODO
//OpenType4 := map[int32]bool{OPEN4_NOCREATE:true, OPEN4_CREATE:true}

// createmode4
const (
	UNCHECKED4 = 0
	GUARDED4 = 1
	EXCLUSIVE4 = 2
)

// fattr4
type FAttr struct {
	Bitmap []uint32
	AttrList string
}

// openflag4
type OpenFlag4 struct {
	OpenType int32   `xdr:"union"`
	CreateHow struct {
		CreateMode int32 `xdr:"union"`
		Attr FAttr `xdr:"unioncalse=0"` // both are the same
		AttrGuarded FAttr `xdr:"unioncalse=1"`
		Verifier [NFS4_VERIFIER_SIZE]byte `xdr:"unioncalse=2"`
	} `xdr:""unioncase=1`
}
// open_delegation_type4
const (
	OPEN_DELEGATE_NONE = 0
	OPEN_DELEGATE_READ = 1
	OPEN_DELEGATE_WRITE = 2
)

// open_claim_type4
const (
	CLAIM_NULL = 0
	CLAIM_PREVIOUS = 1
	CLAIM_DELEGATE_CUR = 2
	CLAIM_DELEGATE_PREV = 3
)

// open_claim_delegate_cur4
type OpenClaimDelegateCur4 struct  {
	DelegateStateId StateId4
	File string
}

// open_claim4
type OpenClaim4 struct  {
	Claim            int32                 `xdr:"union"`
	File             string                `xdr:"unioncase=0"`
	DelegationType   int32                 `xdr:"unioncase=1"`
	DelegateCurInfo  OpenClaimDelegateCur4 `xdr:"unioncase=2"`
	FileDelegatePrev string                `xdr:"unioncase=3"`
}

type OpenOwner4 struct {
	ClientId NfsClientId
	Owner string
}

type OPEN4args struct {
	SeqId uint32
	ShareAccess uint32
	ShareDeny uint32
	OpenOwner OpenOwner4
	OpenHow  OpenFlag4
	Claim OpenClaim4
}

type NfsClientId struct {
	Verifier [NFS4_VERIFIER_SIZE]byte
	Id string // ?
}

type ClientAddr struct {
	NetId string
	Addr string
}

type CallbackClient struct {
	Program uint
	Location ClientAddr
};

type SETCLIENTID4args struct {
	Client NfsClientId
	Callback CallbackClient
	CallbackIdent uint32
}

type SETCLIENTID_CONFIRM4args struct {
	ClientId uint64
	Verifier [NFS4_VERIFIER_SIZE]byte
}

type READDIR4args struct {
	Cookie uint64
	Verifier [NFS4_VERIFIER_SIZE]byte
	Dircount uint32
	Count uint32
	Bitmap []uint32
}

// nfs_argop4
type NfsArgOp4 struct {
	ArgOp              uint32                   `xdr:"union"`
	AttrRequest        []uint32                 `xdr:"unioncase=9"`
	PutFH              PUTFH4args               `xdr:"unioncase=22"`
	ReadDir            READDIR4args             `xdr:"unioncase=26"`
	SetClientId        SETCLIENTID4args         `xdr:"unioncase=35"`
	SetClientIdConfirm SETCLIENTID_CONFIRM4args `xdr:"unioncase=36"`
}

type ArgArrayT struct {
	Args []NfsArgOp4
}

type COMPOUND4args struct {
	Tag string
	MinorVersion uint32
	ArgArray ArgArrayT
}

type CompoundMessage struct {
	Head rpc.Header
	Args COMPOUND4args
}

/////////// server replies /////////////////////////
// enum nfsstat4
const (
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
)

type SETCLIENTID4res struct {

}

type SETCLIENTID_CONFIRM4res struct {
}

// nfs_resop4
type NfsResOp4 struct {
	ResOp              uint32                  `xdr:"union"`
	SetClientId	       SETCLIENTID4res         `xdr:"unioncase=35"`
	SetClientIdConfirm SETCLIENTID_CONFIRM4res `xdr:"unioncase=36"`
}

type COMPOUND4res struct {
	Status    int32
	Tag       string
	ResArray  []NfsResOp4
}