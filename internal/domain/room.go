package domain

type Room struct {
	ID           int64
	Leader       User
	CurrentRound int64
}

func NewRoom(leader User) Room {
	return Room{
		Leader: leader,
	}
}
