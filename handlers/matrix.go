package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/config"
	"github.com/go-chi/chi"
)

// MatrixHook is the endpoint where the user will POST the message
// and matrix channel will be a URL parameter
// api key/secret will be a POST parameter
func MatrixHook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValues := r.Form
	// check if channel is in config file
	channel, exists := config.Config.MatrixChannels[chi.URLParam(r, "channel")]
	if !exists {
		renderJSON(w, r, http.StatusBadRequest, map[string]string{"err": "channel does not exist"})
		return
	}
	// assume api key is empty if not provided
	secret := ""
	if len(formValues["secret"]) > 0 {
		secret = formValues["secret"][0]
	}
	// verify api key
	if channel.Key == secret {
		// send message
		renderJSON(w, r, http.StatusOK, map[string]string{"status": "ok"})
	} else {
		renderJSON(w, r, http.StatusUnauthorized, map[string]string{"err": "unauthorized"})
	}
}
