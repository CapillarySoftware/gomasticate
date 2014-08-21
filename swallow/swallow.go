package swallow

import (
	"github.com/CapillarySoftware/goforward/messaging"
	. "github.com/CapillarySoftware/gomasticate/stomach"
	rep "github.com/CapillarySoftware/goreport"
	log "github.com/cihub/seelog"
	"sync"
)

//Swallow data and insert it into the db
func Swallow(swallowChan <-chan *messaging.Food, stomach Stomach, wg *sync.WaitGroup) {
	log.Info("Ready to swallow!")
	r := rep.NewReporter()
	r.RegisterStatWIndex("swallow", "RFC3164Bad")
	r.RegisterStatWIndex("swallow", "RFC3164Good")
	r.RegisterStatWIndex("swallow", "RFC5424Bad")
	r.RegisterStatWIndex("swallow", "RFC5424Good")
	r.RegisterStatWIndex("swallow", "JSONBad")
	r.RegisterStatWIndex("swallow", "JSONGood")

	for food := range swallowChan {
		fType := food.GetType()
		switch fType {
		case messaging.RFC3164:
			{
				for _, v := range food.Rfc3164 {
					err := stomach.IndexDocument(food.GetIndex(), food.GetIndexType(), v)
					if nil != err {
						log.Error(err)
						r.AddStatWIndex("Swallow", 1, "RFC3164Bad")
					} else {
						r.AddStatWIndex("Swallow", 1, "RFC3164Good")
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
				// log.Trace("JSON :", food)
				// stomach.IndexDocument(food)
			}

		default:
			{
				// log.Error("Invalid message : ", food)
			}
		}
	}
	log.Info("Done Swallowing")
	log.Flush()
	wg.Done()

}
