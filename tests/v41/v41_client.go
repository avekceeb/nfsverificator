package v41tests

import (
	"math/rand"
	"errors"
	"fmt"
	"github.com/avekceeb/nfsverificator/rpc"
	"github.com/avekceeb/nfsverificator/xdr"
	. "github.com/avekceeb/nfsverificator/v41"
	. "github.com/avekceeb/nfsverificator/common"
    . "github.com/onsi/gomega"
)

const (
	Nfs4ClientError = 10099
)

var (
	BeOK = Equal(int32(NFS4_OK))
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

////////// helpers ///////////////

func Uint64ToVerifier(r uint64) (Verifier4) {
	return Verifier4{
		byte(r & 0xff), byte((r & 0xff00) >> 8),
		byte((r & 0xff0000) >> 16), byte((r & 0xff000000) >> 24),
		byte((r & 0xff000000) >> 32), byte((r & 0xff0000000000) >> 40),
		byte((r & 0xff000000000000) >> 48), byte((r & 0xff00000000000000) >> 56),
	}
}

func LastRes(res *([]NfsResop4)) (*NfsResop4) {
	return &((*res)[len(*res)-1])
}

//////////////////////////////////

type ArgArrayT struct {
	Args []NfsArgop4
}

type CompoundMessage struct {
	Head rpc.Header
	Args COMPOUND4args
}

type NFSv41Client struct {
	RpcClient      *rpc.Client
	Auth           rpc.Auth
	AuthSys        AuthsysParms
	ClientId       uint64
	Seq            uint32 // TODO ??
	Verifier       Verifier4
	Id             string
	// TODO: per server / per session
	// Now only one session
	Sid            Sessionid4
	// Now only fore channel
	ForeChAttr         ChannelAttrs4
}

func (cli *NFSv41Client) MockReboot() {
	cli.Seq = 0
	cli.Verifier = Uint64ToVerifier(rand.Uint64())
}

func (cli *NFSv41Client) Close() {
	cli.RpcClient.Close()
}

func (cli *NFSv41Client) Compound(args ...NfsArgop4) (reply COMPOUND4res, err error) {
	res, err := cli.RpcClient.Call(CompoundMessage{
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
			Minorversion: 1,
			Argarray: args,
		},
	})
	if nil != err {
		return COMPOUND4res{Status:Nfs4ClientError}, err
	}
	if nil == res {
		return COMPOUND4res{Status:Nfs4ClientError}, errors.New("Reply is nil")
	}
	// Parse reply at last
	err = xdr.Read(res, &reply)
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

func (cli*NFSv41Client) Null() (error) {
	res, err := cli.RpcClient.Call(rpc.Header{
		Rpcvers: 2,
		Prog:    NFS4_PROGRAM,
		Vers:    NFS_V4,
		Proc:    NFSPROC4_NULL,
		Cred:    cli.Auth,
		Verf:    rpc.AuthNull,
	})
	if nil != err {
		return err
	}
	if nil == res {
		return errors.New("rpc returned nil")
	}
	var b []byte
	res.Read(b)
	if len(b) != 0 {
		return errors.New("NFSv4.1 NULL returned non empty")
	}
	return nil
}


func NewNFSv41Client(srvHost string, srvPort int, authHost string, uid uint32, gid uint32, cid string) (*NFSv41Client) {
	client := NFSv41Client{}
	u := rpc.NewAuthUnix(authHost, uid, gid)
	client.Auth = u.Auth()
	client.AuthSys = AuthsysParms{Stamp:u.Stamp, Uid:uid, Gid:gid, Machinename:authHost, GidLen:0}
	var err error
	client.RpcClient, err = rpc.DialService(srvHost, srvPort)
	if err != nil {
		panic(err.Error())
	}
	client.Id = cid
	client.MockReboot()
	return &client
}

func (t *NFSv41Client) Pass(args ...NfsArgop4) ([]NfsResop4) {
	reply, err := t.Compound(args...)
	Expect(err).To(BeNil())
	Expect(reply.Status).To(BeOK)
	// TODO: ???
	//for _, k := range reply.Resarray {
	//    Expect(GetStatus(&k)).To(Equal(NFS4_OK))
	//}
	return reply.Resarray
}


func (t *NFSv41Client) Fail(stat int32, args ...NfsArgop4) ([]NfsResop4) {
	res, err := t.Compound(args...)
	Expect(err).To(BeNil())
	Expect(res.Status).To(Equal(int32(stat)))
	return res.Resarray
}


func (cli *NFSv41Client) ExchangeId() {
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

func (cli *NFSv41Client) CreateSession() {
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
	//cli.Pass(
	//	Sequence(cli.Sid, cli.Seq, 0, 0, false),
	//	Putrootfh(),
	//	SecinfoNoName(0))
	//cli.Seq++
	//l := cli.Pass(
	//	Sequence(cli.Sid, cli.Seq, 0, 0, false),
	//	Putrootfh(),
	//	Getfh(),
	//	Getattr([]uint32{MakeGetAttrFlags(FATTR4_LEASE_TIME)}))
	//leaseTime := LastRes(&l).Opgetattr.Resok4.ObjAttributes.AttrVals
	//fmt.Println(BytesToUint32(leaseTime))
}