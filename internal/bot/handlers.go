package bot

import (
	"context"

	tgbot "github.com/go-telegram/bot"
	models "github.com/go-telegram/bot/models"
	"github.com/riperaspberry/steam-price-alert/internal/users"
)

func (b *Bot) registerHandlers() {
	b.bot.RegisterHandler(tgbot.HandlerTypeMessageText, "/start", tgbot.MatchTypeExact, b.startHandler)
}

func (b *Bot) startHandler(ctx context.Context, bot *tgbot.Bot, update *models.Update) {
	user := users.User{
		TelegramID: update.Message.From.ID,
		Username:   update.Message.From.Username,
	}

	err := b.userService.Register(ctx, user)
	if err != nil {
		return
	}
	_, err = bot.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Ты зарегистрирован!",
	})
	if err != nil {
		return
	}
}
