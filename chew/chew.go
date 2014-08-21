package chew

//Chew is where we transform the data.

import (
	// "fmt"
	"github.com/CapillarySoftware/goforward/messaging"
	rep "github.com/CapillarySoftware/goreport"
	log "github.com/cihub/seelog"
	"sync"
	"time"
)

func Chew(chewChan <-chan *messaging.Food, swallowChan chan *messaging.Food, wg *sync.WaitGroup) {
	log.Info("Let the chewing begin!")
	defer close(swallowChan)
	r := rep.NewReporter()
	r.RegisterStatWIndex("chew", "good")
	for msg := range chewChan {
		if nil != msg {
			//parsing work here probably change what our message type looks like when swallowed

			date := time.Unix(0, msg.GetTimeNano()).UTC()
			fmtDate := date.Format("2006-01-02")
			indexType := "all"
			customerId := "id" //should exist eventually
			index := "documents-" + fmtDate + "-" + customerId
			msg.Index = &index
			msg.IndexType = &indexType
			r.AddStatWIndex("chew", 1, "good")
			swallowChan <- msg
		}
	}
	log.Info("Done chewing")
	log.Flush()
	wg.Done()

}
