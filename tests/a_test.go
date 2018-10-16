package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/avekceeb/nfsverificator/nfs40"
	"github.com/avekceeb/nfsverificator/rpc"
	"io"
	"github.com/vmware/go-nfs-client/nfs/xdr"
)

var (
	rpcClient      *rpc.Client
	auth           *rpc.AuthUnix
	rpcNfs40Header rpc.Header
	res            io.ReadSeeker
	err            error
	reply          COMPOUND4res
	clientId       uint64
	verifier       [NFS4_VERIFIER_SIZE]byte
)

var _ = Describe("NFSv4.0", func() {

	auth = rpc.NewAuthUnix(Config.ClientHost, 0, 0)
	rpcNfs40Header = rpc.Header{
		Rpcvers: 2,
		Prog:    NFS4_PROGRAM,
		Vers:    NFS_V4,
		Proc:    NFSPROC4_NULL,
		Cred:    auth.Auth(),
		Verf:    rpc.AuthNull,
	}

	BeforeEach(func() {
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
			rpcNfs40Header.Proc = NFSPROC4_NULL
			res, err = rpcClient.Call(rpcNfs40Header)
			Expect(err).To(BeNil())
			Expect(res).ToNot(BeNil())
			var b []byte
			res.Read(b)
			Expect(len(b)).To(Equal(0))
		})

		It("New Client", func() {
			rpcNfs40Header.Proc = NFSPROC4_COMPOUND
			// TODO : random before each
			id := "sdfsdf sdfsdfs sdfsdfff"
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

			res, err = rpcClient.Call(CompoundMessage{
				Head: rpcNfs40Header,
				Args:COMPOUND4args{
					Tag: "",
					MinorVersion: 0,
					ArgArray: ArgArrayT{Args:[]NfsArgOp4{
						{ArgOp:OP_SETCLIENTID, SetClientId:a},
					}},
				},
			})
			Expect(err).To(BeNil())
			Expect(res).ToNot(BeNil())
			err = xdr.Read(res, &reply)
			Expect(err).To(BeNil())
			Expect(reply.Status).To(Equal(int32(NFS4_OK)))
			Expect(len(reply.ResArray)).To(Equal(1))
			Expect(reply.ResArray[0].ResOp).To(Equal(uint32(OP_SETCLIENTID)))
			Expect(reply.ResArray[0].SetClientId.Status).To(Equal(int32(NFS4_OK)))
			clientId = reply.ResArray[0].SetClientId.ResOk.ClientId
			verifier = reply.ResArray[0].SetClientId.ResOk.Verifier

			res, err = rpcClient.Call(CompoundMessage{
				Head:rpcNfs40Header,
				Args:COMPOUND4args{
					Tag:"",
					MinorVersion: 0,
					ArgArray: ArgArrayT{Args:[]NfsArgOp4{
						{ArgOp:OP_SETCLIENTID_CONFIRM,
							SetClientIdConfirm: SETCLIENTID_CONFIRM4args{
								ClientId: clientId,
								Verifier: verifier,
							},
						},
					}},
				},
			})

			Expect(err).To(BeNil())
			Expect(res).ToNot(BeNil())
			err = xdr.Read(res, &reply)
			Expect(err).To(BeNil())
			Expect(reply.Status).To(Equal(int32(NFS4_OK)))
			Expect(len(reply.ResArray)).To(Equal(1))
			Expect(reply.ResArray[0].ResOp).To(Equal(uint32(OP_SETCLIENTID_CONFIRM)))

			// putrootfh | getfh | getattr
			res, err = rpcClient.Call(CompoundMessage{
				Head:rpcNfs40Header,
				Args:COMPOUND4args{
					Tag:"",
					MinorVersion: 0,
					ArgArray: ArgArrayT{Args:[]NfsArgOp4{
						{ArgOp:OP_PUTROOTFH},
						{ArgOp:OP_GETFH},
						{ArgOp:OP_GETATTR,
							AttrRequest:[]uint32{0x0010011a, 0x00b0a23a}},
					}},
				},
			})
			Expect(err).To(BeNil())
			Expect(res).ToNot(BeNil())
			err = xdr.Read(res, &reply)
			Expect(err).To(BeNil())
			Expect(reply.Status).To(Equal(int32(NFS4_OK)))
			Expect(len(reply.ResArray)).To(Equal(3))
			Expect(reply.ResArray[0].ResOp).To(Equal(uint32(OP_PUTROOTFH)))
			Expect(reply.ResArray[1].ResOp).To(Equal(uint32(OP_GETFH)))
			Expect(reply.ResArray[2].ResOp).To(Equal(uint32(OP_GETATTR)))
			var fh string
			fh = reply.ResArray[1].GetFH.FH

			// putfh | readdir
			res, err = rpcClient.Call(CompoundMessage{
				Head: rpcNfs40Header,
				Args: COMPOUND4args{
					Tag:"",
					MinorVersion: 0,
					ArgArray: ArgArrayT{Args:[]NfsArgOp4{
						{ArgOp:OP_PUTFH, PutFH:PUTFH4args{FH:fh}},
						{ArgOp:OP_READDIR, ReadDir:READDIR4args{
							Cookie:0,
							Verifier:[NFS4_VERIFIER_SIZE]byte{0, 0, 0, 0, 0, 0, 0, 0},
							Dircount:8170,
							Count:32680,
							Bitmap:[]uint32{0x0018091a, 0x00b0a23a},
						}},
					}},
				},
			})

		})


	})
})
