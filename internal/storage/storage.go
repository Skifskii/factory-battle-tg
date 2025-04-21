package storage

// import (
// 	"context"
// 	"main/internal/game/domain"
// )

// type Storage interface {
// 	AddUser(ctx context.Context, u domain.User) error
// 	GetUser(ctx context.Context, id int64) (domain.User, error)
// 	IsUserExists(ctx context.Context, id int64) (bool, error)
// 	AddRoom(ctx context.Context, room domain.Room) (int64, error)
// 	IsRoomExists(ctx context.Context, id int64) (bool, error)
// 	AddPlayer(ctx context.Context, p domain.Player) (int64, error)
// 	GetPlayer(ctx context.Context, id int64) (domain.Player, error)
// 	IsUserInRoom(ctx context.Context, userID, roomID int64) (bool, error)
// 	GetUserActiveRoom(ctx context.Context, userID int64) (int64, error)
// 	GetWaitingRoomByLeaderID(ctx context.Context, id int64) (int64, error)
// 	GetUsersByRoomID(ctx context.Context, id int64) ([]domain.User, error)
// 	IncrementCurrentRound(ctx context.Context, roomID int64) error
// }
