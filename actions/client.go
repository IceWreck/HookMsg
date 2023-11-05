package actions

import (
	"github.com/IceWreck/HookMsg/internal/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/matrix-org/gomatrix"
)

// Service struct holds deps like telegram, matrix clients etc needed for this service.
type Service struct {
	// general dependencies
	config config.Config

	// clients for external services
	matrixClient   *gomatrix.Client
	telegramClient *tgbotapi.BotAPI
}

// NewService creates a new instance of this service.
func NewService(cfg config.Config) Service {
	svc := Service{
		config: cfg,
	}
	if svc.config.MatrixEnabled {
		svc.initMatrixClient()
	}
	if svc.config.TelegramEnabled {
		svc.initTelegramClient()
	}

	return svc
}
