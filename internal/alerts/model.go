package alerts

import "time"

type Alert struct {
	ID        int
	UserID    int
	GameID    int
	Type      string
	Active    bool
	CreatedAt time.Time
}
