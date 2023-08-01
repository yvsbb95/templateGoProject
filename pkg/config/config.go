package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	TestInt    int    `mapstructure:"TEST_INT"`
	TestString string `mapstructure:"TEST_STRING"`
}

var configLoaded = false
var config *Config

func LoadConfig() *Config {
	if configLoaded {
		return config
	}
	viper.AddConfigPath("./pkg/config")
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	// Loads config from file specified
	err := viper.ReadInConfig()

	// Loads config from env variables
	viper.AutomaticEnv()

	if err != nil {
		log.Panicln("Error reading config from file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Panicln("Error unmarshalling config", err)
	}

	return config
}
