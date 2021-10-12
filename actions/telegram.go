package actions

import (
	"fmt"

	"github.com/IceWreck/HookMsg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// start the tg poller
func TelegramClientInit(app *config.Application) *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(app.Config.TelegramToken)
	if err != nil {
		app.Logger.Panic().Err(err).Msg("Error connecting to Telegram")
	}
	bot.Debug = false
	app.Logger.Info().Str("account", bot.Self.UserName).Msg("Telegram authorized")
	go tgPoller(app, bot)
	return bot
}

func tgPoller(app *config.Application, bot *tgbotapi.BotAPI) {
	// TODO: repace this with a webhook
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		app.Logger.Panic().Err(err).Msg("Error connecting to Telegram")
	}

	for update := range updates {
		if update.Message == nil {
			// ignore any non-Message Updates
			continue
		}
		msgText := fmt.Sprintf(
			"Hi %s ! Your Chat ID is %d",
			update.Message.From.UserName,
			update.Message.Chat.ID,
		)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
		msg.ReplyToMessageID = update.Message.MessageID
		_, err := bot.Send(msg)
		if err != nil {
			app.Logger.Error().Err(err).Msg("Error sending Telegram message")
		}
	}
}

// SendTelegramText - sent plaintext message
func SendTelegramText(app *config.Application, subject string, body string) {
	//msgTemplate := template.New("TelegramMessage")
	//msgText, err := msgTemplate.Parse("Hello {{.Name}}, your marks are {{.Marks}}%!")
	msgText := fmt.Sprintf("<b><u>%s</u></b>\n\n\n%s", subject, body)
	msg := tgbotapi.NewMessage(app.Config.TelegramUserChatID, msgText)
	msg.ParseMode = "HTML"
	_, err := app.TelegramClient.Send(msg)
	if err != nil {
		app.Logger.Error().Err(err).Msg("Error sending Telegram message")
	}
}
