package swallow

import (
	"bytes"
	"encoding/json"
	"github.com/CapillarySoftware/goforward/messaging"
	_es "github.com/CapillarySoftware/gomasticate/elasticsearch"
	. "github.com/CapillarySoftware/gomasticate/stomach"
	rep "github.com/CapillarySoftware/goreport"
	log "github.com/cihub/seelog"
	"sync"
)

type Swallow struct {
	wg *sync.WaitGroup
}

type Json struct {
	Msg interface{}
	id  string
}

func (this *Json) GetId() string {
	return this.id
}

func (this *Json) String() string {
	return this.id
}

//New Swallowers
func NewSwallow(url string, swallowChan <-chan *messaging.Food, concurrent int) *Swallow {

	wg := sync.WaitGroup{}
	sw := Swallow{wg: &wg}

	wg.Add(concurrent)
	for i := 0; i < concurrent; i++ {
		s := _es.Elasticsearch{}
		s.Connect(url)
		go swallow(swallowChan, &s, &wg)
	}
	return &sw
}

//I'm full!!!!@
func (this *Swallow) Close() {
	this.wg.Wait()
}

//Swallow data and insert it into the db
func swallow(swallowChan <-chan *messaging.Food, stomach Stomach, wg *sync.WaitGroup) {
	log.Info("Ready to swallow!")
	r := rep.NewReporter()
	r.RegisterStatWIndex("swallow", "RFC3164Bad")
	r.RegisterStatWIndex("swallow", "RFC3164Good")
	r.RegisterStatWIndex("swallow", "RFC5424Bad")
	r.RegisterStatWIndex("swallow", "RFC5424Good")
	r.RegisterStatWIndex("swallow", "JSONBad")
	r.RegisterStatWIndex("swallow", "JSONGood")
	r.RegisterStatWIndex("swallow", "Invalid")

	for food := range swallowChan {
		fType := food.GetType()
		switch fType {
		case messaging.RFC3164:
			{
				for _, v := range food.Rfc3164 {
					err := stomach.IndexDocument(food.GetIndex(), food.GetIndexType(), v)
					if nil != err {
						log.Error(err)
						r.AddStatWIndex("swallow", 1, "RFC3164Bad")
					} else {
						r.AddStatWIndex("swallow", 1, "RFC3164Good")
					}
				}
			}

		case messaging.RFC5424:
			{
				// log.Trace("RFC5424 :", food)
				// stomach.IndexDocument(food)
			}

		case messaging.JSON:
			{
				var buf *bytes.Buffer
				var msg interface{}
				var err error
				for _, v := range food.Json {
					buf = bytes.NewBuffer(v.GetJson())
					err = json.NewDecoder(buf).Decode(&msg)
					if nil != err {
						log.Error("Failed to decode json: ", err)
						continue
					}
					err = stomach.IndexDocumentDynamic(food.GetIndex(), food.GetIndexType(), &msg, v.GetId())
					if nil != err {
						log.Error(err)
						r.AddStatWIndex("swallow", 1, "JSONBad")
					} else {
						r.AddStatWIndex("swallow", 1, "JSONGood")
					}
				}
			}

		default:
			{
				r.AddStatWIndex("swallow", 1, "Invalid")
			}
		}
	}
	log.Info("Done Swallowing")
	log.Flush()
	wg.Done()

}
