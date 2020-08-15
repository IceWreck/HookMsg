package actions

import (
	"WebMsg/utils"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var TgBot tgbotapi.BotAPI

// InitializeTG - starts the tg poller
func InitializeTG() {

	bot, err := tgbotapi.NewBotAPI(utils.Config.TelegramToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	TgBot = *bot
	tgPoller()

}

func tgPoller() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := TgBot.GetUpdatesChan(u)
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
		_, err := TgBot.Send(msg)
		if err != nil {
			log.Print(err)
		}
	}
}

func SendMsg(subject string, body string) {
	//msgTemplate := template.New("TelegramMessage")
	//msgText, err := msgTemplate.Parse("Hello {{.Name}}, your marks are {{.Marks}}%!")
	msgText := fmt.Sprintf("<b><u>%s</u></b>\n\n\n%s", subject, body)
	msg := tgbotapi.NewMessage(utils.Config.TelegramUserChatID, msgText)
	msg.ParseMode = "HTML"
	_, err := TgBot.Send(msg)
	if err != nil {
		log.Print(err)
	}
}
