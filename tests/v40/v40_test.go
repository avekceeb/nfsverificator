package v40tests

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
    . "github.com/avekceeb/nfsverificator/util"
    "path/filepath"
    "os"
    "math/rand"
    "time"
    "flag"
)

var (
    Config TestConfig
)

func init() {
    var configFile string
    flag.StringVar(&configFile, "config",
        filepath.Join(os.Getenv("GOPATH"),
        "src/github.com/avekceeb/nfsverificator/config.json"), "Config File")
    flag.Parse()
    Config = ReadConfig(configFile)
    rand.Seed(time.Now().UnixNano())
}

func TestSanity(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "NFSv4.0")
}
