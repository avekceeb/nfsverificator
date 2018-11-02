package v40

/*
func GetStatus(res *NfsResop4) (int32) {
    if res == nil {
		return NFS4ERR_INVAL
	}
	// TODO: reset of ops ...
    switch res.Resop {
    case OP_CLOSE:
        return res.Opclose.Status
    case OP_CREATE:
        return res.Opcreate.Status
    case OP_GETATTR:
        return res.Opgetattr.Status
    case OP_GETFH:
        return res.Opgetfh.Status
    case OP_LOOKUP:
        return res.Oplookup.Status
    case OP_LOCK:
        return res.Oplock.Status
    case OP_LOCKT:
        return res.Oplockt.Status
    case OP_OPEN:
        return res.Opopen.Status
    case OP_OPEN_CONFIRM:
        return res.OpopenConfirm.Status
    case OP_PUTFH:
        return res.Opputfh.Status
    case OP_PUTROOTFH:
        return res.Opputrootfh.Status
    case OP_READDIR:
        return res.Opreaddir.Status
    case OP_SETCLIENTID:
        return res.Opsetclientid.Status
    case OP_SETCLIENTID_CONFIRM:
        return res.OpsetclientidConfirm.Status
    case OP_WRITE:
        return res.Opwrite.Status
    default:
        return NFS4ERR_INVAL
    }
}
*/

func GetBitmap(bits ...int) ([]uint32) {
    b := []uint32{0,0}
    // it will panic in case of bit > 64
    for _, v := range bits {
        b[v/32] |= (1 << uint32(v%32))
    }
    return b
}

func GetPermAttrList(perm uint) (l []byte) {
    l = make([]byte, 4)
    l[3] = byte(perm & 0xff)
    l[2] = byte((perm & 0xff00) >> 8)
    l[1] = byte((perm & 0xff0000) >> 16)
    l[0] = byte((perm & 0xff000000) >> 24)
    return l
}

func FhFromString(h string) (NfsFh4) {
    return NfsFh4([]byte(h))
    //copy(fh[:], s[:NFS4_FHSIZE])
    //return fh
}
