package alerts

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateAlert(ctx context.Context, alert Alert) error {
	return s.repo.Create(ctx, alert)
}
