package tests

import (
	"github.com/onsi/ginkgo"
	"strings"
	"fmt"
)

type ConstToString func(int32) string

type Assertion struct {
	ErrorName ConstToString
}

func (a *Assertion) Assert(condition bool, errMessage string) {
	if ! condition {
		ginkgo.Fail(errMessage)
	}
}

func (a *Assertion) AssertStatus(actual int32, expected int32) {
	a.Assert(actual == expected,
		fmt.Sprintf("Expected: %s  Got: %s",
			a.ErrorName(expected), a.ErrorName(actual)))
}

func (a *Assertion) AssertStatusOneOf(actual int32, expected []int32) {
	list := []string{}
	for _, err := range expected {
		list = append(list, a.ErrorName(err))
		if actual == err {
			return
		}
	}
	ginkgo.Fail(fmt.Sprintf("Expected one of: %s  Got: %s",
			strings.Join(list, ", "), a.ErrorName(actual)))
}

func (a *Assertion) AssertNfsOK(actual int32) {
	a.AssertStatus(actual, 0 /*NFS4_OK*/)
}

func (a *Assertion) AssertNoErr(err error) {
	if err != nil {
		ginkgo.Fail(fmt.Sprintf(
			"Unexpected error: %s", err.Error()))
	}
}

