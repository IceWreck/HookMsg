package handlers

import (
	"encoding/json"
	"net/http"
)

// renders map into JSON
func renderJSON(w http.ResponseWriter, r *http.Request, status int, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
	return
}
