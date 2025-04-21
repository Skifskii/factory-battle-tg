package domain

type Player struct {
	*User
	Score int64
}

func NewPlayer(u *User) *Player {
	return &Player{
		User: u,
	}
}
