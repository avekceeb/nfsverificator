package nfs40

func GetPutFH(fh string) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_PUTFH, PutFH:PUTFH4args{FH:fh}}
}

func GetReadDir() (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_READDIR, ReadDir:READDIR4args{
		Cookie: 0,
		Verifier: [NFS4_VERIFIER_SIZE]byte{0, 0, 0, 0, 0, 0, 0, 0},
		Dircount: 8170,
		Count: 32680,
		Bitmap: GetBitmap(FATTR4_TYPE)},
		//Bitmap: []uint32{0x0018091a, 0x00b0a23a}},
	}
}

func GetPutRootFH() (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_PUTROOTFH}
}

func GetGetFH() (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_GETFH}
}

func GetGetAttr(bits ...int) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_GETATTR,
					AttrRequest:GetBitmap(bits...)}
}

func GetBitmap(bits ...int) ([]uint32) {
	b := []uint32{0,0}
	// it will panic in case of bit > 64
	for _, v := range bits {
		b[v/32] |= (1 << uint32(v%32))
	}
	return b
}
