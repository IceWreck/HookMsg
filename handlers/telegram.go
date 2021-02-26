package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/actions"
	"github.com/IceWreck/HookMsg/config"
)

// TelegramHook is the endpoint where the user will POST the message they wanna send
func TelegramHook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValues := r.Form

	// Send message if secret is in config.json
	isAuthorized := false
	for _, password := range config.Config.TelegramWebhookAuth {
		if password == formValues["secret"][0] {
			actions.SendMsg(formValues["subject"][0], formValues["content"][0])
			isAuthorized = true
			renderJSON(w, r, http.StatusOK, map[string]string{"status": "ok"})
			return
		}
	}
	if !isAuthorized {
		renderJSON(w, r, http.StatusUnauthorized, map[string]string{"err": "unauthorized"})
	}

}
