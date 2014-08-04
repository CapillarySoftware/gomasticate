package chew_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestChew(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Chew Suite")
}
