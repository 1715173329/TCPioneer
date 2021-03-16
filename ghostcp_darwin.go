package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"./header"
)

func StartService() {
	runtime.GOMAXPROCS(1)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if ghostcp.LogLevel > 0 {
		var logFilename string = "ghostcp.log"
		logFile, err := os.OpenFile(logFilename, os.O_RDWR|os.O_CREATE, 0777)
		if err != nil {
			log.Println(err)
			return
		}
		defer logFile.Close()

		ghostcp.Logger = log.New(logFile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)
	}

	err := ghostcp.LoadConfig()
	if err != nil {
		if ghostcp.LogLevel > 0 || !ServiceMode {
			log.Println(err)
		}
		return
	}

	err = ghostcp.LoadHosts("/etc/hosts")
	if err != nil && !ServiceMode {
		log.Println(err)
		return
	}

	if ghostcp.LogLevel == 0 && !ServiceMode {
		ghostcp.LogLevel = 1
	}

	ghostcp.Monitor("")

	fmt.Println("Service Start")
	ghostcp.Wait()
}

func StopService() {
	arg := []string{"-HUP mDNSResponder"}
	cmd := exec.Command("killall", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(d), err)
	}

	os.Exit(0)
}

func InitService(install, remove, start, stop bool) bool {
	// install service
	if install {
		return true
	}

	// remove service
	if remove {
		return true
	}

	// start service
	if start {
		return true
	}

	// stop service
	if stop {
		return true
	}

	return false
}
