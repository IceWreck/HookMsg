package handlers

import (
	"net/http"
)

// telegramHook is the endpoint where the user will POST the message they wanna send
func (app *Application) telegramHook(w http.ResponseWriter, r *http.Request) {
	content := r.PostFormValue("content")
	subject := r.PostFormValue("subject")
	secret := r.PostFormValue("secret")

	// send message if secret is in config file
	isAuthorized := false
	for _, password := range app.config.TelegramKey {
		if password == secret {
			app.actionsService.SendTelegramText(subject, content)
			isAuthorized = true
			app.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"}, nil)
			return
		}
	}
	if !isAuthorized {
		app.errorResponse(w, r, http.StatusUnauthorized, "unauthorized")
	}
}
