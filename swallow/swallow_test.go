package swallow_test

import (
	"github.com/CapillarySoftware/goforward/messaging"
	. "github.com/CapillarySoftware/gomasticate/stomach"
	. "github.com/CapillarySoftware/gomasticate/swallow"
	log "github.com/cihub/seelog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"sync"
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
			wg          sync.WaitGroup
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
			wg.Add(1)
			go Swallow(swallowChan, db, &wg)
			swallowChan <- food
			close(swallowChan)
			wg.Wait()
			count := 0
			for {
				db.Lock()
				if db.Doc != "" {

					log.Info(db.Index)
					db.Unlock()
					break
				}
				db.Unlock()
				time.Sleep(100 * time.Nanosecond)
				count++
			}

			db.Lock()
			Expect(db.Doc).Should(Equal(food.Rfc3164[0].String()))
			db.Unlock()

			// swallow.Swallow()
		})
	})

})

type DB struct {
	sync.Mutex
	Doc       string
	Index     string
	IndexType string
	Id        string
}

func (this *DB) IndexDocument(index string, indexType string, doc Document) (err error) {
	this.Lock()
	this.Index = index
	this.IndexType = indexType
	this.Doc = doc.String()
	this.Unlock()
	return
}
