package hooks

import (
	"WebMsg/actions"
	"net/http"
)

// TelegramHook is the endpoint where the user will POST the message they wanna send
func ScriptHook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValues := r.Form
	var scriptInfo = actions.Script{
		FileName: formValues["filename"][0],
		Shell:    formValues["shell"][0],
	}
	go actions.RunScript(scriptInfo)
	w.Write([]byte("OK"))
}
