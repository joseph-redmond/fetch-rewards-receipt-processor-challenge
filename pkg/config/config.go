package config

import (
	"github.com/spf13/viper"
	"log"
)

// Function that initializes the environment variables
func Init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error occurred while reading config file %s", err)
	}
}
