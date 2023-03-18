package configModule

import (
	. "ScamfariHelper/modules/logModule"
	"github.com/spf13/viper"
)

var WALLETS_FILE = "wallets.txt"
var FOUND_AUDIO_FILE = "old.wav"
var NOT_FOUND_AUDIO_FILE = "new.wav"

func LoadConfig(filename string) bool {
	viper.SetConfigName(filename)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		Logger.Println("Fatal error in config file loading: ", err)
		return false
	}
	WALLETS_FILE = viper.GetString("WALLETS_FILE")
	if len(viper.GetString("FOUND_AUDIO_FILE")) >= 4 {
		FOUND_AUDIO_FILE = viper.GetString("FOUND_AUDIO_FILE")
	}
	if len(viper.GetString("NOT_FOUND_AUDIO_FILE")) >= 4 {
		NOT_FOUND_AUDIO_FILE = viper.GetString("NOT_FOUND_AUDIO_FILE")
	}

	return true
}
