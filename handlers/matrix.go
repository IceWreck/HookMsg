package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/actions"
	"github.com/IceWreck/HookMsg/config"
	"github.com/go-chi/chi"
)

// matrixHook is the endpoint where the user will POST the message
// and matrix channel will be a URL parameter
// api key/secret and content can be url-encoded POST/GET form parameters
// they can also be JSON encoded and sent in the POST body
func matrixHook(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check if channel is in config file
		channel, exists := app.Config.MatrixChannels[chi.URLParam(r, "channel")]
		if !exists {
			errorResponse(app, w, r, http.StatusBadRequest, "channel does not exist")
			return
		}
		quit := false
		var secret, content string

		// for compatibility reasons we need to support both json and url-encoded bodies

		if hasContentType(r, "application/json") {
			var respJSON map[string]string
			err := readJSON(app, w, r, &respJSON)
			if err != nil {
				app.Logger.Error().Err(err).Str("channel", channel.ID).
					Msg("Error parsing JSON for matrix hook")
				errorResponse(app, w, r, http.StatusUnprocessableEntity, "invalid JSON")
				quit = true
				return
			}
			secret = respJSON["secret"]
			content = respJSON["content"]
		} else {
			err := r.ParseForm()
			if err != nil {
				app.Logger.Error().Err(err).Str("channel", channel.ID).
					Msg("Error parsing url-encoded form for matrix hook")
				errorResponse(app, w, r, http.StatusUnprocessableEntity, "invalid URL encoded body")
				quit = true
				return
			}

			secret = r.PostFormValue("secret")
			content = r.PostFormValue("content")
		}

		if quit {
			return
		}

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
