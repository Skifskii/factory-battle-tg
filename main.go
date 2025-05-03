// Package main is the entry point of the application. It initializes configuration,
// sets up the Telegram bot, creates the game manager, and starts the Telegram bot.
package main

import (
	"log"

	"github.com/Skifskii/factory-battle-tg/pkg/telegram"

	"github.com/Skifskii/factory-battle-tg/pkg/game"

	"github.com/Skifskii/factory-battle-tg/pkg/config"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	cfg := config.New()

	// database
	// dbConfig := postgres.Config{
	// 	Username: cfg.DB.User,
	// 	Password: cfg.DB.Password,
	// 	Host:     cfg.DB.Host,
	// 	Port:     cfg.DB.Port,
	// 	Database: cfg.DB.Name,
	// }
	// pgClient, err := postgres.NewClient(context.TODO(), dbConfig, 3)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// repo := postgres.NewRepository(pgClient)

	// telegram
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramBotToken)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	roomManager := game.NewGameManager()

	telegramBot := telegram.New(bot, roomManager)
	telegramBot.Start()
}
