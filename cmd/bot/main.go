package main

import (
	"context"
	"fmt"

	"github.com/riperaspberry/steam-price-alert/internal/alerts"
	"github.com/riperaspberry/steam-price-alert/internal/bot"
	"github.com/riperaspberry/steam-price-alert/internal/config"
	"github.com/riperaspberry/steam-price-alert/internal/database"
	"github.com/riperaspberry/steam-price-alert/internal/steam"
	"github.com/riperaspberry/steam-price-alert/internal/users"
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

	userRepo := users.NewPostgresRepository(db)
	userService := users.NewService(userRepo)
	steamRepo := steam.NewPostgresRepository(db)
	steamService := steam.NewService(steamRepo)
	alertRepo := alerts.NewPostgresRepository(db)
	alertService := alerts.NewService(alertRepo)

	tgBot, err := bot.New(cfg.TelegramToken, cfg.ProxyURL, userService, steamService, alertService)
	if err != nil {
		panic(err)
	}

	tgBot.Start(context.Background())
}
