package postgres

// import (
// 	"context"
// 	"main/internal/game/domain"
// )

// // ToDo: protect from SQL injection.
// // AddUser adds a new user to the database.
// func (r *Repository) AddUser(ctx context.Context, u domain.User) error {
// 	q := `INSERT INTO users (user_id) VALUES($1)`

// 	if _, err := r.Client.Exec(ctx, q, u.ID); err != nil {
// 		return err
// 	}

// 	return nil
// }

// // GetUser returns a user by id.
// func (r *Repository) GetUser(ctx context.Context, id int64) (domain.User, error) {
// 	q := `SELECT user_id, user_name FROM users WHERE user_id = $1`

// 	var u domain.User
// 	if err := r.Client.QueryRow(ctx, q, id).Scan(&u.ID, &u.Name); err != nil {
// 		return domain.User{}, err
// 	}

// 	return u, nil
// }

// // IsUserExists checks if user exists in database.
// func (r *Repository) IsUserExists(ctx context.Context, id int64) (bool, error) {
// 	q := `SELECT COUNT(*) FROM users WHERE user_id = $1`

// 	var count int64
// 	if err := r.Client.QueryRow(ctx, q, id).Scan(&count); err != nil {
// 		return false, err
// 	}

// 	return count > 0, nil
// }

// // GetUsers returns a user by id.
// func (r *Repository) GetUsersByRoomID(ctx context.Context, id int64) ([]domain.User, error) {
// 	q := `SELECT user_id FROM players WHERE room_id = $1`

// 	rows, err := r.Client.Query(ctx, q, id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	users := make([]domain.User, 0)

// 	for rows.Next() {
// 		var u domain.User

// 		err = rows.Scan(&u.ID)
// 		if err != nil {
// 			return nil, err
// 		}

// 		users = append(users, u)
// 	}

// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }
