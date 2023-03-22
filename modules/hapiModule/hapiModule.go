package hapiModule

import (
	. "ScamfariHelper/modules/logModule"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Catlog struct {
	Wallet string `json:"wallet"`
	Chain  string `json:"chain"`
}

var HapiDB []Catlog

func ReadJSON() bool {
	file, _ := ioutil.ReadFile("hapiDB.json")

	data := []Catlog{}

	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		Logger.Println("Error in hapiModule:ReadJSON() function. Error: ", err)
	}

	HapiDB = data

	fmt.Println("HapiDB records count: ", len(HapiDB))
	if len(HapiDB) >= 1 {
		return true
	} else {
		return false
	}
}
