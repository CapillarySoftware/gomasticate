package start

//Start manages the main run loop of the application
import (
	"github.com/CapillarySoftware/goforward/messaging"
	"github.com/CapillarySoftware/gomasticate/chew"
	_es "github.com/CapillarySoftware/gomasticate/elasticsearch"
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
func Run() {
	var wg sync.WaitGroup
	log.Info("Starting gomasticate")
	conf := new(Conf)
	err := conf.InitConf("conf.yaml")
	if nil != err {
		log.Error(err)
		return
	}
	log.Info(conf)
	es := new(_es.Elasticsearch)
	es.Connect("localhost")
	chewChan := make(chan *messaging.Food, 1000)
	swallowChan := make(chan *messaging.Food, 1000)
	done := make(chan interface{})

	wg.Add(8)
	go lips.OpenWide(chewChan, done, &wg)
	go chew.Chew(chewChan, swallowChan, &wg)
	go swallow.Swallow(swallowChan, es, &wg)
	go swallow.Swallow(swallowChan, es, &wg)
	go swallow.Swallow(swallowChan, es, &wg)
	go swallow.Swallow(swallowChan, es, &wg)
	go swallow.Swallow(swallowChan, es, &wg)
	go swallow.Swallow(swallowChan, es, &wg)

	//handle signals
	c := make(chan os.Signal, 1)
	s := make(chan int, 1)
	signal.Notify(c)
	go Death(c, s)
	death := <-s //time for shutdown
	log.Debug("Death return code: ", death)
	close(done)
	log.Info("Waiting for goroutines to finish...")
	wg.Wait()
	log.Info("Exiting")
}
