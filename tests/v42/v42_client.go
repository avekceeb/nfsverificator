package v42tests

// TODO: may be it is 99% identical to v41...

import (
	"errors"
	"fmt"
	"github.com/avekceeb/nfsverificator/rpc"
	"github.com/avekceeb/nfsverificator/xdr"
	. "github.com/avekceeb/nfsverificator/tests"
	. "github.com/avekceeb/nfsverificator/v42"
	. "github.com/avekceeb/nfsverificator/common"
	"math/rand"
)

const (
	Nfs4ClientError = 10099
)

var (
	Assert42 Assertion
	DefProtect = StateProtect4A{SpaHow:0}
	DefImpl = []NfsImplID4{{
				NiiDate:Nfstime4{Seconds: 0, Nseconds: 0},
				NiiDomain:"kernel.org",
				NiiName:"Linux"}}
	DefExchgFlags = MakeUint32Flags(
		EXCHGID4_FLAG_BIND_PRINC_STATEID,
		EXCHGID4_FLAG_SUPP_MOVED_MIGR,
		EXCHGID4_FLAG_SUPP_MOVED_REFER)
	DefCsFlags = MakeUint32Flags(
		CREATE_SESSION4_FLAG_PERSIST,
		CREATE_SESSION4_FLAG_CONN_BACK_CHAN)
	DefChannelAttrs = ChannelAttrs4{
		CaHeaderpadsize:0,
		CaMaxrequestsize:1049620,
		CaMaxresponsesize:1049480,
		CaMaxresponsesizeCached:3428,
		CaMaxoperations:8,
		CaMaxrequests:64,
	}
)

func init() {
	Assert42 = Assertion{ErrorName:ErrorNameNfs42}
}

type ArgArrayT struct {
	Args []NfsArgop4
}

type CompoundMessage struct {
	Head rpc.Header
	Args COMPOUND4args
}

type NFSv42Client struct {
	*Nfs4Client
	Verifier       Verifier4
	// TODO: per server / per session
	// Now only one session
	Sid            Sessionid4
	EidFlags       uint32
	// Now only fore channel
	ForeChAttr     ChannelAttrs4
}

