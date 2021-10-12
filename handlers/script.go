package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/actions"
	"github.com/IceWreck/HookMsg/config"

	"github.com/go-chi/chi"
)

// scriptHook is the endpoint where the user will POST the script's name
func scriptHook(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		formValues := r.Form
		formResults := map[string]string{
			"endpoint": chi.URLParam(r, "endpoint"),
			"secret":   formValues.Get("secret"),
		}

		go actions.RunScript(app, formResults)
		renderJSON(w, r, http.StatusOK, map[string]string{"status": "ok"})
	}
}
