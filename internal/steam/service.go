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

func (s *Service) AddGame(ctx context.Context, game Game) error {
	_, err := s.repo.GetBySteamID(ctx, game.SteamID)
	if err == nil {
		return nil
	}
	if err != sql.ErrNoRows {
		return err
	}
	return s.repo.Create(ctx, game)
}
