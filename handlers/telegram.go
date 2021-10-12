package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/actions"
	"github.com/IceWreck/HookMsg/config"
)

// telegramHook is the endpoint where the user will POST the message they wanna send
func telegramHook(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content := r.PostFormValue("content")
		subject := r.PostFormValue("subject")
		secret := r.PostFormValue("secret")

		// send message if secret is in config file
		isAuthorized := false
		for _, password := range app.Config.TelegramKey {
			if password == secret {
				actions.SendTelegramText(app, subject, content)
				isAuthorized = true
				writeJSON(app, w, http.StatusOK, map[string]string{"status": "ok"}, nil)
				return
			}
		}
		if !isAuthorized {
			errorResponse(app, w, r, http.StatusUnauthorized, "unauthorized")
		}
	}
}
