package lips_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMouth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mouth Suite")
}
