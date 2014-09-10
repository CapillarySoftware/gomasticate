package start

//Start manages the main run loop of the application
import (
	"github.com/CapillarySoftware/goforward/messaging"
	"github.com/CapillarySoftware/gomasticate/chew"
	"github.com/CapillarySoftware/gomasticate/lips"
	"github.com/CapillarySoftware/gomasticate/swallow"

	log "github.com/cihub/seelog"
	"os"
	"os/signal"
	"sync"
)

//Manage death of application by signal
func Death(c <-chan os.Signal, death chan int) {
	for sig := range c {
		switch sig.String() {
		case "terminated":
			{
				death <- 1
			}
		case "interrupt":
			{
				death <- 2
			}
		default:
			{
				death <- 3
			}
		}

	}
}

//Run the app.
func Run() (err error) {
	var wg sync.WaitGroup
	log.Info("Starting gomasticate")
	conf, err := NewConf("conf.yaml")
	if nil != err {
		log.Error(err)
		return
	}
	log.Info(conf)
	chewChan := make(chan *messaging.Food, 2000)
	swallowChan := make(chan *messaging.Food, 4000)

	done := make(chan interface{})

	wg.Add(2)
	go lips.OpenWide(chewChan, done, &wg, conf.LipsPort())
	go chew.Chew(chewChan, swallowChan, &wg)

	sw := swallow.NewSwallow(conf.EsHost(), swallowChan, 10)

	//handle signals
	c := make(chan os.Signal, 1)
	s := make(chan int, 1)
	signal.Notify(c)
	go Death(c, s)
	death := <-s //time for shutdown
	log.Debug("Death return code: ", death)
	close(done)
	sw.Close()
	log.Info("Waiting for goroutines to finish...")
	wg.Wait()
	log.Info("Exiting")
	return
}
