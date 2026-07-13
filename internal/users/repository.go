package users

import "context"

type Repository interface {
	Create(ctx context.Context, user User) error
	GetByTelegramID(ctx context.Context, telegramID int64) (User, error)
}
