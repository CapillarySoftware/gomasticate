package lips

import (
	"github.com/CapillarySoftware/goforward/messaging"
	log "github.com/cihub/seelog"
	nano "github.com/op/go-nanomsg"
	"time"
)

//Simple pusher for testing
func Pusher(count int, finished chan int) {
	socket, err := nano.NewPushSocket()
	if nil != err {
		log.Error(err)
	}
	defer socket.Close()
	_, err = socket.Connect("tcp://localhost:2025")
	if nil != err {
		log.Error(err)
		return
	}
	log.Info("Connected and ready to send data")
	tot := 0
	bytes := []byte{'h', 'e', 'l', 'l', '\xc3', '\xb8', 'x', 'f', 'c', 'x', 'f'}
	for {
		_, err := socket.Send(bytes, 0) //blocking
		if nil != err {
			log.Error(err)
		} else {
			tot++
		}
		if tot >= count {
			break
		}
	}
	finished <- tot
}

//Simple pusher for testing
func PusherProto(count int, finished chan int, msg *messaging.Food) {
	log.Info("Starting pusher")
	socket, err := nano.NewPushSocket()
	if nil != err {
		log.Error(err)
	}
	defer socket.Close()
	socket.SetSendTimeout(500 * time.Millisecond)
	_, err = socket.Connect("tcp://localhost:2025")
	if nil != err {
		log.Error(err)
		return
	}
	log.Info("Connected and ready to send data")
	tot := 0
	for {
		bytes, _ := msg.Marshal()
		_, err := socket.Send(bytes, 0) //blocking
		if nil != err {
			log.Error(err)
			continue
		} else {
			tot++
		}
		if tot >= count {
			break
		}
	}
	log.Info("Finished sending data exiting")
	finished <- tot
}
