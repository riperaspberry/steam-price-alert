package steam

import "context"

type Repository interface {
	Create(ctx context.Context, game Game) error
	GetBySteamID(ctx context.Context, steamID int64) (Game, error)
	UpdatePrice(ctx context.Context, steamID int64, price float64) error
	// GetAll(ctx context.Context) ([]Game, error)
}
