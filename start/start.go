package start

//Start manages the main run loop of the application
import (
	"github.com/CapillarySoftware/goforward/messaging"
	"github.com/CapillarySoftware/gomasticate/chew"
	"github.com/CapillarySoftware/gomasticate/swallow"
	log "github.com/cihub/seelog"
	// yaml "gopkg.in/yaml.v1"
	"os"
	"os/signal"
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
	log.Info("Starting gomasticate")
	conf, err := GetConf("conf.yaml")
	if nil != err {
		log.Error(err)
		return
	}
	log.Info(conf)
	swallowChan := make(chan *messaging.Food, 1000)
	go chew.Chew(swallowChan)
	go swallow.Swallow(swallowChan)
	c := make(chan os.Signal, 1)
	s := make(chan int, 1)
	signal.Notify(c)
	go Death(c, s)
	death := <-s //time for shutdown
	log.Debug("Death return code: ", death)
	log.Info("Exiting")
}
