package actions

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rs/zerolog/log"
)

func (svc *Service) initTelegramClient() {
	bot, err := tgbotapi.NewBotAPI(svc.config.TelegramToken)
	if err != nil {
		log.Fatal().Err(err).Msg("Error creating telegram client")
	}
	bot.Debug = false
	log.Info().Str("account", bot.Self.UserName).Msg("Created telegram client")
	svc.telegramClient = bot

	go startTelegramPoller(svc.telegramClient)
}

func startTelegramPoller(bot *tgbotapi.BotAPI) {
	// TODO: repace this with a webhook
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		// okay to die here because it only happens at program init
		log.Fatal().Err(err).Msg("Error connecting to Telegram")
	}

	// poll for updates
	for update := range updates {
		if update.Message == nil {
			// ignore any non-message updates
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
			log.Error().Err(err).Msg("Error sending telegram reply")
		}
	}
}

// SendTelegramText sends a plaintext message.
func (svc *Service) SendTelegramText(subject string, body string) {
	if svc.telegramClient == nil {
		log.Error().Msg("Cannot send telegram text, client has not been initialized")
		return
	}

	//msgTemplate := template.New("TelegramMessage")
	//msgText, err := msgTemplate.Parse("Hello {{.Name}}, your marks are {{.Marks}}%!")
	msgText := fmt.Sprintf("<b><u>%s</u></b>\n\n\n%s", subject, body)
	msg := tgbotapi.NewMessage(svc.config.TelegramUserChatID, msgText)
	msg.ParseMode = "HTML"
	_, err := svc.telegramClient.Send(msg)
	if err != nil {
		log.Error().Err(err).Msg("Error sending Telegram message")
	}
}
