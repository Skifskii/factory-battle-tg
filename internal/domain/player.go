package domain

type Player struct {
	User
	ID    int64
	Room  Room
	Score int64
}
