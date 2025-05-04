package game

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Skifskii/factory-battle-tg/pkg/game/domain"
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

func (rm *RoomManager) ProcessSession(nCh chan domain.Notification) {
	rm.Room.Status = domain.StatusPlaying
	rm.Room.CurrentRound = 1

	for rm.Room.CurrentRound <= rm.Room.MaxRounds {
		rm.processRound(nCh)
		rm.Room.CurrentRound++
	}

	rm.determineWinners(nCh)

	rm.Room.Status = domain.StatusFinished

	close(rm.mCh) // Close the move channel to signal no more moves
}

func (rm *RoomManager) determineWinners(nCh chan domain.Notification) {
	var maxScore int64 = -1 << 63 // minimum int64 value
	for _, player := range rm.Room.Players {
		if player.Score > maxScore {
			maxScore = player.Score
		}
	}

	for _, player := range rm.Room.Players {
		if player.Score == maxScore {
			nCh <- domain.Notification{
				ToPlayer: *player,
				Text:     fmt.Sprintf("Поздравляем\\! Вы выиграли с %d очками\\!", player.Score),
			}
		}
	}
}

func (rm *RoomManager) ProcessMove(playerID int64, cardName string) error {
	if rm.Room.Status != domain.StatusPlaying {
		return errors.New("game is not running")
	}

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

func (rm *RoomManager) processRound(nCh chan domain.Notification) {
	moves := make([]domain.Move, 0, len(rm.Room.Players))
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
		nCh <- domain.Notification{
			ToPlayer: *p,
			Text:     rep,
		}
	}
}

func (rm *RoomManager) createRoundReport(moves []domain.Move) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("Раунд %d\n\n", rm.Room.CurrentRound))

	// Prepare slices for columns to calculate max width
	var ids []string
	var movesStr []string
	var scores []string

	// Add header names first
	ids = append(ids, "Имя")
	movesStr = append(movesStr, "Ход")
	scores = append(scores, "Очки")

	for _, move := range moves {
		player := rm.Room.Players[move.PlayerID]
		ids = append(ids, fmt.Sprintf("%d", player.ID))
		movesStr = append(movesStr, string(move.CardName))
		scores = append(scores, fmt.Sprintf("%d", player.Score))
	}

	// Calculate max width for each column
	maxLen := func(arr []string) int {
		max := 0
		for _, s := range arr {
			if len([]rune(s)) > max {
				max = len(s)
			}
		}
		return max
	}

	idWidth := maxLen(ids)
	moveWidth := maxLen(movesStr)
	scoreWidth := maxLen(scores)

	builder.WriteString("```\n")
	// Format header
	builder.WriteString(fmt.Sprintf("%-*s | %-*s | %-*s\n", idWidth, "Имя", moveWidth, "Ход", scoreWidth, "Очки"))

	// Format rows
	for i := 1; i < len(ids); i++ {
		builder.WriteString(fmt.Sprintf("%-*s | %-*s | %-*s\n", idWidth, ids[i], moveWidth, movesStr[i], scoreWidth, scores[i]))
	}
	builder.WriteString("```\n")

	return builder.String()
}
