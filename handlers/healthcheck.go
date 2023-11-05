package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/internal/config"
)

func (app *Application) healthCheck(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"status": "available",
		"system_info": map[string]string{
			"environment": "dev",
			"version":     config.Version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
