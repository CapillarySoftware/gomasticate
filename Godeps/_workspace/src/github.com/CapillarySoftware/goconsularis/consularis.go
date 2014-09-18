package goconsularis

import (
	consul "github.com/armon/consul-api"
	log "github.com/cihub/seelog"
	"strconv"
	"time"
)

//Register service and keep it registered.
func RegisterService(name string, port int, ttl int) {
	go registerService(name, port, ttl)
}

func registerCheckTtl(name string, ttl int, agent *consul.Agent) {
	reg := &consul.AgentCheckRegistration{
		Name: name,
	}
	reg.TTL = strconv.Itoa(ttl) + "s"

	if err := agent.CheckRegister(reg); err != nil {
		log.Error("Failed to register check: ", err)
	}
}

//continue to register a service preferably ran in a go routine.
func registerService(name string, port int, ttl int) {

	reportInterval := make(chan bool, 1)
	go func() {
		for {
			time.Sleep(time.Duration(ttl) / 2 * time.Second)
			reportInterval <- true
		}
	}()

	client, err := consul.NewClient(consul.DefaultConfig())
	if nil != err {
		log.Error("Failed to get consul client")
	}
	agent := client.Agent()

	serviceRegister(name, port, ttl, agent)
	for {
		select {
		case <-reportInterval: //report registration
			{
				servicePassing(name, agent)

			}
		}
	}

}

func servicePassing(name string, agent *consul.Agent) {
	agent.PassTTL("service:"+name, "Service up and ready!")
}

func serviceRegister(name string, port int, ttl int, agent *consul.Agent) {
	reg := &consul.AgentServiceRegistration{
		Name: name,
		Port: port,
		Check: &consul.AgentServiceCheck{
			TTL: strconv.Itoa(ttl) + "s",
		},
	}
	if err := agent.ServiceRegister(reg); err != nil {
		log.Error("err: ", err)
	}
}
