package swallow

import (
	"github.com/CapillarySoftware/goforward/messaging"
	. "github.com/CapillarySoftware/gomasticate/stomach"

	log "github.com/cihub/seelog"
	"sync"
)

//Swallow data and insert it into the db
func Swallow(swallowChan <-chan *messaging.Food, stomach Stomach, wg *sync.WaitGroup) {
	log.Info("Ready to swallow!")
	for food := range swallowChan {
		fType := food.GetType()
		switch fType {
		case messaging.RFC3164:
			{
				// log.Trace("RFC3164 : ", food)
				err := stomach.IndexDocument(food.GetIndex(), food.GetIndexType(), food.GetId(), food.Rfc3164[0])
				if nil != err {
					log.Error(err)
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
