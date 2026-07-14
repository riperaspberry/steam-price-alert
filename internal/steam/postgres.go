package steam

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

func (r *PostgresRepository) Create(ctx context.Context, game Game) error {
	query := `INSERT INTO games(steam_id, name, price) VALUES ($1,$2,$3)`
	_, err := r.db.ExecContext(ctx, query, game.SteamID, game.Name, game.Price)
	return err
}

func (r *PostgresRepository) GetBySteamID(ctx context.Context, steamID int64) (Game, error) {
	query := `SELECT id, steam_id, name, price FROM games WHERE steam_id = $1`
	row := r.db.QueryRowContext(ctx, query, steamID)
	var game Game
	err := row.Scan(&game.ID, &game.SteamID, &game.Name, &game.Price)
	if err != nil {
		return Game{}, err
	}
	return game, nil
}

func (r *PostgresRepository) UpdatePrice(ctx context.Context, steamID int64, price float64) error {
	query := `UPDATE games SET price = $1 WHERE steam_id = $2`
	_, err := r.db.ExecContext(ctx, query, price, steamID)
	return err
}
