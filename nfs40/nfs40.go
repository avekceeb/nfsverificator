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
	OPEN4_SHARE_ACCESS_READ = 0x00000001
	OPEN4_SHARE_ACCESS_WRITE = 0x00000002
	OPEN4_SHARE_ACCESS_BOTH = 0x00000003
	OPEN4_SHARE_DENY_NONE = 0x00000000
	OPEN4_SHARE_DENY_READ = 0x00000001
	OPEN4_SHARE_DENY_WRITE = 0x00000002
	OPEN4_SHARE_DENY_BOTH = 0x00000003
)

const (
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
)

// nfs_ftype4
const (
	NF4REG = 1
	NF4DIR = 2
	NF4BLK = 3
	NF4CHR = 4
	NF4LNK = 5
	NF4SOCK = 6
	NF4FIFO = 7
	NF4ATTRDIR = 8
	NF4NAMEDATTR = 9
)

// stable_how4
const (
	UNSTABLE4 = 0
	DATA_SYNC4 = 1
	FILE_SYNC4 = 2
)

const (
	FH4_PERSISTENT = 0x00000000
	FH4_NOEXPIRE_WITH_OPEN = 0x00000001
	FH4_VOLATILE_ANY = 0x00000002
	FH4_VOL_MIGRATION = 0x00000004
	FH4_VOL_RENAME = 0x00000008
)

