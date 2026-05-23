package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	AppEnv     string
	AppVersion string
	PORT       int
	IP         string
}

var configuration *Config

func loadConfigs() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		panic(err)
	}

	configuration = &Config{
		AppName:    os.Getenv("APP_NAME"),
		AppEnv:     os.Getenv("APP_ENV"),
		AppVersion: os.Getenv("APP_VERSION"),
		PORT:       int(port),
		IP:         os.Getenv("IP_ADDRESS"),
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfigs()
	}

	return configuration
}
