// +build hooktelegram

package actions

import (
	"fmt"
	"log"

	"github.com/IceWreck/HookMsg/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var tgBot tgbotapi.BotAPI

// start the tg poller
func init() {
	go func() {
		bot, err := tgbotapi.NewBotAPI(config.Config.TelegramToken)
		if err != nil {
			log.Panic(err)
		}
		bot.Debug = false
		log.Printf("Authorized on account %s", bot.Self.UserName)
		tgBot = *bot
		tgPoller()
	}()
}

func tgPoller() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := tgBot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
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
		_, err := tgBot.Send(msg)
		if err != nil {
			log.Print(err)
		}
	}
}

// SendMsg - sent plaintext message
func SendMsg(subject string, body string) {
	//msgTemplate := template.New("TelegramMessage")
	//msgText, err := msgTemplate.Parse("Hello {{.Name}}, your marks are {{.Marks}}%!")
	msgText := fmt.Sprintf("<b><u>%s</u></b>\n\n\n%s", subject, body)
	msg := tgbotapi.NewMessage(config.Config.TelegramUserChatID, msgText)
	msg.ParseMode = "HTML"
	_, err := tgBot.Send(msg)
	if err != nil {
		log.Print(err)
	}
}
