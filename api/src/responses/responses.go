package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a requisition response in JSON format
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	// Sets the response format to JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error returns an error in JSON
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}