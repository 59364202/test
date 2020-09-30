package main

import (
	"haii.or.th/api/server/model/backgroundjob/bgrunner"
)

const envPrefix = "TW30"

var BuildVersion = ""

var runner = bgrunner.NewRunner(envPrefix, BuildVersion)

func main() {
	runner.RunExit()
}
