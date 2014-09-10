#!/bin/bash
#Don't change this to $home it will fail in ansible
go get github.com/tools/godep
go install github.com/tools/godep
godep restore
godep go build github.com/CapillarySoftware/gomasticate
