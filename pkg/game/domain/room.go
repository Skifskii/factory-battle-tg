package domain

type RoomStatus string

const (
	StatusWaiting  RoomStatus = "waiting"
	StatusPlaying  RoomStatus = "playing"
	StatusFinished RoomStatus = "finished"
)

type Room struct {
	ID           int64
	LeaderID     int64
	CurrentRound int64
	Players      map[int64]*Player
	Status       RoomStatus
	MaxRounds    int64
}

func NewRoom(id, leaderID int64) *Room {
	return &Room{
		ID:        id,
		LeaderID:  leaderID,
		MaxRounds: 2, // ToDo: make this configurable
		Status:    StatusWaiting,
		Players:   make(map[int64]*Player),
	}
}
