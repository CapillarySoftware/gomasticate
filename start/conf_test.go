package start_test

import (
	. "github.com/CapillarySoftware/gomasticate/start"

	log "github.com/cihub/seelog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conf", func() {
	It("Fail to read conf that doesn't exist", func() {
		conf, err := GetConf("test.yaml_not_exist")
		Expect(err).ShouldNot(Equal(BeNil()))
		Expect(conf).Should(BeNil())
	})

	It("Fail read invalid file", func() {
		_, err := GetConf("conf.go")
		Expect(err).ShouldNot(Equal(BeNil()))

	})

	Describe("Read file and parse out objects by expected type", func() {
		var (
			conf map[interface{}]interface{}
			err  error
		)
		BeforeEach(func() {
			conf, err = GetConf("test.yaml")
			if nil != err {
				log.Error(err)
			}
		})

		It("Read chew map out of yaml ", func() {
			chew := conf["chew"].(map[interface{}]interface{})
			port := chew["port"].(int)
			host := chew["host"].(string)
			Expect(chew).ShouldNot(BeNil())
			Expect(port).Should(Equal(2025))
			Expect(host).Should(Equal("localhost"))
		})
		It("Read es map out of yaml ", func() {
			es := conf["es"].(map[interface{}]interface{})
			port := es["port"].(int)
			host := es["host"].(string)
			Expect(es).ShouldNot(BeNil())
			Expect(port).Should(Equal(9054))
			Expect(host).Should(Equal("localhost"))
		})

	})

})
