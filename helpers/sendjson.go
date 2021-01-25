package helpers

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// SendJSON recieves a ResponseWriter, a status, and a body then sends a json to the
// ResponseWriter
func SendJSON(w http.ResponseWriter, status int, body interface{}) {
	writeHeader(&w, status)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Printf("error while encoding, %v", err)
	}
}

// writeHeader writes the headers required for setting the application to type json and 
// allowing cors
func writeHeader(w *http.ResponseWriter, status int) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).WriteHeader(status)
}
