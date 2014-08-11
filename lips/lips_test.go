package lips_test

import (
	"github.com/CapillarySoftware/goforward/messaging"
	"github.com/CapillarySoftware/gomasticate/lips"
	log "github.com/cihub/seelog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sync"
	"time"
)

var _ = Describe("Lips", func() {

	var (
		timestamp int64
		hostname  string
		tag       string
		content   string
		priority  int32
		facility  int32
		severity  int32
		food      *messaging.Food
		chewChan  chan *messaging.Food
		done      chan interface{}
	)

	BeforeEach(func() {

		timestamp = int64(time.Now().Unix())
		hostname = "hostname"
		tag = "tag"
		content = "content"
		priority = 1
		facility = 7
		severity = 2
		fType := messaging.RFC3164
		food = new(messaging.Food)
		food.Type = &fType

		msg := new(messaging.Rfc3164)
		msg.Timestamp = &timestamp
		msg.Hostname = &hostname
		msg.Tag = &tag
		msg.Content = &content
		msg.Priority = &priority
		msg.Severity = &severity

		food.Rfc3164 = append(food.Rfc3164, msg)
		chewChan = make(chan *messaging.Food, 100) //blocking
		done = make(chan interface{})

	})

	It("Test simple message through the lips", func() {
		msgCount := 10
		var wg sync.WaitGroup
		tot := make(chan int)
		wg.Add(1)
		go lips.PusherProto(msgCount, tot, food)
		go lips.OpenWide(chewChan, done, &wg)
		log.Info("Waiting for sends to be finished")
		count := <-tot //sync until it's been sent
		Expect(count).Should(Equal(msgCount))

		rcvCount := 0
		log.Info("Receiving msgs")
		for _ = range chewChan {
			rcvCount++
			if rcvCount >= 10 {
				break
			}
		}
		Expect(len(chewChan)).Should(Equal(0))
		close(done)
		log.Info("Waiting for shutdown")
		wg.Wait()
	})
})
