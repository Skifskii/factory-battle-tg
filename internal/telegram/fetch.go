package telegram

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) fetchUpdate(ctx context.Context, upd tgbotapi.Update) {
	// not a message
	if upd.Message == nil {
		return
	}

	// command
	if upd.Message.IsCommand() {
		b.processCommand(ctx, upd.Message)
		return
	}

	// message
	b.processMessage(upd.Message)
}
