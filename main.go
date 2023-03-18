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
	searchModule.LoadFile()
	fmt.Printf(`
Scamfari helper initialized. Make sure you filled your wallets.txt file with wallets
and old.wav, new.wav audiofile is present. You can fold/minimize console window.
Once you copy wallet via Ctrl+C you will hear audio new.wav if unique wallet(not present in wallets.txt)
In case it's already in file old.wav will be played. 
For support write in official telegram group: https://t.me/scamfari_public
`)
	fmt.Println("Wallets in file = ", len(searchModule.WalletsList))
	WG.Add(1)
	searchModule.ReadKeyboardInput(WG)
	defer WG.Wait()
}
