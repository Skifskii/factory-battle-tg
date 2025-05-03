package postgres

// import (
// 	"context"
// 	"github.com/Skifskii/factory-battle-tg/pkg/game/domain"
// )

// // ToDo: protect from SQL injection.
// func (r *Repository) AddPlayer(ctx context.Context, p domain.Player) (int64, error) {
// 	q := `INSERT INTO players (room_id, user_id) VALUES($1, $2) RETURNING room_id`

// 	var playerID int64
// 	if err := r.Client.QueryRow(ctx, q, p.RoomID, p.UserID).Scan(&playerID); err != nil {
// 		return 0, err
// 	}

// 	return playerID, nil
// }

// func (r *Repository) GetPlayer(ctx context.Context, id int64) (domain.Player, error) {
// 	q := `SELECT player_id, room_id, user_id, score FROM players WHERE player_id = $1`

// 	var p domain.Player
// 	if err := r.Client.QueryRow(ctx, q, id).Scan(&p.ID, &p.RoomID, &p.UserID, &p.Score); err != nil {
// 		return domain.Player{}, err
// 	}

// 	return p, nil
// }

// // IsUserExists checks if user exists in database.
// func (r *Repository) IsUserInRoom(ctx context.Context, userID, roomID int64) (bool, error) {
// 	q := `SELECT COUNT(*) FROM players WHERE user_id = $1 AND room_id = $2`

// 	var count int64
// 	if err := r.Client.QueryRow(ctx, q, userID, roomID).Scan(&count); err != nil {
// 		return false, err
// 	}

// 	return count > 0, nil
// }
