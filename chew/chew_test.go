package chew_test

import (
	"github.com/CapillarySoftware/gomasticate/chew"
	log "github.com/cihub/seelog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	// "time"
)

var _ = Describe("Push and chew test", func() {

	Measure("10 pushers single puller performance", func(b Benchmarker) {
		count := 50000
		pushers := 10
		runtime := b.Time("runtime", func() {
			finished := make(chan int, 10)

			go chew.Chew(count*pushers, finished)
			for i := 0; i < pushers; i++ {
				log.Info("Starting pusher")
				go chew.Pusher(count, finished)
			}
			log.Info("Waiting for messages")
			for i := 0; i < pushers+1; i++ {
				log.Info("Checking finished pushers and chewers")
				c := <-finished
				Expect(c).Should(BeNumerically(">=", count))
			}
			close(finished)
		})

		Ω(runtime.Seconds()).Should(BeNumerically("<", 10), "Under expected performance number")

		b.RecordValue("msgs", float64(count*pushers))
	}, 5)
})
