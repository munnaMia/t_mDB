package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName      string
	AppEnv       string
	AppVersion   string
	PORT         int
	IP           string
	KEY_SIZE     int
	PAYLOAD_SIZE int
}

var configuration *Config

func loadConfigs() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	keysize, err := strconv.ParseInt(os.Getenv("KEY_SIZE"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	paylaodsize, err := strconv.ParseInt(os.Getenv("PAYLOAD_SIZE"), 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	configuration = &Config{
		AppName:      os.Getenv("APP_NAME"),
		AppEnv:       os.Getenv("APP_ENV"),
		AppVersion:   os.Getenv("APP_VERSION"),
		PORT:         int(port),
		IP:           os.Getenv("IP_ADDRESS"),
		KEY_SIZE:     int(keysize),
		PAYLOAD_SIZE: int(paylaodsize),
	}
}

func GetConfig() *Config {
	if configuration == nil {
		loadConfigs()
	}

	return configuration
}
