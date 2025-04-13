package postgres

import (
	"context"
	"main/internal/domain"
)

// ToDo: protect from SQL injection.
// AddUser adds a new user to the database.
func (r *Repository) AddRoom(ctx context.Context, room domain.Room) (int64, error) {
	q := `INSERT INTO rooms (leader) VALUES($1) RETURNING room_id`

	var roomID int64
	if err := r.Client.QueryRow(ctx, q, room.LeaderID).Scan(&roomID); err != nil {
		return 0, err
	}

	return roomID, nil
}

func (r *Repository) IsRoomExists(ctx context.Context, id int64) (bool, error) {
	q := `SELECT COUNT(*) FROM rooms WHERE room_id = $1`

	var count int64
	if err := r.Client.QueryRow(ctx, q, id).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}
