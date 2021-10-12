package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IceWreck/HookMsg/actions"
	"github.com/IceWreck/HookMsg/config"
	"github.com/IceWreck/HookMsg/handlers"
	"github.com/rs/zerolog"
)

func main() {

	app := &config.Application{
		Config: config.LoadConfig(),
		Logger: zerolog.New(os.Stderr).With().Logger(),
	}

	app.Logger.Info().
		Bool("script_hook", true).
		Bool("tg_hook", app.Config.TelegramEnabled).
		Bool("matrix_hook", app.Config.MatrixEnabled).
		Msg("")

	if app.Config.TelegramEnabled {
		app.TelegramClient = actions.TelegramClientInit(app) // Login to Telegram
	}

	if app.Config.MatrixEnabled {
		app.MatrixClient = actions.MatrixClientInit(app) // Login to Matrix
	}

	// Initialize Router
	r := handlers.Routes(app)

	app.Logger.Info().Int("port", app.Config.DeploymentPort).Msg("HookMsg Running")
	err := http.ListenAndServe(fmt.Sprintf(":%d", app.Config.DeploymentPort), r)
	if err != nil {
		app.Logger.Err(err).Msg("")
	}
}
