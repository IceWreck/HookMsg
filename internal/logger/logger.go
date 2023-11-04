package logger

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Instead of passing zerolog.Logger everywhere, its easier to just change
// the default one while starting the app and then calling it everywhere.
func InitLogging() {
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out: os.Stderr,
	})
}
