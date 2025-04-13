package main

import (
	"context"
	"log"
	"main/internal/config"
	"main/internal/storage/postgres"
	"main/internal/telegram"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg := config.New()

	// database
	dbConfig := postgres.Config{
		Username: cfg.DB.User,
		Password: cfg.DB.Password,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Database: cfg.DB.Name,
	}
	pgClient, err := postgres.NewClient(context.TODO(), dbConfig, 3)
	if err != nil {
		log.Fatal(err)
	}
	repo := postgres.NewRepository(pgClient)

	// telegram
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	telegramBot := telegram.New(bot, repo)
	telegramBot.Start()
}
