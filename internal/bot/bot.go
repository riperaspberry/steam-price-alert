package bot

import (
	"context"
	"net/http"
	"net/url"
	"time"

	tgbot "github.com/go-telegram/bot"
	"github.com/riperaspberry/steam-price-alert/internal/users"
)

type Bot struct {
	bot         *tgbot.Bot
	userService *users.Service
}

func New(token string, proxyURL string, userService *users.Service) (*Bot, error) {
	proxy, err := url.Parse(proxyURL)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxy),
		},
	}

	b, err := tgbot.New(
		token,
		tgbot.WithHTTPClient(30*time.Second, client),
	)
	if err != nil {
		return nil, err
	}

	myBot := &Bot{
		bot:         b,
		userService: userService,
	}

	myBot.registerHandlers()

	return myBot, nil
}

func (b *Bot) Start(ctx context.Context) {
	b.bot.Start(ctx)
}
