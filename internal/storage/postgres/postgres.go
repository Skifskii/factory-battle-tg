package postgres

import (
	"context"
	"fmt"
	"log"
	"main/internal/domain"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Config is the configuration for the database.
type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

// Repository is a database repository.
type Repository struct {
	Client Client
}

func NewRepository(client Client) *Repository {
	return &Repository{Client: client}
}

// Client is a wrapper around the pgxpool.Pool.
type Client interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
}

// NewClient returns a new Client.
func NewClient(ctx context.Context, cfg Config, maxAttempts int64) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	err = doWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, maxAttempts, 5*time.Second)

	if err != nil {
		log.Fatal("error do with tries postgres")
	}

	return pool, nil
}

// ToDo: protect from SQL injection.
// AddUser adds a new user to the database.
func (r *Repository) AddUser(ctx context.Context, u domain.User) error {
	q := `INSERT INTO users (id) VALUES($1)`

	if _, err := r.Client.Exec(ctx, q, u.ID); err != nil {
		return err
	}

	return nil
}

// GetUser returns a user by id.
func (r *Repository) GetUser(ctx context.Context, id int64) (domain.User, error) {
	q := `SELECT id, user_name FROM users WHERE id = $1`

	var u domain.User
	if err := r.Client.QueryRow(ctx, q, id).Scan(&u.ID, &u.Name); err != nil {
		return domain.User{}, err
	}

	return u, nil
}

// IsExists checks if user exists in database.
func (r *Repository) IsExists(ctx context.Context, id int64) (bool, error) {
	q := `SELECT COUNT(*) FROM users WHERE id = $1`

	var count int64
	if err := r.Client.QueryRow(ctx, q, id).Scan(&count); err != nil {
		return false, err
	}

	return count > 0, nil
}

func doWithTries(fn func() error, attemtps int64, delay time.Duration) (err error) {
	for attemtps > 0 {
		if err = fn(); err != nil {
			time.Sleep(delay)
			attemtps--

			continue
		}

		return nil
	}

	return
}
