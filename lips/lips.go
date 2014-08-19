package lips

//Entry point for data

import (
	"github.com/CapillarySoftware/goforward/messaging"
	rep "github.com/CapillarySoftware/goreport"
	log "github.com/cihub/seelog"
	nano "github.com/op/go-nanomsg"
	"sync"
	"time"
)

//Injest data from queue and ship the data off to be swallowed
func OpenWide(chewChan chan *messaging.Food, done chan interface{}, wg *sync.WaitGroup) {
	var (
		msg []byte
		err error
	)

	defer close(chewChan)
	socket, err := nano.NewPullSocket()
	r := rep.NewReporter()

	if nil != err {
		log.Error(err)
	}
	defer socket.Close()
	socket.SetRecvTimeout(1000 * time.Millisecond)
	_, err = socket.Bind("tcp://*:2025")
	if nil != err {
		log.Error(err)
	}

	log.Info("Connected and ready to receive data")
main:
	for {
		select {
		case <-done:
			{
				log.Info("Got done signal")
				break main
			}
		default:
			{
				msg, err = socket.Recv(0)
				if nil != err {
					r.AddStatWIndex("lips", 1, "timeout")
					//we hit timeout
				}
				if nil != msg {
					food := new(messaging.Food)
					err = food.Unmarshal(msg)
					if nil != err {
						log.Error("Invalid message: ", err)
						r.AddStatWIndex("lips", 1, "bad")
						continue
					}
					r.AddStatWIndex("lips", 1, "good")
					chewChan <- food

				}
			}
		}
	}

	log.Info("Closing lips")
	log.Flush()
	wg.Done()

}
