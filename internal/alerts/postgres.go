package alerts

import (
	"context"
	"database/sql"
)

type PostgresRepository struct {
	db *sql.DB
}

var _ Repository = (*PostgresRepository)(nil)

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Create(ctx context.Context, alert Alert) error {
	query := `INSERT INTO alerts (user_id, game_id, type, active)VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, alert.UserID, alert.GameID, alert.Type, alert.Active)
	return err
}

func (r *PostgresRepository) GetByID(ctx context.Context, id int) (Alert, error) {
	query := `SELECT id, user_id, game_id, type, active, created_at FROM alerts WHERE id = $1`
	row := r.db.QueryRowContext(ctx, query, id)
	var alert Alert
	err := row.Scan(&alert.ID, &alert.UserID, &alert.GameID, &alert.Type, &alert.Active, &alert.CreatedAt)
	if err != nil {
		return Alert{}, err
	}
	return alert, nil
}

func (r *PostgresRepository) GetUserAlerts(ctx context.Context, userID int) ([]UserAlert, error) {
	query := `SELECT alerts.id, games.id, games.name, games.steam_id, games.price, alerts.active FROM alerts JOIN games ON games.id = alerts.game_id WHERE alerts.user_id = $1`
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var alerts []UserAlert
	for rows.Next() {
		var alert UserAlert
		err := rows.Scan(&alert.ID, &alert.GameID, &alert.Name, &alert.SteamID, &alert.Price, &alert.Active)
		if err != nil {
			return nil, err
		}
		alerts = append(alerts, alert)
	}

	return alerts, rows.Err()
}

func (r *PostgresRepository) Deactivate(ctx context.Context, id int) error {
	query := `UPDATE alerts SET active = false WHERE id =$1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
