package start

import (
	log "github.com/cihub/seelog"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

type Conf struct {
	c     map[interface{}]interface{}
	mouth map[interface{}]interface{}
	es    map[interface{}]interface{}
}

//initialize our config file
func (this *Conf) InitConf(path string) (err error) {
	cont, err := ioutil.ReadFile(path)
	if nil != err {
		log.Error(err)
		return
	}
	this.c = make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(cont), &this.c)
	if nil != err {
		log.Error(err)
		return
	}
	this.mouth = this.c["mouth"].(map[interface{}]interface{})
	this.es = this.c["es"].(map[interface{}]interface{})
	return
}

//Hostname or ip of elasticsearch server
func (this *Conf) EsHost() (host string) {
	host = this.es["host"].(string)
	return
}

// Port used to feed the mouth
func (this *Conf) MouthPort() (port int) {
	port = this.mouth["port"].(int)
	return
}
