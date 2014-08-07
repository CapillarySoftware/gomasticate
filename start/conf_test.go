package start_test

import (
	. "github.com/CapillarySoftware/gomasticate/start"

	log "github.com/cihub/seelog"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Conf", func() {
	It("Fail to read conf that doesn't exist", func() {
		conf := new(Conf)
		err := conf.InitConf("test.yaml_not_exist")
		Expect(err).ShouldNot(Equal(BeNil()))
	})

	Describe("Read test file and parse out objects by expected type", func() {
		var (
			conf *Conf
			err  error
		)
		BeforeEach(func() {
			conf = new(Conf)
			err = conf.InitConf("test.yaml")
			if nil != err {
				log.Error(err)
			}
			log.Debug("test yaml loaded")
		})

		It("Read All values out of test yaml", func() {
			esHost := conf.EsHost()
			Expect(esHost).Should(Equal("localhost"))
			mouthPort := conf.MouthPort()
			Expect(mouthPort).Should(Equal(2025))
		})

	})

})