func (cli *NFSv42Client) Compound(args ...NfsArgop4) (reply COMPOUND4res, err error) {
	msg := CompoundMessage{
		Head: rpc.Header{
			Rpcvers: 2,
			Prog:    NFS4_PROGRAM,
			Vers:    NFS_V4,
			Proc:    NFSPROC4_COMPOUND,
			Cred:    cli.Auth,
			Verf:    rpc.AuthNull,
		},
		Args: COMPOUND4args{
			Tag: "",
			Minorversion: 2,
			Argarray: args,
		},
	}
	cli.Trace(msg)
	res, err := cli.RpcClient.Call(msg)
	if nil != err {
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	if nil == res {
		return COMPOUND4res{Status:Nfs4ClientError}, errors.New("Reply is nil")
	}
	// Parse reply at last
	err = xdr.Read(res, &reply)
	cli.Trace(reply)
	if nil != err {
		fmt.Printf("%s", err.Error())
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	if args[0].Argop == OP_SEQUENCE {
		if reply.Resarray[0].Opsequence.SrStatus == int32(NFS4_OK) {
			cli.Seq++
		}
	}

	//for _, a := range args {
    //    if a.Argop == OP_SEQUENCE { // TODO: other ops?
		//	cli.Seq++
		//	break
    //    }
    //}
	return reply, nil
}

func NewNFSv41Client(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string, trace bool) (*NFSv42Client) {
	r := rand.Uint64()
	client := NFSv42Client{NewNfs4Client(
		srvHost, srvPort, authHost, uid, gid, cid, trace), Verifier4{
			byte(r&0xff), byte((r&0xff00)>>8),
			byte((r&0xff0000)>>16), byte((r&0xff000000)>>24),
			byte((r&0xff000000)>>32), byte((r&0xff0000000000)>>40),
			byte((r&0xff000000000000)>>48), byte((r&0xff00000000000000)>>56),
	}, Sessionid4{}, 0, ChannelAttrs4{}}
	return &client
}

func NewNFSv42DefaultClient() (*NFSv42Client) {
	return NewNFSv41Client(Config.Server, Config.Port,
		RandString(8) + ".fake.net", 0, 0, RandString(8), Config.Trace)
}

// HELPERS FOR CHECK COMPOUND STATUS //

func (t *NFSv42Client) Pass(args ...NfsArgop4) ([]NfsResop4) {
	reply, err := t.Compound(args...)
	Assert42.AssertNoErr(err)
	Assert42.AssertNfsOK(reply.Status)
	// Note: not checking every op in compound, only overall status
	return reply.Resarray
}


func (t *NFSv42Client) Fail(stat int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	Assert42.AssertNoErr(err)
	Assert42.AssertStatus(res.Status, stat)
	return res.Resarray
}

func (cli *NFSv42Client) ExchangeId() {
	r := cli.Pass(ExchangeId(
		ClientOwner4{
			CoOwnerid: RandString(14),
			CoVerifier: Verifier4{}},
		DefExchgFlags,
		DefProtect,
		DefImpl))
	ei := LastRes(&r).OpexchangeID.EirResok4
	cli.ClientId = ei.EirClientid
	cli.Seq = ei.EirSequenceid
}

func (cli *NFSv42Client) CreateSession() {
	/*
As previously stated, CREATE_SESSION can be sent with or without a
preceding SEQUENCE operation.  Even if a SEQUENCE precedes
CREATE_SESSION, the server MUST maintain the CREATE_SESSION reply
cache, which is separate from the reply cache for the session
associated with a SEQUENCE.  If CREATE_SESSION was originally sent by
itself, the client MAY send a retry of the CREATE_SESSION operation
within a COMPOUND preceded by a SEQUENCE.  If CREATE_SESSION was
originally sent in a COMPOUND that started with a SEQUENCE, then the
client SHOULD send a retry in a COMPOUND that starts with a SEQUENCE
that has the same session ID as the SEQUENCE of the original request.
However, the client MAY send a retry in a COMPOUND that either has no
preceding SEQUENCE, or has a preceding SEQUENCE that refers to a
different session than the original CREATE_SESSION.  This might be
necessary if the client sends a CREATE_SESSION in a COMPOUND preceded
by a SEQUENCE with session ID X, and session X no longer exists.
Regardless, any retry of CREATE_SESSION, with or without a preceding
SEQUENCE, MUST use the same value of csa_sequence as the original.
*/
	s := cli.Pass(CreateSession(
		cli.ClientId,
		cli.Seq,
		DefCsFlags,
		DefChannelAttrs,
		DefChannelAttrs,
		0x40000000, // CallBack Program
		[]CallbackSecParms4{{
			CbSecflavor:1,
			CbspSysCred:cli.AuthSys}}))
	cli.Sid = LastRes(&s).OpcreateSession.CsrResok4.CsrSessionid
	// TODO: now only fore channel
	cli.ForeChAttr = LastRes(&s).OpcreateSession.CsrResok4.CsrForeChanAttrs
	/*
   Once the session is created, the first SEQUENCE or CB_SEQUENCE
   received on a slot MUST have a sequence ID equal to 1; if not, the
   replier MUST return NFS4ERR_SEQ_MISORDERED.

	BTW, Linux disregards this

 */

	cli.Seq = 1

	cli.Pass(
		Sequence(cli.Sid, cli.Seq, 0, 0, false),
		ReclaimComplete(false))
}

func (cli *NFSv42Client) GetSomeAttr() {
	cli.Pass(
		Sequence(cli.Sid, cli.Seq, 0, 0, false),
		Putrootfh(),
		SecinfoNoName(0))
	// TODO
	l := cli.Pass(
		Sequence(cli.Sid, cli.Seq, 0, 0, false),
		Putrootfh(),
		Getfh(),
		Getattr(MakeGetAttrFlags(FATTR4_LEASE_TIME)))
	cli.LeaseTime = BytesToUint32(LastRes(&l).Opgetattr.Resok4.ObjAttributes.AttrVals)
	//r := cli.Pass(
	//	Sequence(cli.Sid, cli.Seq, 0, 0, false),
	//	Putrootfh(),
	//	Access(MakeUint32Flags(ACCESS4_DELETE, ACCESS4_EXTEND, ACCESS4_LOOKUP, ACCESS4_MODIFY,
	//		ACCESS4_READ, ACCESS4_EXECUTE)),
	//)
	//access := LastRes(&r).Opaccess.Resok4.Access
	//cli.DL = CheckFlag(access, ACCESS4_DELETE)
	//cli.XT = CheckFlag(access, ACCESS4_EXTEND)
	//cli.LU = CheckFlag(access, ACCESS4_LOOKUP)
	//cli.MD = CheckFlag(access, ACCESS4_MODIFY)
	//cli.RD = CheckFlag(access, ACCESS4_READ)
	// TODO : execute ???

}

// HELPERS FOR BUILDING NfsArgop4 FOR COMPOUNDS //

func (t *NFSv42Client) SequenceArgs() NfsArgop4 {
	return Sequence(t.Sid, t.Seq, 0, 0, false)
}
