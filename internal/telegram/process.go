package telegram

import (
	"context"
	"fmt"
	"log"
	"main/internal/domain"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const alreadyRegisteredAns = "Вы уже зарегистрированы!"
const successfullyRegisteredAns = "Вы успешно зарегистрированы!"
const successfullyCreatedRoomAns = "Комната создана! Ее ID = %d"

const unknownCommandAns = "Я не знаю такую команду"

func (b *Bot) processCommand(ctx context.Context, msg *tgbotapi.Message) error {
	switch msg.Command() {
	case "start":
		return b.processStartCommand(ctx, msg)
	case "create_room":
		return b.processCreateRoomCommand(ctx, msg)
	default:
		return b.processUnknownCommand(msg)
	}
}

func (b *Bot) processMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
}

func (b *Bot) processStartCommand(ctx context.Context, msg *tgbotapi.Message) error {
	exists, err := b.storage.IsUserExists(ctx, msg.Chat.ID)
	if err != nil {
		return err
	}

	// already registered
	if exists {
		_, err := b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, alreadyRegisteredAns))
		return err
	}

	// new user
	if err := b.storage.AddUser(ctx, domain.NewUser(msg.Chat.ID)); err != nil {
		return err
	}
	_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, successfullyRegisteredAns))

	return err
}

func (b *Bot) processCreateRoomCommand(ctx context.Context, msg *tgbotapi.Message) error {
	roomID, err := b.storage.AddRoom(ctx, domain.NewRoom(domain.NewUser(msg.Chat.ID)))
	if err != nil {
		return err
	}

	_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(successfullyCreatedRoomAns, roomID)))

	return err
}

func (b *Bot) processUnknownCommand(msg *tgbotapi.Message) error {
	_, err := b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, unknownCommandAns))
	return err
}
