package domain

type Room struct {
	ID      int64
	Leader  User
	Players []User
}
