package handlers

import (
	"fmt"
	"net/http"

	"github.com/IceWreck/HookMsg/actions"
	"github.com/IceWreck/HookMsg/internal/config"
	"github.com/rs/zerolog/log"
)

// Application struct to hold the dependencies for our application.
type Application struct {
	config         config.Config
	actionsService *actions.Service
}

func (app *Application) Start() error {
	log.Info().Int("port", app.config.DeploymentPort).Msg("HookMsg running")
	return http.ListenAndServe(fmt.Sprintf(":%d", app.config.DeploymentPort), app.routes())
}

func NewApplication() Application {
	cfg := config.LoadConfig()
	actionsSvc := actions.NewService(cfg)
	app := Application{
		config:         cfg,
		actionsService: &actionsSvc,
	}

	log.Info().
		Bool("script_hook", true).
		Bool("tg_hook", app.config.TelegramEnabled).
		Bool("matrix_hook", app.config.MatrixEnabled).
		Msg("")

	return app
}
