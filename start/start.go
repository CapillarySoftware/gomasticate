package start

//Start manages the main run loop of the application
import (
	// "flag"
	log "github.com/cihub/seelog"
	"os"
	"os/signal"
	// "strconv"
	// "strings"
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
	// flag.Parse()
	c := make(chan os.Signal, 1)
	s := make(chan int, 1)
	signal.Notify(c)
	go Death(c, s)
	death := <-s //time for shutdown
	log.Debug("Death return code: ", death)
	log.Info("Exiting")
}
