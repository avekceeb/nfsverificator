# NFS Server tests
Based on github.com/vmware/go-nfs-client using github.com/rasky/go-xdr/xdr2 

## Quick Start

## Installation

#### Install golang

#### Install Ginkgo
	go get github.com/onsi/ginkgo/ginkgo
	# TODO: this is no longer needed
	go get github.com/onsi/gomega/...
	cd ${GOPATH}/src/github.com/onsi/ginkgo/ginkgo
	go install
	export PATH=${PATH}:${GOPATH}/bin

#### Run Tests
    make all
    make focus="PyNFS" v41
    make skip="BUG"
    make config="/path/to/custom.json"

## Features To Be Done
- receive CallBacks (at least CB_NULL)
- select secure/insecure port
- config: provide user
- multi-server
- config: provide server reboot command
- option to run test on every share provided by default server (exported with different options)
- various UTF8 data
- check if NFS4ERR_GRACE before suite?
- coverage (or at least nfsstat)
- config entries for refer option
- trace: print crc32 for fh
- wipe out inUnion from spew::dump