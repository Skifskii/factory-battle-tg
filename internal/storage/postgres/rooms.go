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
	if err := r.Client.QueryRow(ctx, q, room.Leader.ID).Scan(&roomID); err != nil {
		return 0, err
	}

	return roomID, nil
}
