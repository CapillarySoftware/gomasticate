package chew

//Chew is where we transform the data.

import (
	"github.com/CapillarySoftware/goforward/messaging"
	log "github.com/cihub/seelog"
	"sync"
)

func Chew(chewChan <-chan *messaging.Food, swallowChan chan *messaging.Food, wg *sync.WaitGroup) {
	log.Info("Let the chewing begin!")
	defer close(swallowChan)
	for msg := range chewChan {
		if nil != msg {
			//parsing work here probably change what our message type looks like when swallowed
			swallowChan <- msg
		}
	}
	log.Info("Done chewing")
	log.Flush()
	wg.Done()

}
