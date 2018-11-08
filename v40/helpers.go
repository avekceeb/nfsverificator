package v40

func FhFromString(h string) (NfsFh4) {
    return NfsFh4([]byte(h))
    //copy(fh[:], s[:NFS4_FHSIZE])
    //return fh
}
