package v41tests

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
    "flag"
    . "github.com/avekceeb/nfsverificator/common"
)

func init() {
    flag.Parse()
    Config = ReadConfig(ConfigFile)
}

func TestSanity(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "NFSv4.1")
}

