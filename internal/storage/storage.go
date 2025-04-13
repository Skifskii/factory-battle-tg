package storage

import (
	"context"
	"main/internal/domain"
)

type Storage interface {
	AddUser(ctx context.Context, u domain.User) error
	GetUser(ctx context.Context, id int64) (domain.User, error)
	IsUserExists(ctx context.Context, id int64) (bool, error)
	AddRoom(ctx context.Context, room domain.Room) (int64, error)
	IsRoomExists(ctx context.Context, id int64) (bool, error)
	AddPlayer(ctx context.Context, p domain.Player) (int64, error)
	GetPlayer(ctx context.Context, id int64) (domain.Player, error)
	IsUserInRoom(ctx context.Context, userID, roomID int64) (bool, error)
}
