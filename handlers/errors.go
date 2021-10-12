package handlers

import (
	"net/http"

	"github.com/IceWreck/HookMsg/config"
)

// The logError() method is a generic helper for logging an error message.
func logError(app *config.Application, r *http.Request, err error) {
	app.Logger.Error().Stack().Err(err).Str("request_method", r.Method).Str("request_url", r.URL.String()).Msg("")
}

// The errorResponse() method is a generic helper for sending JSON-formatted error
// messages to the client with a given status code. Note that we're using an interface{}
// type for the message parameter, rather than just a string type, as this gives us
// more flexibility over the values that we can include in the response.
func errorResponse(app *config.Application, w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	data := map[string]interface{}{"error": message}
	// Write the response using the writeJSON() helper. If this happens to return an
	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.
	err := writeJSON(app, w, status, data, nil)
	if err != nil {
		logError(app, r, err)
		w.WriteHeader(500)
	}
}

// The serverErrorResponse() method will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses the
// errorResponse() helper to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func serverErrorResponse(app *config.Application, w http.ResponseWriter, r *http.Request, err error) {
	logError(app, r, err)
	message := "the server encountered a problem and could not process your request"
	errorResponse(app, w, r, http.StatusInternalServerError, message)
}
