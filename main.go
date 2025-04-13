package main

import (
	"context"
	"log"
	"main/internal/storage/postgres"
	"main/internal/telegram"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	// environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("error with .env load: ", err)
	}

	// database
	dbConfig := postgres.Config{
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Database: os.Getenv("DB_NAME"),
	}
	pgClient, err := postgres.NewClient(context.TODO(), dbConfig, 3)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewRepository(pgClient)

	// telegram
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	telegramBot := telegram.New(bot, repo)
	telegramBot.Start()
}
