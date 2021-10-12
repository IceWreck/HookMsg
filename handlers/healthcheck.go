package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/config"
)

func healthCheck(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"status": "available",
			"system_info": map[string]string{
				"environment": "dev",
				"version":     config.Version,
			},
		}
		err := writeJSON(app, w, http.StatusOK, data, nil)
		if err != nil {
			serverErrorResponse(app, w, r, err)
		}
	}
}
