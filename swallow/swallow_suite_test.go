package swallow_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSwallow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Swallow Suite")
}
