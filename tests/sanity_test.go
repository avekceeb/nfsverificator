package tests

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
    "testing"
	. "github.com/avekceeb/nfsverificator/util"
	"path/filepath"
	"os"
	"math/rand"
	"time"
)

var Config TestConfig

func init() {
    rand.Seed(time.Now().UnixNano())
	// TODO: option to provide non-default file
    configPath := filepath.Join(os.Getenv("GOPATH"),
        "src/github.com/avekceeb/nfsverificator/config.json")
	Config = ReadConfig(configPath)
}

const letters = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

func RandString(n int) string {
    var l int64 = int64(len(letters))
    b := make([]byte, n)
    for i := 0; i < n; i++ {
        b[i] = letters[int(rand.Int63n(l))]
    }
    return string(b)
}

func TestSanity(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "Sanity")
}
