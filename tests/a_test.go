package tests

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"fmt"
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
		fmt.Println(">>>>>>>>>>>>>>>>>")
		rpcClient, err = rpc.DialService(Config.ServerHost, Config.ServerPort)
		//defer client.Close()
		if err != nil {
			panic(err.Error())
		}
		// TODO: check client created OK
	})

	AfterEach(func() {
		rpcClient.Close()
		fmt.Println("<<<<<<<<<<<<<<<<<")
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
			id := "blah-blah-blah"
			cb_cl := NfsClientId{
				Verifier:[NFS4_VERIFIER_SIZE]byte{
					0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b},
				Id:id}

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
					Tag:"",
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
		})

	})
})
