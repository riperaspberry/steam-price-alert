package users

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

var _ Repository = (*PostgresRepository)(nil)

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{db: db}
}

func (r *PostgresRepository) Create(ctx context.Context, user User) error {
	query := `INSERT INTO users (telegram_id, username) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, user.TelegramID, user.Username)
	return err
}

func (r *PostgresRepository) GetByTelegramID(ctx context.Context, telegramID int64) (User, error) {
	query := `SELECT id,telegram_id, username, created_at FROM users WHERE telegram_id = $1`
	row := r.db.QueryRowContext(ctx, query, telegramID)
	var user User
	err := row.Scan(&user.ID, &user.TelegramID, &user.Username, &user.CreatedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}
