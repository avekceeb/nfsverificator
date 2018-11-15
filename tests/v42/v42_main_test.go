package v42tests

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
    Config.SuiteName = "NFSv4.2"
    RunSpecs(t, "NFSv4.2")
}

