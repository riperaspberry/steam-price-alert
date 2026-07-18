package steam

import (
	"context"
	"database/sql"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddGame(ctx context.Context, game Game) (Game, error) {
	existing, err := s.repo.GetBySteamID(ctx, game.SteamID)

	if err == nil {
		return existing, nil
	}

	if err != sql.ErrNoRows {
		return Game{}, err
	}

	err = s.repo.Create(ctx, game)
	if err != nil {
		return Game{}, err
	}

	return s.repo.GetBySteamID(ctx, game.SteamID)
}
