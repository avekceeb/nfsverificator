# NFS Server tests
Based on github.com/vmware/go-nfs-client using github.com/rasky/go-xdr/xdr2 


## Installation

#### Install golang

#### Install Ginkgo
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega/...
	cd ${GOPATH}/src/github.com/onsi/ginkgo/ginkgo
	go install
	export PATH=${PATH}:${GOPATH}/bin

#### Run Tests
    make all

## Features To Be Done
- receive CallBacks (at least CB_NULL)
- select secure/insecure port
- config: provide user
- multi-server
- config: provide server reboot command