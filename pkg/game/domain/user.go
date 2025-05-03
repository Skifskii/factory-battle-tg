package domain

type User struct {
	ID   int64
	Name string
}

func NewUser(id int64) *User {
	return &User{ID: id}
}
