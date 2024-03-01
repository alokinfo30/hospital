// response_handler.go

package Helper

import (
	"encoding/json"
	"net/http"
)

// RespondJSON writes a JSON response with the given status code and data
func RespondJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if data != nil {
		// Encode the data as JSON and write it to the response writer
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

// RespondError writes an error response with the given status code and message
func RespondError(w http.ResponseWriter, statusCode int, errorMessage string) {
	RespondJSON(w, statusCode, map[string]string{"error": errorMessage})
}
