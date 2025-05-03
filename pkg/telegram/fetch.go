package telegram

import (
	"context"

	"github.com/Skifskii/factory-battle-tg/pkg/game/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) fetchUpdate(ctx context.Context, upd tgbotapi.Update, nCh chan domain.Notification) {
	// not a message
	if upd.Message == nil {
		return
	}

	// command
	if upd.Message.IsCommand() {
		b.processCommand(ctx, upd.Message, nCh)
		return
	}

	// message
	b.processMessage(upd.Message)
}
