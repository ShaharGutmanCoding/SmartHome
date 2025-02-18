package utils

import (
	"encoding/json"
	"net/http"
)

// JSONResponse sends a JSON response with a given status
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
