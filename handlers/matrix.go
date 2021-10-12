package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/actions"
	"github.com/IceWreck/HookMsg/config"
	"github.com/go-chi/chi"
)

// matrixHook is the endpoint where the user will POST the message
// and matrix channel will be a URL parameter
// api key/secret will be a POST parameter
func matrixHook(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if channel is in config file
		channel, exists := app.Config.MatrixChannels[chi.URLParam(r, "channel")]
		if !exists {
			errorResponse(app, w, r, http.StatusBadRequest, "channel does not exist")
			return
		}

		err := r.ParseForm()
		if err != nil {
			app.Logger.Error().Err(err).Str("channel", channel.ID).
				Msg("Error parsing url-encoded form for matrix hook")
		}

		secret := r.PostFormValue("secret")
		content := r.PostFormValue("content")
		// verify api key
		if channel.Key == secret {
			// send message
			actions.SendMatrixText(app, channel.ID, content)
			writeJSON(app, w, http.StatusOK, map[string]string{"status": "ok"}, nil)
		} else {
			errorResponse(app, w, r, http.StatusUnauthorized, "unauthorized")
		}
	}
}
