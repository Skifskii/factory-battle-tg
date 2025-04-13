package telegram

import (
	"context"
	"log"
	"main/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const commandStart = "start"
const alreadyRegistered = "Вы уже зарегистрированы!"
const successfullyRegistered = "Вы успешно зарегистрированы!"
const unknownCommand = "Я не знаю такую команду"

func (b *Bot) handleCommand(ctx context.Context, msg *tgbotapi.Message) error {
	switch msg.Command() {
	case commandStart:
		return b.handleStartCommand(ctx, msg)
	default:
		ans := tgbotapi.NewMessage(msg.Chat.ID, unknownCommand)
		_, err := b.bot.Send(ans)
		return err
	}
}

func (b *Bot) handleStartCommand(ctx context.Context, msg *tgbotapi.Message) error {
	exists, err := b.storage.IsExists(ctx, msg.Chat.ID)
	if err != nil {
		return err
	}

	// already registered
	if exists {
		ans := tgbotapi.NewMessage(msg.Chat.ID, alreadyRegistered)
		_, err := b.bot.Send(ans)
		return err
	}

	// new user
	if err := b.storage.AddUser(ctx, domain.NewUser(msg.Chat.ID)); err != nil {
		return err
	}
	ans := tgbotapi.NewMessage(msg.Chat.ID, successfullyRegistered)
	_, err = b.bot.Send(ans)

	return err
}

func (b *Bot) handleMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	b.bot.Send(msg)
}
