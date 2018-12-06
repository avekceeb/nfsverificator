package v42tests

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/avekceeb/nfsverificator/common"
	"flag"
)

func init() {
	flag.Parse()
	Config = ParseOptions()
}

func TestSanity(t *testing.T) {
	Config.SuiteName = "NFSv4.2"
	RunSpecs(t, "NFSv4.2")
}

