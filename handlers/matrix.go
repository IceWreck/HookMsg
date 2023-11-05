package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog/log"
)

// matrixHook is the endpoint where the user will POST the message
// and matrix channel will be a URL parameter
// api key/secret and content can be url-encoded POST/GET form parameters
// they can also be JSON encoded and sent in the POST body
func (app *Application) matrixHook(w http.ResponseWriter, r *http.Request) {
	// check if channel is in config file
	channel, exists := app.config.MatrixChannels[chi.URLParam(r, "channel")]
	if !exists {
		app.errorResponse(w, r, http.StatusBadRequest, "channel does not exist")
		return
	}
	quit := false
	var secret, content string

	// for compatibility reasons we need to support both json and url-encoded bodies

	if app.hasContentType(r, "application/json") {
		var respJSON map[string]string
		err := app.readJSON(w, r, &respJSON)
		if err != nil {
			log.Error().Err(err).Str("channel", channel.ID).
				Msg("Error parsing JSON for matrix hook")
			app.errorResponse(w, r, http.StatusUnprocessableEntity, "invalid JSON")
			quit = true
			return
		}
		secret = respJSON["secret"]
		content = respJSON["content"]
	} else {
		err := r.ParseForm()
		if err != nil {
			log.Error().Err(err).Str("channel", channel.ID).
				Msg("Error parsing url-encoded form for matrix hook")
			app.errorResponse(w, r, http.StatusUnprocessableEntity, "invalid URL encoded body")
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
		app.actionsService.SendMatrixText(channel.ID, content)
		app.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"}, nil)
	} else {
		app.errorResponse(w, r, http.StatusUnauthorized, "unauthorized")
	}
}
