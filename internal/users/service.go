package users

import "context"

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
	return s.repo.Create(ctx, user)
}
