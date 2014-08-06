package swallow_test

import (
	"github.com/CapillarySoftware/goforward/messaging"
	. "github.com/CapillarySoftware/gomasticate/stomach"
	. "github.com/CapillarySoftware/gomasticate/swallow"
	log "github.com/cihub/seelog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Swallow", func() {

	Describe("Positive tests", func() {
		var (
			timestamp   int64
			hostname    string
			tag         string
			content     string
			priority    int32
			facility    int32
			severity    int32
			food        *messaging.Food
			swallowChan chan *messaging.Food
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
			swallowChan = make(chan *messaging.Food, 1) //blocking

		})
		It("Test RFC3164", func() {
			// log.Info(food)
			db := new(DB)
			go Swallow(swallowChan, db)
			swallowChan <- food
			close(swallowChan)
			count := 0
			for {
				if db.Index != "" {
					log.Info(db.Index)
					break
				}
				time.Sleep(100 * time.Nanosecond)
				count++
			}
			Expect(db.Index).Should(Equal(food.String()))

			// swallow.Swallow()
		})
	})

})

type DB struct {
	Index string
}

func (this *DB) IndexDocument(doc Document) (err error) {
	this.Index = doc.String()
	return
}
