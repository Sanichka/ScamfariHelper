package searchModule

import (
	"ScamfariHelper/modules/audio"
	"ScamfariHelper/modules/configModule"
	"ScamfariHelper/modules/hapiModule"
	"ScamfariHelper/modules/keyboardCapture"
	. "ScamfariHelper/modules/logModule"
	"bufio"
	"golang.design/x/clipboard"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

var WalletsList []string
var ClipboardContent string

// TODO: Near Postponed due to there is no length limit according to documentation https://nomicon.io/DataStructures/Account
//var NearRegexp = regexp.MustCompile("^(([a-z\\d]+[\\-_])*[a-z\\d]+\\.)*([a-z\\d]+[\\-_])*[a-z\\d]+$")
// TODO add Aurora regexp
// TODO add Solana validation
// TODO Improve to searching wallet even if clipboard contains some other text
// ETHRegexp Works for Polygon(Matic) and OKC chain since they are layer 2 networks
var ETHRegexp = regexp.MustCompile("^0x[a-fA-F0-9]{40}$")

// BSCRegexp & BNBRegexp BNB chain has multiple variants https://dev.binance.vision/t/what-is-the-bnb-java-kotlin-regex/11571
var BSCRegexp = regexp.MustCompile("^(0x)[0-9A-Fa-f]{40}$")
var BNBRegexp = regexp.MustCompile("^(bnb1)[0-9a-z]{38}$")
var BTCRegexp = regexp.MustCompile("^(bc1|[13])[a-zA-HJ-NP-Z0-9]{25,39}$")
var TRONRegexp = regexp.MustCompile("^T[A-Za-z1-9]{33}$")

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
	if configModule.HAPI_VALIDATION {
		for _, v := range hapiModule.HapiDB {
			if v.Wallet == wallet {
				return true
			}
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
					ClipboardContent = strings.TrimSpace(readClipboard())
					// Use advanced validation for wallets if flag is enabled in config
					if configModule.WALLET_VALIDATION {
						if len(ClipboardContent) >= 30 && WalletValidation(ClipboardContent) {
							Search()
						}
					} else if len(ClipboardContent) >= 30 {
						Search()
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

func Search() {
	switch IsDuplicate(ClipboardContent) {
	case true:
		//fmt.Println("Streaming found audio file")
		go audio.StreamAudio(configModule.FOUND_AUDIO_FILE)
	case false:
		//fmt.Println("Streaming NOT found audio file")
		go audio.StreamAudio(configModule.NOT_FOUND_AUDIO_FILE)
		if configModule.AUTO_UPDATE_WALLETS {
			UpdateFile(ClipboardContent)
		}
	default:
		Logger.Println("Error in switch IsDuplicate. Default case.")
	}
}

func readClipboard() string {
	return string(clipboard.Read(clipboard.FmtText))
}

func UpdateFile(newWallet string) {
	WalletsList = append(WalletsList, newWallet)
	f, err := os.OpenFile(configModule.WALLETS_FILE, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	if _, err = f.WriteString(newWallet + " \n"); err != nil {
		panic(err)
	}
}

// WalletValidation TODO refactor in more aesthetic way
func WalletValidation(wallet string) bool {
	IsETH := ETHRegexp.MatchString(wallet)
	IsBNC := BSCRegexp.MatchString(wallet)
	IsBNB := BNBRegexp.MatchString(wallet)
	IsBTC := BTCRegexp.MatchString(wallet)
	IsTRON := TRONRegexp.MatchString(wallet)
	return IsETH || IsBNC || IsBNB || IsBTC || IsTRON
}
