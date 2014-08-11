package lips_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestLips(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lips Suite")
}
