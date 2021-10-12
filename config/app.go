package config

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/matrix-org/gomatrix"
	"github.com/rs/zerolog"
)

// Application struct to hold the dependencies for our application.
type Application struct {

	// General Dependencies
	Config Config
	Logger zerolog.Logger

	// Hook Specific Dependencies

	MatrixClient   *gomatrix.Client
	TelegramClient *tgbotapi.BotAPI
}
