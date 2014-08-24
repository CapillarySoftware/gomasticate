package swallow

import (
	"github.com/CapillarySoftware/goforward/messaging"
	. "github.com/CapillarySoftware/gomasticate/stomach"
	log "github.com/cihub/seelog"
	gi "github.com/onsi/ginkgo"
	gom "github.com/onsi/gomega"
	"sync"
	"time"
)

var _ = gi.Describe("Swallow", func() {

	gi.Describe("NewSwallowers", func() {
		gi.It("Test swallower new method, with close", func() {
			foodChan := make(chan *messaging.Food)
			sw := NewSwallow("test", foodChan, 2)
			close(foodChan)
			sw.Close()
		})
	})
	gi.Describe("Positive tests", func() {
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
		gi.BeforeEach(func() {
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
		gi.It("Test RFC3164", func() {
			// log.Info(food)
			db := new(DB)
			wg.Add(1)
			go swallow(swallowChan, db, &wg)
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
			gom.Expect(db.Doc).Should(gom.Equal(food.Rfc3164[0].String()))
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
func (this *DB) Close() {

}
