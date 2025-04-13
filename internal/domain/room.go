package domain

type Room struct {
	ID           int64
	Leader       User
	CurrentRound int64
}
