package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"./header"
	"github.com/chai2010/winsvc"
)

func StartService() {
	runtime.GOMAXPROCS(1)
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if ghostcp.LogLevel > 0 {
		var logFilename string = "tcpioneer.log"
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

	Windir := os.Getenv("WINDIR")
	err = ghostcp.LoadHosts(Windir + "\\System32\\drivers\\etc\\hosts")
	if err != nil && !ServiceMode {
		log.Println(err)
		return
	}

	if ghostcp.LogLevel == 0 && !ServiceMode {
		ghostcp.LogLevel = 1
	}

	ghostcp.TCPDaemon(":443", false)
	ghostcp.TCPDaemon(":80", false)
	ghostcp.UDPDaemon(443, false)
	ghostcp.TCPRecv(443, false)

	if ghostcp.Forward {
		ghostcp.TCPDaemon(":443", true)
		ghostcp.TCPDaemon(":80", true)
		ghostcp.UDPDaemon(443, true)
		ghostcp.TCPRecv(443, true)
	}

	if ghostcp.DNS == "" {
		ghostcp.DNSRecvDaemon()
	} else {
		ghostcp.TCPDaemon(ghostcp.DNS, false)
		ghostcp.DNSDaemon()
	}

	if ScanIPRange != "" {
		ghostcp.ScanURL = ScanURL
		ghostcp.ScanTimeout = ScanTimeout

		go ghostcp.Scan(ScanIPRange, ScanSpeed)
	}

	fmt.Println("Service Start")
	ghostcp.Wait()
}

func StopService() {
	arg := []string{"/flushdns"}
	cmd := exec.Command("ipconfig", arg...)
	d, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(d), err)
	}

	os.Exit(0)
}

func InitService(install, remove, start, stop bool) bool {
	serviceName := "Ghostcp"

	appPath, err := winsvc.GetAppPath()
	if err != nil {
		log.Fatal(err)
	}

	// install service
	if install {
		if err := winsvc.InstallService(appPath, serviceName, ""); err != nil {
			log.Fatalf("installService(%s, %s): %v\n", serviceName, "", err)
		}
		log.Printf("Done\n")
		return true
	}

	// remove service
	if remove {
		if err := winsvc.RemoveService(serviceName); err != nil {
			log.Fatalln("removeService:", err)
		}
		log.Printf("Done\n")
		return true
	}

	// start service
	if start {
		if err := winsvc.StartService(serviceName); err != nil {
			log.Fatalln("startService:", err)
		}
		log.Printf("Done\n")
		return true
	}

	// stop service
	if stop {
		if err := winsvc.StopService(serviceName); err != nil {
			log.Fatalln("stopService:", err)
		}
		log.Printf("Done\n")
		return true
	}

	// run as service
	if !winsvc.IsAnInteractiveSession() {
		log.Println("main:", "runService")

		if err := os.Chdir(filepath.Dir(appPath)); err != nil {
			log.Fatal(err)
		}

		if err := winsvc.RunAsService(serviceName, StartService, StopService, false); err != nil {
			log.Fatalf("svc.Run: %v\n", err)
		}
		return true
	}

	return false
}
