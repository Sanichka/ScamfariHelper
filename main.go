package main

import (
	"ScamfariHelper/modules/configModule"
	"ScamfariHelper/modules/logModule"
	"ScamfariHelper/modules/searchModule"
	"fmt"
	"sync"
)

var WG = new(sync.WaitGroup)

func main() {
	logModule.Logfile, logModule.Logger = logModule.InitLogs()
	logModule.Logger.Println("Log writing initiated...")
	errcfcg := configModule.LoadConfig("cfg")
	if errcfcg != true {
		logModule.Logger.Println("Error during loading cfg.env file. Please check logs file!")
		fmt.Println("Error during loading cfg.env file. Please check logs file!")
	}
	//searchModule.StartGrabber(searchModule.KeysChan)
	//keyboardCapture.Start()
	WG.Add(1)
	searchModule.ReadKeyboardInput(WG)
	defer WG.Wait()
}
