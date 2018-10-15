package tests

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
	. "github.com/avekceeb/nfsverificator/util"
	"path/filepath"
	"os"
)

var Config TestConfig

func init() {
	// TODO: option to provide non-default file
    configPath := filepath.Join(os.Getenv("GOPATH"),
        "src/github.com/avekceeb/nfsverificator/config.json")
	Config = ReadConfig(configPath)
}

func TestSanity(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Sanity")
}
