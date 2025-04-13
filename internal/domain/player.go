package domain

type Player struct {
	ID     int64
	UserID int64
	RoomID int64
	Score  int64
}

func NewPlayer(id, userID, roomID int64) Player {
	return Player{
		ID:     id,
		UserID: userID,
		RoomID: roomID,
	}
}
