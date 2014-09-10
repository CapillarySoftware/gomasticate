package lips_test

import (
	"github.com/CapillarySoftware/gomasticate/lips"
	log "github.com/cihub/seelog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strconv"
	// "time"
	nano "github.com/op/go-nanomsg"
)

var _ = Describe("Push and Pull", func() {
	Describe("Positive tests", func() {

	})
	Describe("Benchmarks", func() {
		Measure("25 pushers single puller performance", func(b Benchmarker) {
			port := 9331
			count := 20000
			pushers := 25
			runtime := b.Time("runtime", func() {
				finished := make(chan int, 10)

				go Pull(count*pushers, finished, port)
				for i := 0; i < pushers; i++ {
					log.Info("Starting pusher")
					go lips.Pusher(count, finished, port)
				}
				log.Info("Waiting for messages")
				for i := 0; i < pushers+1; i++ {
					log.Info("Checking finished pushers and puller")
					c := <-finished
					Expect(c).Should(BeNumerically(">=", count))
				}
				close(finished)
			})

			Ω(runtime.Seconds()).Should(BeNumerically("<", 10), "Under expected performance number")

			b.RecordValue("msgs", float64(count*pushers))
		}, 5)
	})

})

//Simple nano puller
func Pull(count int, finished chan int, port int) {
	var (
		msg []byte
		err error
	)
	socket, err := nano.NewPullSocket()

	if nil != err {
		log.Error(err)
	}
	defer socket.Close()
	sport := strconv.Itoa(port)
	_, err = socket.Bind("tcp://*:" + sport)
	if nil != err {
		log.Error(err)
	}
	log.Info("Connected and ready to receive data")
	tot := 0

	for {
		msg, err = socket.Recv(0) //blocking
		if nil != err {
			log.Error(err)
		} else {
			// log.Trace(msg)
			if nil != msg {
				tot++
			}
		}
		if tot >= count {
			break
		}
	}
	finished <- tot

}
