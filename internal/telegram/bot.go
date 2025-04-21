package telegram

import (
	"context"
	"log"
	"main/internal/game"
	"main/internal/game/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	bot *tgbotapi.BotAPI
	gm  *game.GameManager
}

func New(bot *tgbotapi.BotAPI, gm *game.GameManager) *Bot {
	return &Bot{
		bot: bot,
		gm:  gm,
	}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)

	updates := b.initUpdatesChannel()
	nCh := make(chan domain.Notification)

	go b.handleUpdates(context.TODO(), updates, nCh)

	b.handleNotifications(context.TODO(), nCh)
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel, nCh chan domain.Notification) {
	for u := range updates {
		b.fetchUpdate(ctx, u, nCh)
	}
}

func (b *Bot) handleNotifications(ctx context.Context, nCh chan domain.Notification) {
	for n := range nCh {
		b.bot.Send(tgbotapi.NewMessage(n.ToPlayer.ID, n.Text))
	}
}
