package domain

type Room struct {
	ID           int64
	LeaderID     int64
	CurrentRound int64
}

func NewRoom(leaderID int64) Room {
	return Room{
		LeaderID: leaderID,
	}
}
