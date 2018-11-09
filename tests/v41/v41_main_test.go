package v41tests

import (
    "testing"
    "flag"
    . "github.com/onsi/ginkgo"
    . "github.com/avekceeb/nfsverificator/common"
)

func init() {
    flag.Parse()
    Config = ReadConfig(ConfigFile)
}

func TestSanity(t *testing.T) {
    RunSpecs(t, "NFSv4.1")
}

