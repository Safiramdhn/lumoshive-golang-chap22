package utils

import (
	"encoding/json"
	"golang-beginner-22/models"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(models.Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	})
}
