package postgres

// import (
// 	"context"
// 	"main/internal/game/domain"
// )

// // ToDo: protect from SQL injection.
// // AddUser adds a new user to the database.
// func (r *Repository) AddRoom(ctx context.Context, room domain.Room) (int64, error) {
// 	q := `INSERT INTO rooms (leader) VALUES($1) RETURNING room_id`

// 	var roomID int64
// 	if err := r.Client.QueryRow(ctx, q, room.LeaderID).Scan(&roomID); err != nil {
// 		return 0, err
// 	}

// 	return roomID, nil
// }

// func (r *Repository) IsRoomExists(ctx context.Context, id int64) (bool, error) {
// 	q := `SELECT COUNT(*) FROM rooms WHERE room_id = $1`

// 	var count int64
// 	if err := r.Client.QueryRow(ctx, q, id).Scan(&count); err != nil {
// 		return false, err
// 	}

// 	return count > 0, nil
// }

// // GetUserActiveRoom checks if the user is in active rooms.
// func (r *Repository) GetUserActiveRoom(ctx context.Context, id int64) (int64, error) {
// 	q := `
// 	SELECT r.room_id
// 	FROM rooms r
// 	JOIN players rp ON r.room_id = rp.room_id
// 	WHERE rp.user_id = $1
// 	AND r.winner_id = 0
// 	LIMIT 1;
// 	`

// 	var roomID int64
// 	if err := r.Client.QueryRow(ctx, q, id).Scan(&roomID); err != nil {
// 		return 0, err
// 	}

// 	return roomID, nil
// }

// // GetUserActiveRoom checks if the user is in active rooms.
// func (r *Repository) GetWaitingRoomByLeaderID(ctx context.Context, id int64) (int64, error) {
// 	q := `
// 	SELECT room_id
// 	FROM rooms
// 	WHERE leader = $1 AND winner_id = 0 AND current_round = 0
// 	LIMIT 1;
// 	`

// 	var roomID int64
// 	if err := r.Client.QueryRow(ctx, q, id).Scan(&roomID); err != nil {
// 		return 0, err
// 	}

// 	return roomID, nil
// }

// // GetUserActiveRoom checks if the user is in active rooms.
// func (r *Repository) IncrementCurrentRound(ctx context.Context, roomID int64) error {
// 	q := `UPDATE rooms SET current_round = current_round + 1 WHERE room_id = $1`
// 	if _, err := r.Client.Exec(ctx, q, roomID); err != nil {
// 		return err
// 	}

// 	return nil
// }
