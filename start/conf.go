package start

import (
	log "github.com/cihub/seelog"
	"gopkg.in/yaml.v1"
	"io/ioutil"
)

func GetConf(path string) (conf map[interface{}]interface{}, err error) {
	cont, err := ioutil.ReadFile(path)
	if nil != err {
		log.Error(err)
		return
	}
	conf = make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(cont), &conf)
	if err != nil {
		log.Error("error: %v", err)
	}
	return

}
