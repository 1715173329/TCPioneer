package main

import (
	"flag"
)

var ServiceMode bool = true
var ScanIPRange string = ""
var ScanSpeed int = 1
var ScanURL string = ""
var ScanTimeout uint = 0

func main() {
	var flagServiceInstall bool
	var flagServiceRemove bool
	var flagServiceStart bool
	var flagServiceStop bool

	flag.BoolVar(&flagServiceInstall, "install", false, "Install service")
	flag.BoolVar(&flagServiceRemove, "remove", false, "Remove service")
	flag.BoolVar(&flagServiceStart, "start", false, "Start service")
	flag.BoolVar(&flagServiceStop, "stop", false, "Stop service")
	flag.StringVar(&ScanIPRange, "scanip", "", "Scan IP Range")
	flag.IntVar(&ScanSpeed, "scanspeed", 1, "Scan Speed")
	flag.StringVar(&ScanURL, "scanurl", "", "Scan URL")
	flag.UintVar(&ScanTimeout, "scantimeout", 0, "Scan Timeout")
	flag.Parse()

	ServiceMode = InitService(flagServiceInstall, flagServiceRemove, flagServiceStart, flagServiceStop)

	if !ServiceMode {
		StartService()
	}
}
