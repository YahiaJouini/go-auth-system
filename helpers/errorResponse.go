package helpers
import (
	"encoding/json"
	"net/http"
)


type ErrorResponse struct {
    Message string `json:"message"`
    Code    int    `json:"code"`
}


func RespondWithError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	response := ErrorResponse{
		Message: message,
		Code:    code,
	}
	json.NewEncoder(w).Encode(response)
}