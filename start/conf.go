package start

import (
	log "github.com/cihub/seelog"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

type Conf struct {
	c    map[interface{}]interface{}
	lips map[interface{}]interface{}
	es   map[interface{}]interface{}
}

//initialize our config file
func NewConf(path string) (conf *Conf, err error) {
	conf = new(Conf)
	cont, err := ioutil.ReadFile(path)
	if nil != err {
		log.Error(err)
		return
	}
	conf.c = make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(cont), &conf.c)
	if nil != err {
		log.Error(err)
		return
	}
	conf.lips = conf.c["lips"].(map[interface{}]interface{})
	conf.es = conf.c["es"].(map[interface{}]interface{})

	return
}

//Hostname or ip of elasticsearch server
func (this *Conf) EsHost() (host string) {
	host = this.es["host"].(string)
	return
}

// Port used to feed the mouth
func (this *Conf) LipsPort() (port int) {
	port = this.lips["port"].(int)
	return
}
