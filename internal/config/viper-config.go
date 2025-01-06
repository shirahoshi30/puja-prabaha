package config

import (
	"log"

	"github.com/spf13/viper"
)

// viper configures json file
func ConfigViper() {

	viper.SetConfigFile("./internal/config/config.json")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("error loading the file.")
	}
}
