package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/nfs40"
	"github.com/avekceeb/nfsverificator/rpc"
	"github.com/vmware/go-nfs-client/nfs/xdr"
	"fmt"
)

var (
	rpcClient      *rpc.Client
	auth           rpc.Auth
	clientId       uint64
	verifier       [NFS4_VERIFIER_SIZE]byte
)

func Compoundv40(args ...NfsArgOp4) ([]NfsResOp4) {
	res, err := rpcClient.Call(CompoundMessage{
		Head: rpc.Header{
			Rpcvers: 2,
			Prog:    NFS4_PROGRAM,
			Vers:    NFS_V4,
			Proc:    NFSPROC4_COMPOUND,
			Cred:    auth,
			Verf:    rpc.AuthNull,
		},
		Args:COMPOUND4args{
			Tag: "",
			MinorVersion: 0,
			ArgArray: ArgArrayT{Args:args},
		},
	})
	Expect(err).To(BeNil())
	Expect(res).ToNot(BeNil())
	var reply COMPOUND4res
	err = xdr.Read(res, &reply)
	fmt.Printf(" REPLY: \n%v\n", reply)
	Expect(err).To(BeNil())
	Expect(reply.Status).To(Equal(int32(NFS4_OK)))
	Expect(len(reply.ResArray)).To(Equal(len(args)))
	// TODO: 0 != 26
	for k := range args {
		Expect(reply.ResArray[k].ResOp).To(Equal(args[k].ArgOp))
	}
	return reply.ResArray
}

var _ = Describe("NFSv4.0", func() {

	auth = rpc.NewAuthUnix(Config.ClientHost, 0, 0).Auth()

	BeforeEach(func() {
		var err error
		rpcClient, err = rpc.DialService(Config.ServerHost, Config.ServerPort)
		//defer client.Close()
		if err != nil {
			panic(err.Error())
		}
		// TODO: check client created OK
	})

	AfterEach(func() {
		rpcClient.Close()
	})

	Context("Basic", func() {

		It("NULL Call", func() {
			res, err := rpcClient.Call(rpc.Header{
				Rpcvers: 2,
				Prog:    NFS4_PROGRAM,
				Vers:    NFS_V4,
				Proc:    NFSPROC4_NULL,
				Cred:    auth,
				Verf:    rpc.AuthNull,
			})

			Expect(err).To(BeNil())
			Expect(res).ToNot(BeNil())
			var b []byte
			res.Read(b)
			Expect(len(b)).To(Equal(0))
		})

		It("New Client", func() {
			id := RandString(8)
			cb_cl := NfsClientId{
				Verifier:[NFS4_VERIFIER_SIZE]byte{
					0x04, 0x05, 0x06, 0x07, 0x08, 0x19, 0x1a, 0x1b},
				Id:id}
			// TODO: real callback server
			cb := CallbackClient{
				Program:0x40000000,
				Location: ClientAddr{NetId:"tcp", Addr:"127.0.0.1.139.249"}}

			a := SETCLIENTID4args{
				Client:cb_cl,
				Callback:cb,
				CallbackIdent:1}

			ret := Compoundv40(NfsArgOp4{ArgOp:OP_SETCLIENTID, SetClientId:a})

			Expect(ret[0].SetClientId.Status).To(Equal(int32(NFS4_OK)))
			clientId = ret[0].SetClientId.ResOk.ClientId
			verifier = ret[0].SetClientId.ResOk.Verifier

			ret = Compoundv40(NfsArgOp4{
					ArgOp:OP_SETCLIENTID_CONFIRM,
					SetClientIdConfirm: SETCLIENTID_CONFIRM4args{
						ClientId: clientId,
						Verifier: verifier},
			})
			//Expect(ret[0].SetClientIdConfirm. TODO: Status).To(Equal())

			// putrootfh | getfh | getattr
			ret = Compoundv40(GetPutRootFH(),
					GetGetFH(),
					GetGetAttr(FATTR4_MODE, FATTR4_SIZE))

			Expect(ret[0].GetFH.Status).To(Equal(int32(NFS4_OK)))
			var fh string
			fh = ret[1].GetFH.FH

			// putfh | readdir
			ret = Compoundv40(GetPutFH(fh),	GetReadDir())
			// TODO
			//for _, e := range ret[1].ReadDir.Result.DirList.Entries {
			//	fmt.Printf(" entry: %s\n", e.Name)
			//}

			// Create File:
			ret = Compoundv40([]NfsArgOp4{
				GetPutFH(fh),
				{ArgOp:OP_OPEN, Open: OPEN4args{SeqId:0,
					ShareAccess: OPEN4_SHARE_ACCESS_WRITE,
					ShareDeny: OPEN4_SHARE_DENY_NONE,
					OpenOwner:OpenOwner4{ClientId:clientId, Owner:id},
					OpenHow: OpenFlag4{OpenType:OPEN4_CREATE,
						CreateHow: CreateHowT{CreateMode:UNCHECKED4,
							Attr:FAttr4{
								Bitmap: GetBitmap(FATTR4_MODE),
								AttrList:"\x00\x00\x01\xa4"},
						},
					},
					Claim: OpenClaim4{Claim:CLAIM_NULL, File: RandString(8)}},
				},
			}...)

		})
	})
})
