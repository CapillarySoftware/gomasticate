package lips

//Entry point for data

import (
	"github.com/CapillarySoftware/goforward/messaging"
	log "github.com/cihub/seelog"
	nano "github.com/op/go-nanomsg"
	"sync"
)

//Injest data from queue and ship the data off to be swallowed
func OpenWide(chewChan chan *messaging.Food, done chan interface{}, wg *sync.WaitGroup) {
	var (
		msg []byte
		err error
	)
	defer close(chewChan)
	socket, err := nano.NewPullSocket()

	if nil != err {
		log.Error(err)
	}
	defer socket.Close()
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
					log.Error(err)
					break main
				}
				if nil != msg {
					food := new(messaging.Food)
					err = food.Unmarshal(msg)
					if nil != err {
						log.Error("Invalid message: ", err)
						continue
					}
					chewChan <- food

				} else {
					log.Warn("Null mesage...? Possible shutdown.")
				}
			}
		}
	}

	log.Info("Closing lips")
	log.Flush()
	wg.Done()

}
