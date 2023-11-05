package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Instead of passing zerolog.Logger everywhere, its easier to just change
// the default one while starting the app and then calling it everywhere.
func SetupLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	})
	// zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
