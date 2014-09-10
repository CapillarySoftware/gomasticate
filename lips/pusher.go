package lips

import (
	"github.com/CapillarySoftware/goforward/messaging"
	log "github.com/cihub/seelog"
	nano "github.com/op/go-nanomsg"
	"strconv"
	"time"
)

//Simple pusher for testing
func Pusher(count int, finished chan int, port int) {
	socket, err := nano.NewPushSocket()
	if nil != err {
		log.Error(err)
	}
	defer socket.Close()
	sport := strconv.Itoa(port)
	_, err = socket.Connect("tcp://localhost:" + sport)
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
func PusherProto(count int, finished chan int, msg *messaging.Food, port int) {
	log.Info("Starting pusher")
	socket, err := nano.NewPushSocket()
	if nil != err {
		log.Error(err)
	}
	defer socket.Close()
	socket.SetSendTimeout(500 * time.Millisecond)
	sport := strconv.Itoa(port)
	_, err = socket.Connect("tcp://localhost:" + sport)
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
