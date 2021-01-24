// Package helpers stores all the helper functions that are used for parsing requests to the api.
package helpers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// SendJSON recieves a ResponseWriter, a status, and a body then sends a json to the
// ResponseWriter
func SendJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Printf("error while encoding, %v", err)
	}
}
