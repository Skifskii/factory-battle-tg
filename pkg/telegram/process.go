package telegram

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Skifskii/factory-battle-tg/pkg/game/domain"

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

func (b *Bot) processCommand(ctx context.Context, msg *tgbotapi.Message, nCh chan domain.Notification) error {
	switch msg.Command() {
	case "start":
		return b.processStartCommand(ctx, msg)
	case "create_room":
		return b.processCreateRoomCommand(ctx, msg)
	case "join_room":
		return b.processJoinRoomCommand(ctx, msg)
	case "start_game":
		return b.processStartGameCommand(ctx, msg, nCh)
	case "move":
		return b.processMoveCommand(ctx, msg)
	default:
		return b.processUnknownCommand(msg)
	}
}

func (b *Bot) processMessage(message *tgbotapi.Message) {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
}

func (b *Bot) processStartCommand(ctx context.Context, msg *tgbotapi.Message) error {
	_, err := b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, successfullyRegisteredAns))

	return err
}

func (b *Bot) processCreateRoomCommand(ctx context.Context, msg *tgbotapi.Message) error {
	roomID, err := b.gm.AddRoom(msg.Chat.ID)
	if err != nil {
		return err
	}

	_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(successfullyCreatedRoomAns, roomID)))
	return err
}

func (b *Bot) processJoinRoomCommand(ctx context.Context, msg *tgbotapi.Message) error {
	roomID, err := parseInt64(msg.CommandArguments())
	if err != nil {
		return err
	}

	if err = b.gm.AddPlayerToRoom(msg.Chat.ID, roomID); err != nil {
		_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "не удается подключиться к комнате"))
		return err
	}

	_, err = b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, fmt.Sprintf(successfullyJoinedRoomAns, roomID)))
	return err
}

func (b *Bot) processStartGameCommand(ctx context.Context, msg *tgbotapi.Message, nCh chan domain.Notification) error {
	roomID, exists := b.gm.RoomIDByPlayerID[msg.From.ID]
	if !exists {
		_, err := b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "вы не находитесь в комнате"))
		return err
	}

	if err := b.gm.StartGame(roomID, msg.Chat.ID, nCh); err != nil {
		return err
	}

	_, err := b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "игра началась! 1 раунд"))
	return err
}

func (b *Bot) processMoveCommand(ctx context.Context, msg *tgbotapi.Message) error {
	cardName := msg.CommandArguments()

	roomID := b.gm.RoomIDByPlayerID[msg.Chat.ID]
	rm := b.gm.Rooms[roomID]

	if err := rm.ProcessMove(msg.Chat.ID, cardName); err != nil {
		return nil
	}

	_, err := b.bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "ход записан"))
	return err
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
