package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IceWreck/HookMsg/config"
)

// renders map into JSON
func renderJSON(w http.ResponseWriter, r *http.Request, status int, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
	return
}

// The writeJSON() method will render your data structure as a JSON object with relevant headers.
func writeJSON(app *config.Application, w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	// js, err := json.MarshalIndent(data, "", "\t")
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	js = append(js, '\n')
	for key, value := range headers {
		w.Header()[key] = value
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return nil
}
