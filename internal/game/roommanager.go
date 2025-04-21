package game

import (
	"errors"
	"fmt"
	"main/internal/game/domain"
	"strings"
)

type RoomManager struct {
	Room *domain.Room
	mCh  chan domain.Move
}

func NewRoomManager(r *domain.Room) *RoomManager {
	return &RoomManager{
		Room: r,
		mCh:  make(chan domain.Move),
	}
}

func (rm *RoomManager) AddPlayer(p *domain.Player) error {
	if _, exist := rm.Room.Players[p.ID]; exist {
		return errors.New("player already exists in room")
	}
	rm.Room.Players[p.ID] = p

	return nil
}

func (rm *RoomManager) ProcessSession(notificationsChanel chan domain.Notification) {
	rm.Room.Status = domain.StatusPlaying
	rm.Room.CurrentRound = 1

	for rm.Room.CurrentRound <= rm.Room.MaxRounds {
		rm.processRound(notificationsChanel)
		rm.Room.CurrentRound++
	}

	rm.Room.Status = domain.StatusFinished
}

func (rm *RoomManager) ProcessMove(playerID int64, cardName string) error {
	m := domain.Move{PlayerID: playerID}

	switch cardName {
	case "inc":
		m.CardName = domain.CardIncrease
	case "dec":
		m.CardName = domain.CardDecrease
	default:
		return errors.New("invalid card name")
	}

	rm.mCh <- m

	return nil
}

func (rm *RoomManager) processRound(notificationsChanel chan domain.Notification) {
	moves := make([]domain.Move, len(rm.Room.Players))
	for i := 0; i < len(rm.Room.Players); i++ {
		moves = append(moves, <-rm.mCh)
	}

	// ToDo: реализовать логику обработки раунда
	for _, move := range moves {
		switch move.CardName {
		case domain.CardIncrease:
			rm.Room.Players[move.PlayerID].Score += 10
		case domain.CardDecrease:
			rm.Room.Players[move.PlayerID].Score -= 5
		}
	}

	rep := rm.createRoundReport(moves)
	for _, p := range rm.Room.Players {
		notificationsChanel <- domain.Notification{
			ToPlayer: *p,
			Text:     rep,
		}
	}
}

func (rm *RoomManager) createRoundReport(moves []domain.Move) string {
	// ToDo: реализовать логику создания отчета раунда
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Раунд %d.\n\n", rm.Room.CurrentRound))

	for _, move := range moves {
		builder.WriteString(fmt.Sprintf("Игрок %d: %s\n", move.PlayerID, move.CardName))
	}

	return builder.String()
}
