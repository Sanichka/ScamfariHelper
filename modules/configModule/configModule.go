package configModule

import (
	. "ScamfariHelper/modules/logModule"
	"github.com/spf13/viper"
)

var WALLETS_FILE = "wallets.txt"

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

	return true
}
