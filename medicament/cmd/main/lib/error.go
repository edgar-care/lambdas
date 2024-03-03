package lib

import (
	"encoding/json"
	"net/http"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type ErrorResponse struct {
	Message string `json:"message"`
}

// WriteError envoie une réponse d'erreur JSON avec le code HTTP spécifié et le message donné.
func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	response := ErrorResponse{Message: message}
	json.NewEncoder(w).Encode(response)
}
