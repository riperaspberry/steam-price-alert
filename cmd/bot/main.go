package main

import (
	"fmt"

	"github.com/riperaspberry/steam-price-alert/internal/config"
	"github.com/riperaspberry/steam-price-alert/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	db, err := database.New(cfg.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("db connected")
}
