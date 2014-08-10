package lips_test

import (
	"github.com/CapillarySoftware/goforward/messaging"
	. "github.com/CapillarySoftware/gomasticate/lips"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("Lips", func() {

	It("Test simple message through the lips", func() {
		chewChan := make(chan *messaging.Food, 1000)
		// go mouth.OpenWide(chewChan)
		//create and send nano message protobuffer

		//then receive on chewChan and get expected message back
	})
})
