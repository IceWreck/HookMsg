package handlers

import (
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog/log"
)

// scriptHook is the endpoint where the user will GET/POST the script's name, secret
// all headers, payload will be passed to the script as JSON
// this is like CGI essentially
func (app *Application) scriptHook(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValues := r.Form
	endpoint := chi.URLParam(r, "endpoint")
	secret := formValues.Get("secret")

	// request data will be passed to the script after stripping out the HookMsg secret
	formValues.Set("secret", "")

	webhookData := map[string]interface{}{}
	webhookData["Parameters"] = formValues
	webhookData["Method"] = r.Method
	webhookData["Headers"] = r.Header
	resBody, err := io.ReadAll(r.Body)
	if err == nil {
		webhookData["Body"] = string(resBody)
	} else {
		webhookData["Body"] = ""
	}

	go app.actionsService.RunScript(endpoint, secret, webhookData)
	app.writeJSON(w, http.StatusAccepted, map[string]string{"status": "will be scheduled if everything checks out"}, nil)
}

func (app *Application) listScripts(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	formValues := r.Form
	secret := formValues.Get("secret")

	// verify api key
	if app.config.ScriptsListKey != secret {
		app.errorResponse(w, r, http.StatusUnauthorized, "unauthorized")
		return
	}

	scripts, err := app.actionsService.GetAvailableScripts()
	if err != nil {
		log.Error().Err(err).Msg("Could not read scripts file")
		app.errorResponse(w, r, http.StatusInternalServerError, "could not read scripts file")
		return
	}
	// clear out the secret to run the script
	for index := range scripts {
		scripts[index].Secret = ""
	}

	app.writeJSON(w, http.StatusOK, scripts, nil)
}
