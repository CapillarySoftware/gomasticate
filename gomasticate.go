package main

import (
	"github.com/CapillarySoftware/gomasticate/start"
	rep "github.com/CapillarySoftware/goreport"
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	rep.ReporterConfig("ipc:///temp/testSender.ipc", 0)
	r := rep.NewReporter()
	defer r.Close()
	if err != nil {
		log.Warn("Failed to load config", err)
	}

	log.ReplaceLogger(logger)
	start.Run()
}
