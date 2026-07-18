package alerts

type UserAlert struct {
	ID      int
	GameID  int
	Name    string
	SteamID int64
	Price   float64
	Active  bool
}
