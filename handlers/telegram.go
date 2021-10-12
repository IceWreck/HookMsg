package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/actions"
	"github.com/IceWreck/HookMsg/config"
)

// TelegramHook is the endpoint where the user will POST the message they wanna send
func TelegramHook(w http.ResponseWriter, r *http.Request) {
	content := r.PostFormValue("content")
	subject := r.PostFormValue("subject")
	secret := r.PostFormValue("secret")

	// send message if secret is in config file
	isAuthorized := false
	for _, password := range config.Config.TelegramKey {
		if password == secret {
			actions.SendMsg(subject, content)
			isAuthorized = true
			renderJSON(w, r, http.StatusOK, map[string]string{"status": "ok"})
			return
		}
	}
	if !isAuthorized {
		renderJSON(w, r, http.StatusUnauthorized, map[string]string{"err": "unauthorized"})
	}

}
