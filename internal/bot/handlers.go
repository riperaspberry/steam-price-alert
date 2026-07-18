package bot

import (
	"context"
	"strconv"
	"strings"

	tgbot "github.com/go-telegram/bot"
	models "github.com/go-telegram/bot/models"
	"github.com/riperaspberry/steam-price-alert/internal/alerts"
	"github.com/riperaspberry/steam-price-alert/internal/steam"
	"github.com/riperaspberry/steam-price-alert/internal/users"
)

func (b *Bot) registerHandlers() {
	b.bot.RegisterHandler(tgbot.HandlerTypeMessageText, "/start", tgbot.MatchTypeExact, b.startHandler)
	b.bot.RegisterHandler(tgbot.HandlerTypeMessageText, "/add", tgbot.MatchTypePrefix, b.addHandler)
	b.bot.RegisterHandler(tgbot.HandlerTypeMessageText, "/list", tgbot.MatchTypeExact, b.listHandler)
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
	user, err := b.userService.GetByTelegramID(
		ctx,
		update.Message.From.ID,
	)
	if err != nil {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Пользователь не найден. Напишите /start",
		})
		return
	}
	game := steam.Game{
		SteamID: steamID,
		Name:    "Unknown",
		Price:   0,
	}
	game, err = b.steamService.AddGame(ctx, game)
	if err != nil {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Не удалось добавить игру.",
		})
		return
	}

	alert := alerts.Alert{
		UserID: user.ID,
		GameID: game.ID,
		Type:   "price",
		Active: true,
	}

	err = b.alertService.CreateAlert(ctx, alert)
	if err != nil {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Не удалось создать уведомление.",
		})
		return
	}
	_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Игра добавлена.",
	})
}

func (b *Bot) listHandler(ctx context.Context, bot *tgbot.Bot, update *models.Update) {
	user, err := b.userService.GetByTelegramID(
		ctx,
		update.Message.From.ID,
	)
	if err != nil {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Сначала напишите /start",
		})
		return
	}

	alerts, err := b.alertService.GetUserAlerts(ctx, user.ID)
	if err != nil {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "Не удалось получить список игр.",
		})
		return
	}

	if len(alerts) == 0 {
		_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   "У тебя пока нет отслеживаемых игр.",
		})
		return
	}

	text := "Твои игры:\n\n"

	for i, alert := range alerts {
		text += strconv.Itoa(i+1) + ". "
		text += alert.Name + "\n"
		text += "Steam ID: " + strconv.FormatInt(alert.SteamID, 10) + "\n"
		text += "Цена: " + strconv.FormatFloat(alert.Price, 'f', 2, 64) + "\n\n"
	}

	_, _ = bot.SendMessage(ctx, &tgbot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   text,
	})
}
