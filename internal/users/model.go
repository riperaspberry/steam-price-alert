package users

import "time"

type User struct {
	ID         int
	TelegramID int64
	UserName   string
	CreatedAt  time.Time
}
