package alerts

import "context"

type Repository interface {
	Create(ctx context.Context, alert Alert) error
	GetByID(ctx context.Context, id int) (Alert, error)
	GetUserAlerts(ctx context.Context, userID int) ([]UserAlert, error)
	Deactivate(ctx context.Context, id int) error
}