// nfs_lock_type4
const (
	READ_LT = 1
	WRITE_LT = 2
	READW_LT = 3
	WRITEW_LT = 4
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

//////////////////////////////////////////////

type Verifier4 [NFS4_VERIFIER_SIZE]byte
// nfs_fh4
type FH4 []byte

type PUTFH4args struct {
	FH FH4
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

// createmode4
const (
	UNCHECKED4 = 0
	GUARDED4 = 1
	EXCLUSIVE4 = 2
)

// fattr4
type FAttr4 struct {
	Bitmap []uint32 // bitmap4
	AttrList []byte // attrlist4
}
/*
   The bitmap is a counted array of 32 bit integers used to contain bit
   values.  The position of the integer in the array that contains bit n
   can be computed from the expression (n / 32) and its bit within that
   integer is (n mod 32).

                           0            1
         +-----------+-----------+-----------+--
         |  count    | 31  ..  0 | 63  .. 32 |
         +-----------+-----------+-----------+--

*/

type CreateHowT struct {
		CreateMode  int32     `xdr:"union"`
		Attr        FAttr4    `xdr:"unioncase=0"` // both are the same
		AttrGuarded FAttr4    `xdr:"unioncase=1"`
		Verifier    Verifier4 `xdr:"unioncase=2"`
}

// openflag4
type OpenFlag4 struct {
	OpenType  int32      `xdr:"union"`
	CreateHow CreateHowT `xdr:"unioncase=1"`
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

// open_owner4
type OpenOwner4 struct {
	ClientId uint64 // clientid4
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

type OPEN_CONFIRM4args struct {
	State StateId4
	SeqId uint32 // ???????????? in *.h
}

type NfsClientId struct {
	Verifier Verifier4
	Id string // ?
}

// clientaddr4
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
	Verifier Verifier4
}

type READDIR4args struct {
	Cookie uint64
	Verifier Verifier4
	Dircount uint32
	Count uint32
	Bitmap []uint32
}

type SETATTR4args struct {
	StateId StateId4  //stateid4 stateid;
	Attr FAttr4       //fattr4 obj_attributes;
}

type LOOKUP4args struct {
	Name string //component4 objname;
}

// createtype4
type CreateType4 struct {
	Type     int32	   `xdr:"union"`
	// TODO: check RFC
	SpecBlk  [2]uint32 `xdr:"unioncase=3"`
	SpecChr  [2]uint32 `xdr:"unioncase=4"`
	Link     string    `xdr:"unioncase=5"`
}

type CREATE4args struct {
	CreateType CreateType4
	Name       string
	Attr       FAttr4
}

type WRITE4args struct {
	State  StateId4
	Offset uint64 //offset4
	Stable int32  //stable_how4
	Data   []byte
}

type CLOSE4args struct {
	SeqId   uint32
	StateId StateId4
}

//////////// LOCKS ///////////////////////////

// lock_owner4
type LockOwner4 struct {
	ClientId uint64 // clientid4
	Owner string
}

//  open_to_lock_owner4
type OpenToLockOwner4 struct {
	SeqId uint32 // seqid4
	StateId StateId4 // stateid4
	LockSeqId uint32 // seqid4
	LockOwner LockOwner4
}

// exist_lock_owner4
type ExistLockOwner4 struct {
	StateId StateId4
	SeqId uint32 // seqid4
}

// locker4
type Locker4 struct {
	New       bool             `xdr:"union"`           // new_lock_owner;
	LockOwner ExistLockOwner4  `xdr:"unioncase=0"` // exist_lock_owner4
	OpenOwner OpenToLockOwner4 `xdr:"unioncase=1"`  // open_to_lock_owner4
}

type LOCK4args struct {
	LockType int32 // nfs_lock_type4
	Reclaim bool
	Offset uint64
	Length uint64
	Locker Locker4
}

type LOCKT4args struct {
	LockType int32 // nfs_lock_type4
	Offset uint64
	Length uint64
	Owner LockOwner4
}

type RELEASE_LOCKOWNER4args struct {
	LockOwner LockOwner4
}

// nfs_argop4
type NfsArgOp4 struct {
	ArgOp              uint32                   `xdr:"union"`
	Close              CLOSE4args               `xdr:"unioncase=4"`
	Create             CREATE4args              `xdr:"unioncase=6"`
	AttrRequest        []uint32                 `xdr:"unioncase=9"`
	Lock               LOCK4args                `xdr:"unioncase=12"`
	LockT              LOCKT4args               `xdr:"unioncase=13"`
	Lookup             LOOKUP4args              `xdr:"unioncase=15"`
	Open               OPEN4args                `xdr:"unioncase=18"`
	OpenConfirm        OPEN_CONFIRM4args        `xdr:"unioncase=20"`
	PutFH              PUTFH4args               `xdr:"unioncase=22"`
	ReadDir            READDIR4args             `xdr:"unioncase=26"`
	SetAttr            SETATTR4args             `xdr:"unioncase=34"`
	SetClientId        SETCLIENTID4args         `xdr:"unioncase=35"`
	SetClientIdConfirm SETCLIENTID_CONFIRM4args `xdr:"unioncase=36"`
	Write              WRITE4args               `xdr:"unioncase=38"`
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

type SETCLIENTID4resok struct {
	ClientId uint64 //clientid4
	Verifier [NFS4_VERIFIER_SIZE]byte //verifier4
}

type SETCLIENTID4res struct {
	Status int32
	//union {
	ResOk SETCLIENTID4resok
	//	ClientAddr clientaddr4 client_using;
	//} SETCLIENTID4res_u;
}

type PUTROOTFH4res struct {
	Status int32
}

type PUTFH4res struct {
	Status int32
}

type GETFH4res struct {
	Status int32  `xdr:"union"`
	FH     FH4    `xdr:"unioncase=0"`// nfs_fh4
}

type GETATTR4res struct {
	Status int32 `xdr:"union"`
	Attr   FAttr4 `xdr:"unioncase=0"`
}

type SETCLIENTID_CONFIRM4res struct {
	Status int32
}

// entry4
type DirListEntry struct  {
	Cookie   uint64 // nfs_cookie4
	Name     string // component4
	Attrs    FAttr4
	// this is the trick:
	// in RFC here is the
	// entry4 *nextentry;
	// but I want it to be prepended by length (1 or 0)
	NexEntry []DirListEntry
}

// dirlist4
type DirList4 struct {
	Entries []DirListEntry
	Eof bool
}

type READDIR4resok struct {
	Cookie Verifier4
	DirList DirList4
}

type READDIR4res struct {
	Status int32 `xdr:"union"`
	Result READDIR4resok `xdr:"unioncase=0"`
}

// change_info4
type ChangeInfo4 struct {
	Atomic bool   //bool_t atomic;
	Before uint64 //changeid4 before;
	After  uint64 //changeid4 after;
}

// open_read_delegation4
type OpenReadDelegation4 struct {
	//stateid4 stateid;
	//bool_t recall;
	//nfsace4 permissions;
}

// open_write_delegation4
type OpenWriteDelegation4 struct {
	//stateid4 stateid;
	//bool_t recall;
	//nfs_space_limit4 space_limit;
	//nfsace4 permissions;
}

// open_delegation4
type OpenDelegation4 struct {
	Type int32 `xdr:"union"`
	Read OpenReadDelegation4 `xdr:"unioncase=1"`
	Write OpenWriteDelegation4 `xdr:"unioncase=2"`
}

type OPEN4resok struct {
	StateId StateId4 //stateid4 stateid;
	ChangeInfo ChangeInfo4 //change_info4 cinfo;
	RFlags uint32 //uint32_t rflags;
	Bitmap []uint32 //bitmap4 attrset;
	Delegation OpenDelegation4 //open_delegation4 delegation;
}

type OPEN4res struct {
	Status int32  `xdr:"union"`
	Result OPEN4resok `xdr:"unioncase=0"`
}

type OPEN_CONFIRM4resok struct {
	State StateId4
}

type OPEN_CONFIRM4res struct {
	Status int32  `xdr:"union"`
	Result OPEN_CONFIRM4resok `xdr:"unioncase=0"`
}

type LOOKUP4res struct {
	Status int32
}

type CREATE4resok struct {
	ChangeInfo ChangeInfo4
	Bitmap     []uint32
}

type CREATE4res struct {
	Status int32        `xdr:"union"`
	Result CREATE4resok `xdr:"unioncase=0"`
}

type WRITE4resok struct {
	Count     uint32
	Committed int32     // stable_how4
	Verifier  Verifier4
}

type WRITE4res struct {
	Status int32       `xdr:"union"`
	Result WRITE4resok  `xdr:"unioncase=0"`
}

type CLOSE4res struct {
	Status int32 `xdr:"union"`
	StateId StateId4 `xdr:"unioncase=0"`
}

type LOCK4denied struct {
	Offset uint64
	Length uint64
	LockType int32 // nfs_lock_type4
	Owner LockOwner4
}

type LOCK4resok struct {
	StateId StateId4
}

type LOCK4res struct {
	Status int32 `xdr:"union"`
	Result LOCK4resok `xdr:"unioncase=0"`
	Denied	LOCK4denied `xdr:"unioncase=1"` // TODO
}

type LOCKT4res struct {
	Status int32       `xdr:"union"`
	Denied LOCK4denied `xdr:"unioncase=10010"` // NFS4ERR_DENIED
}

// nfs_resop4
type NfsResOp4 struct {
	ResOp              uint32                  `xdr:"union"`
	Close              CLOSE4res               `xdr:"unioncase=4"`
	Create             CREATE4res              `xdr:"unioncase=6"`
	GetAttr            GETATTR4res             `xdr:"unioncase=9"`
	GetFH              GETFH4res               `xdr:"unioncase=10"`
	Lock               LOCK4res                `xdr:"unioncase=12"`
	LockT              LOCKT4res               `xdr:"unioncase=13"`
	Lookup             LOOKUP4res              `xdr:"unioncase=15"`
	Open               OPEN4res                `xdr:"unioncase=18"`
	OpenConfirm        OPEN_CONFIRM4res        `xdr:"unioncase=20"`
	PutFH              PUTFH4res               `xdr:"unioncase=22"`
	PutRootFH          PUTROOTFH4res           `xdr:"unioncase=24"`
	ReadDir            READDIR4res             `xdr:"unioncase=26"`
	SetClientId	       SETCLIENTID4res         `xdr:"unioncase=35"`
	SetClientIdConfirm SETCLIENTID_CONFIRM4res `xdr:"unioncase=36"`
	Write              WRITE4res               `xdr:"unioncase=38"`
}

func GetResStatus(res *NfsResOp4) (int32) {
	// res != nil
	switch res.ResOp {
	case OP_CLOSE:
		return res.Close.Status
	case OP_CREATE:
		return res.Create.Status
	case OP_GETATTR:
		return res.GetAttr.Status
	case OP_GETFH:
		return res.GetFH.Status
	case OP_LOOKUP:
		return res.Lookup.Status
	case OP_LOCK:
		return res.Lock.Status
	case OP_LOCKT:
		return res.LockT.Status
	case OP_OPEN:
		return res.Open.Status
	case OP_OPEN_CONFIRM:
		return res.OpenConfirm.Status
	case OP_PUTFH:
		return res.PutFH.Status
	case OP_PUTROOTFH:
		return res.PutRootFH.Status
	case OP_READDIR:
		return res.ReadDir.Status
	case OP_SETCLIENTID:
		return res.SetClientId.Status
	case OP_SETCLIENTID_CONFIRM:
		return res.SetClientIdConfirm.Status
	case OP_WRITE:
		return res.Write.Status
	default:
		return NFS4ERR_INVAL
	}
}

type COMPOUND4res struct {
	Status    int32
	Tag       string
	ResArray  []NfsResOp4
}