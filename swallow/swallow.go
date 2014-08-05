package swallow

import (
	"github.com/CapillarySoftware/goforward/messaging"
	log "github.com/cihub/seelog"
)

func Swallow(swallowChan <-chan *messaging.Food) {
	log.Info("Ready to swallow!")
	for food := range swallowChan {
		fType := food.GetType()
		switch fType {
		case messaging.RFC3164:
			{
				log.Trace("RFC3164 : ", food)
			}

		case messaging.RFC5424:
			{
				log.Trace("RFC5424 :", food)
			}

		case messaging.JSON:
			{
				log.Trace("JSON :", food)
			}

		default:
			{
				log.Error("Invalid message : ", food)
			}
		}
	}
}
