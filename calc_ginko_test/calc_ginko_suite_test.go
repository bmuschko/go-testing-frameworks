package calc_ginko_test

import (
	"testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCalc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Calculator Suite")
}