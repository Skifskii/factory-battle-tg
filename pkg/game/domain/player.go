// Package domain defines the core domain models used in the game,
// such as Player, User, Room, and related entities.
package domain

// Player - игрок.
type Player struct {
	*User
	Score int64
}

// NewPlayer - конструктор игрока.
func NewPlayer(u *User) *Player {
	return &Player{
		User: u,
	}
}
