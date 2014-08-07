package mouth

//Entry point for data

//Injest data from queue and ship the data off to be swallowed
func OpenWide(chewChan chan *messaging.Food) {
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

	for {
		msg, err = socket.Recv(0) //blocking
		if nil != err {
			log.Error(err)
		}
		if nil != msg {
			food := new(messaging.Food)
			food.Unmarshal(msg)
			chewChan <- food

		} else {
			log.Warn("Null mesage...? Possible shutdown.")
		}
	}

}
