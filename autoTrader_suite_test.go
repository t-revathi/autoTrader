package autoTrader_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAutoTrader(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AutoTrader Suite")
}
