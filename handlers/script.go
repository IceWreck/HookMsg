package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/actions"

	"github.com/go-chi/chi"
)

// ScriptHook is the endpoint where the user will POST the script's name
func ScriptHook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValues := r.Form
	formResults := map[string]string{
		"endpoint": chi.URLParam(r, "endpoint"),
		"secret":   formValues.Get("secret"),
	}

	go actions.RunScript(formResults)
	renderJSON(w, r, http.StatusOK, map[string]string{"status": "ok"})
}
