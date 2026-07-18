package bot

import (
	"context"
	"strconv"
	"strings"

	tgbot "github.com/go-telegram/bot"
	models "github.com/go-telegram/bot/models"
	"github.com/riperaspberry/steam-price-alert/internal/steam"
	"github.com/riperaspberry/steam-price-alert/internal/users"
)

func (b *Bot) registerHandlers() {
	b.bot.RegisterHandler(tgbot.HandlerTypeMessageText, "/start", tgbot.MatchTypeExact, b.startHandler)
	b.bot.RegisterHandler(tgbot.HandlerTypeMessageText, "/add", tgbot.MatchTypePrefix, b.addHandler)
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

func (b *Bot) addHandler(ctx context.Context, bot *tgbot.Bot, update *models.Update) {
	text := update.Message.Text
	parts := strings.Fields(text)
	if len(parts) != 2 {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Использование: /add <steam_id>",
		})
		return
	}
	steamID, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Steam ID должен быть числом.",
		})
		return
	}
	game := steam.Game{
		SteamID: steamID,
		Name:    "Unknown",
		Price:   0,
	}
	err = b.steamService.AddGame(ctx, game)
	if err != nil {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Не удалось добавить игру.",
		})
		return
	}
	_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Игра добавлена.",
	})
}
