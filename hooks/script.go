package hooks

import (
	"WebMsg/actions"
	"net/http"

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
	w.Write([]byte("OK"))
}
