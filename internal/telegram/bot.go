package telegram

import (
	"context"
	"log"
	"main/internal/storage"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	storage storage.Storage
}

func New(bot *tgbotapi.BotAPI, s storage.Storage) *Bot {
	return &Bot{
		bot:     bot,
		storage: s,
	}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()
	b.handleUpdates(context.TODO(), updates)
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(ctx, update.Message)
			continue
		}

		b.handleMessage(update.Message)
	}
}
