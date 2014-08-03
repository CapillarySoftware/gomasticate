package main

import (
	"github.com/CapillarySoftware/gomasticate/start"
	log "github.com/cihub/seelog"
)

func main() {
	defer log.Flush()
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")

	if err != nil {
		log.Warn("Failed to load config", err)
	}

	log.ReplaceLogger(logger)
	start.Run()
}
