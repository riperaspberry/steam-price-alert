package alerts

import (
	"context"
	"errors"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateAlert(ctx context.Context, alert Alert) error {
	return s.repo.Create(ctx, alert)
}

func (s *Service) GetUserAlerts(ctx context.Context, userID int) ([]UserAlert, error) {
	return s.repo.GetUserAlerts(ctx, userID)
}

func (s *Service) Deactivate(ctx context.Context, id int) error {
	return s.repo.Deactivate(ctx, id)
}

func (s *Service) DeactivateUserAlert(ctx context.Context, alertID int, userID int) error {
	alert, err := s.repo.GetByID(ctx, alertID)
	if err != nil {
		return err
	}

	if alert.UserID != userID {
		return errors.New("not your alert")
	}

	return s.repo.Deactivate(ctx, alertID)
}
