package v40tests

import (
    . "github.com/onsi/ginkgo"
    "testing"
    "flag"
    . "github.com/avekceeb/nfsverificator/common"
)

func init() {
    flag.Parse()
    Config = ReadConfig(ConfigFile)
}

func TestSanity(t *testing.T) {
    RunSpecs(t, "NFSv4.0")
}
