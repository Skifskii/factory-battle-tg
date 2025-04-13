package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramBotToken string
	DB               Database
}

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

func New() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error with .env load: ", err)
	}
	return Config{
		TelegramBotToken: os.Getenv("TELEGRAM_BOT_TOKEN"),
		DB: Database{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			Name:     os.Getenv("DB_NAME"),
		},
	}
}
