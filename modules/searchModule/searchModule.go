package searchModule

import (
	"ScamfariHelper/modules/configModule"
	"ScamfariHelper/modules/keyboardCapture"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

var WalletsList []string

func LoadFile() bool {
	file, err := os.Open(configModule.WALLETS_FILE)
	var result []string
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		result = append(result, strings.Replace(scanner.Text(), "\"", "", -1))
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(result) <= 0 {
		return false
	}

	WalletsList = result
	return true
}

func IsDuplicate(wallet string) bool {
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
	out := make(chan string, 100)
	currentKey := ""
	lastKey := ""
	go keyboardCapture.Start(out)
	go func() {
		for {
			currentKey = <-out
			switch currentKey {
			case "67":
				//fmt.Printf("Last key:%v Current key: %v\n", lastKey, currentKey)
				if lastKey == "162" {
					//fmt.Printf("CTRL+C pressed: Last key:%v Current key: %v\n", lastKey, currentKey)
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
