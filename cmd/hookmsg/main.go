package main

import (
	"github.com/IceWreck/HookMsg/handlers"
	"github.com/IceWreck/HookMsg/internal/logger"
	"github.com/rs/zerolog/log"
)

func main() {
	logger.SetupLogging()

	app := handlers.NewApplication()
	err := app.Start()
	if err != nil {
		log.Err(err).Msg("")
	}
}
