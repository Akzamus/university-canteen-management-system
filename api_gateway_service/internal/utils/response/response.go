package response

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode response data", http.StatusInternalServerError)
		return
	}
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	response := map[string]string{"error": message}
	RespondWithJSON(w, statusCode, response)
}
