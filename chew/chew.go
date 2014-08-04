package chew

import (
	log "github.com/cihub/seelog"
	nano "github.com/op/go-nanomsg"
)

func Chew(count int, finished chan int) {
	var (
		msg []byte
		err error
	)
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
	tot := 0

	for {
		msg, err = socket.Recv(0) //blocking
		if nil != err {
			log.Error(err)
		} else {
			// log.Trace(msg)
			if nil != msg {
				tot++
			}
		}
		if tot >= count {
			break
		}
	}
	finished <- tot

}
