package main

import (
	"fmt"

	"github.com/riperaspberry/steam-price-alert/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println("start")
	fmt.Println(cfg.AppName)
	fmt.Println(cfg.CheckInterval)
}
