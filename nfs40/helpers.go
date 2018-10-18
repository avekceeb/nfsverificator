package nfs40

func GetPutFH(fh FH4) (NfsArgOp4) {
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

func GetLookup(file string) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_LOOKUP, Lookup: LOOKUP4args{Name: file}}
}

func GetGetAttr(bits ...int) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_GETATTR,
					AttrRequest:GetBitmap(bits...)}
}

func GetSetAttr(s_id StateId4, bm []uint32, attr []byte) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_SETATTR,
					SetAttr:SETATTR4args{StateId:s_id,
						Attr:FAttr4{Bitmap: bm, AttrList: attr}}}
}

func GetBitmap(bits ...int) ([]uint32) {
	b := []uint32{0,0}
	// it will panic in case of bit > 64
	for _, v := range bits {
		b[v/32] |= (1 << uint32(v%32))
	}
	return b
}

func CreateDir(name string) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_CREATE,
			Create:CREATE4args{CreateType:CreateType4{Type:NF4DIR},
				Name:name,
				Attr:FAttr4{Bitmap:GetBitmap(FATTR4_MODE),
					AttrList: GetPermAttrList(0644)},
			}}
}

func GetPermAttrList(perm uint) (l []byte) {
	l = make([]byte, 4)
	l[3] = byte(perm & 0xff)
	l[2] = byte((perm & 0xff00) >> 8)
	l[1] = byte((perm & 0xff0000) >> 16)
	l[0] = byte((perm & 0xff000000) >> 24)
	return l
}

func GetSetClientConfirm(clientId uint64, verifier Verifier4) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_SETCLIENTID_CONFIRM,
		SetClientIdConfirm: SETCLIENTID_CONFIRM4args{ClientId: clientId,
			Verifier: verifier}}
}

func GetOpen(seq uint32, clientId uint64, owner string, name string) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_OPEN,
		Open: OPEN4args{SeqId:seq,
			ShareAccess: OPEN4_SHARE_ACCESS_WRITE,
			ShareDeny: OPEN4_SHARE_DENY_NONE,
			OpenOwner: OpenOwner4{ClientId: clientId,
				Owner: owner},
			OpenHow: OpenFlag4{OpenType:OPEN4_CREATE,
				CreateHow: CreateHowT{CreateMode:UNCHECKED4,
					Attr:FAttr4{
						Bitmap: GetBitmap(FATTR4_MODE),
						AttrList: GetPermAttrList(0644)},
				},
			},
			Claim: OpenClaim4{Claim:CLAIM_NULL, File: name}},
	}
}

func GetOpenConfirm(stateId StateId4, seq uint32) (NfsArgOp4) {
	return NfsArgOp4{ArgOp:OP_OPEN_CONFIRM,
			OpenConfirm:OPEN_CONFIRM4args{State:stateId, SeqId:seq}}
}

func GetWrite(stateId StateId4, data *[]byte, offset uint64) (NfsArgOp4) {
	return NfsArgOp4{ArgOp: OP_WRITE,
		Write: WRITE4args{State:stateId, Offset: offset, Stable: FILE_SYNC4, Data: *data}}
}

func GetClose(seq uint32, stateId StateId4) (NfsArgOp4) {
	return NfsArgOp4{ArgOp: OP_CLOSE, Close:CLOSE4args{SeqId:seq, StateId: stateId}}
}
