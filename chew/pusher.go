package chew

import (
	log "github.com/cihub/seelog"
	nano "github.com/op/go-nanomsg"
)

//Simple pusher for testing
func Pusher(count int, finished chan int) {
	socket, err := nano.NewPushSocket()
	if nil != err {
		log.Error(err)
	}
	defer socket.Close()
	_, err = socket.Connect("tls://localhost:2025")
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
