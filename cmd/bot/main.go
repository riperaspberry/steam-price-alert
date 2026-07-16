package main

import (
	"context"
	"fmt"

	"github.com/riperaspberry/steam-price-alert/internal/bot"
	"github.com/riperaspberry/steam-price-alert/internal/config"
	"github.com/riperaspberry/steam-price-alert/internal/database"
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
	tgBot, err := bot.New(cfg.TelegramToken, cfg.ProxyURL, userService)
	if err != nil {
		panic(err)
	}

	tgBot.Start(context.Background())
}
