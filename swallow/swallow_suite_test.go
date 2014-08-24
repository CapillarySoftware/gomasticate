package swallow

import (
	gi "github.com/onsi/ginkgo"
	gom "github.com/onsi/gomega"

	"testing"
)

func TestSwallow(t *testing.T) {
	gom.RegisterFailHandler(gi.Fail)
	gi.RunSpecs(t, "Swallow Suite")
}
