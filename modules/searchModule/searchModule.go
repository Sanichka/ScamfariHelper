package searchModule

import (
	"ScamfariHelper/modules/audio"
	"ScamfariHelper/modules/configModule"
	"ScamfariHelper/modules/keyboardCapture"
	. "ScamfariHelper/modules/logModule"
	"bufio"
	"golang.design/x/clipboard"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

var WalletsList []string
var ClipboardContent string

func LoadFile() bool {
	file, err := os.Open(configModule.WALLETS_FILE)
	var result []string
	if err != nil {
		Logger.Printf("Error opening file %v. Please check if it's present!", configModule.WALLETS_FILE)
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		for _, v := range strings.SplitAfter(scanner.Text(), "\"") {
			for _, vv := range strings.SplitAfter(v, " ") {
				cleanstr := strings.TrimSpace(vv)
				//fmt.Println(cleanstr)
				if len(cleanstr) >= 30 {
					result = append(result, cleanstr)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		Logger.Printf("Error scanning file %v. Error: ", configModule.WALLETS_FILE, err)
		log.Fatal(err)
	}
	if len(result) <= 0 {
		Logger.Println("WARNING: Wallets file is empty!")
		return false
	}

	WalletsList = result
	return true
}

func IsDuplicate(wallet string) bool {
	if len(wallet) <= 30 {
		return false
	}
	for _, v := range WalletsList {
		if v == wallet {
			return true
		}
	}
	return false
}

// Ctrl = 162; c = 67
func ReadKeyboardInput(wg *sync.WaitGroup) {
	run := true
	//Clipboard = clipboard.Init()
	out := make(chan string, 100)
	currentKey := ""
	lastKey := ""
	err := clipboard.Init()
	if err != nil {
		Logger.Printf("Error in clipboard init: %v", err)
	}
	go keyboardCapture.Start(out)
	go func() {
		for {
			currentKey = <-out
			switch currentKey {
			case "67":
				//fmt.Printf("Last key:%v Current key: %v\n", lastKey, currentKey)
				if lastKey == "162" {
					//fmt.Printf("CTRL+C pressed: Last key:%v Current key: %v\n", lastKey, currentKey)
					time.Sleep(20 * time.Millisecond)
					ClipboardContent = readClipboard()
					//fmt.Println("ClipboardContent:", ClipboardContent)
					// Validate string length in order to not annoy on each Ctrl+c press
					if len(ClipboardContent) >= 30 {
						switch IsDuplicate(ClipboardContent) {
						case true:
							//fmt.Println("Streaming found audio file")
							go audio.StreamAudio(configModule.FOUND_AUDIO_FILE)
						case false:
							//fmt.Println("Streaming NOT found audio file")
							go audio.StreamAudio(configModule.NOT_FOUND_AUDIO_FILE)
						default:
							Logger.Println("Error in switch IsDuplicate. Default case.")
						}
					}
					lastKey = currentKey
				}
			default:
				//fmt.Printf("1Last key:%v Current key: %v\n", lastKey, currentKey)
				lastKey = currentKey
			}
		}
	}()
	// TODO: add option to pause keyboard capturing
	if !run {
		wg.Done()
	}
}

func readClipboard() string {
	return string(clipboard.Read(clipboard.FmtText))
}
