package start_test

import (
	. "github.com/CapillarySoftware/goforward/start"
	sys "github.com/CapillarySoftware/goforward/syslogService"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Start", func() {
	Describe("Valid tests", func() {
		It("Protocol switch TCP", func() {
			proto := ProcessProtocol("tcP")
			Expect(proto).Should(Equal(sys.TCP))

		})

		It("Protocol switch UDP", func() {
			proto := ProcessProtocol("UdP")
			Expect(proto).Should(Equal(sys.UDP))

		})

	})

})
