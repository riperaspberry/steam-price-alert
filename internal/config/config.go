package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName       string
	TelegramToken string
	DatabaseURL   string
	CheckInterval int
}

func Load() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	checkInterval, err := strconv.Atoi(os.Getenv("CHECK_INTERVAL"))
	if err != nil {
		panic("invalid check interval")
	}

	return Config{
		AppName:       os.Getenv("APP_NAME"),
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		DatabaseURL:   os.Getenv("DATABASE_URL"),
		CheckInterval: checkInterval,
	}, nil
}
