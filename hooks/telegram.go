package hooks

import (
	"WebMsg/actions"
	"net/http"
)

// TelegramHook is the endpoint where the user will POST the message they wanna send
func TelegramHook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValues := r.Form
	actions.SendMsg(formValues["subject"][0], formValues["content"][0])
	w.Write([]byte("OK"))
}
