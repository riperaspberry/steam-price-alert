package users

import (
	"context"
	"database/sql"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Register(ctx context.Context, user User) error {
	_, err := s.repo.GetByTelegramID(ctx, user.TelegramID)
	if err == nil {
		return nil
	}
	if err != sql.ErrNoRows {
		return err
	}
	return s.repo.Create(ctx, user)
}

func (s *Service) GetByTelegramID(ctx context.Context, telegramID int64) (User, error) {
	return s.repo.GetByTelegramID(ctx, telegramID)
}
