package telegram

import (
	"context"
	"fmt"
	"log"
	"main/internal/domain"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const alreadyRegisteredAns = "Вы уже зарегистрированы!"
const successfullyRegisteredAns = "Вы успешно зарегистрированы!"
const successfullyCreatedRoomAns = "Комната создана! Ее ID = %d"

const unknownRoomID = "Комнаты с ID=%d не существует :("
const alreadyInAnotherRoom = "Вы не можете присоединиться к другой комнате. Сейчас вы находитесь в комнате с ID = %d"
const alreadyInRoom = "Вы уже находитесь в этой комнате!"
const successfullyJoinedRoomAns = "Теперь вы находитесь в комнате с ID = %d!"

const gameStartedNotification = "Игра началась! Вы находитесь в комнате с ID = %d.\n\nХод номер 1."

const unknownCommandAns = "Я не знаю такую команду"

func (b *Bot) processCommand(ctx context.Context, msg *tgbotapi.Message) error {
	switch msg.Command() {
	case "start":
		return b.processStartCommand(ctx, msg)
	case "create_room":
		return b.processCreateRoomCommand(ctx, msg)
	case "join_room":
		return b.processJoinRoomCommand(ctx, msg)
	case "start_game":
		return b.processStartGameCommand(ctx, msg)
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
	// reject if the user is in another room
	roomID, err := b.storage.GetUserActiveRoom(ctx, msg.Chat.ID)
	if err != nil {
		return err
	}
	if roomID != 0 {
		_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(alreadyInAnotherRoom, roomID)))
		return err
	}

	roomID, err = b.storage.AddRoom(ctx, domain.NewRoom(msg.Chat.ID))
	if err != nil {
		return err
	}

	// create Player and add to Room
	_, err = b.storage.AddPlayer(ctx, domain.NewPlayer(0, msg.Chat.ID, roomID))
	if err != nil {
		return err
	}

	_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(successfullyCreatedRoomAns, roomID)))

	return err
}

func (b *Bot) processJoinRoomCommand(ctx context.Context, msg *tgbotapi.Message) error {
	// reject if the user is in another room
	roomID, _ := b.storage.GetUserActiveRoom(ctx, msg.Chat.ID)
	if roomID != 0 {
		_, err := b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(alreadyInAnotherRoom, roomID)))
		return err
	}

	roomID, err := parseInt64(msg.CommandArguments())
	if err != nil {
		return err
	}

	// search for a room
	roomExists, err := b.storage.IsRoomExists(ctx, roomID)
	if err != nil {
		return err
	}
	if !roomExists {
		_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(unknownRoomID, roomID)))
		return err
	}

	// create Player and add to Room
	var inRoom bool
	inRoom, err = b.storage.IsUserInRoom(ctx, msg.Chat.ID, roomID)
	if err != nil {
		return err
	}
	if inRoom {
		_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, alreadyInRoom))
		return err
	}

	_, err = b.storage.AddPlayer(ctx, domain.NewPlayer(0, msg.Chat.ID, roomID))
	if err != nil {
		return err
	}

	_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(successfullyJoinedRoomAns, roomID)))

	return err
}

func (b *Bot) processStartGameCommand(ctx context.Context, msg *tgbotapi.Message) error {
	roomID, err := b.storage.GetWaitingRoomByLeaderID(ctx, msg.Chat.ID)
	if err != nil {
		return err
	}

	users, err := b.storage.GetUsersByRoomID(ctx, roomID)
	if err != nil {
		return err
	}
	for _, user := range users {
		_, err = b.bot.Send(tgbotapi.NewMessage(user.ID, fmt.Sprintf(gameStartedNotification, roomID)))
		if err != nil {
			return err
		}
	}

	return b.storage.IncrementCurrentRound(ctx, roomID)
}

func (b *Bot) processUnknownCommand(msg *tgbotapi.Message) error {
	_, err := b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, unknownCommandAns))
	return err
}

// parseInt64 attempts to parse a string to int64
// Returns the parsed value and nil error if successful
// Returns 0 and error if parsing fails
func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
