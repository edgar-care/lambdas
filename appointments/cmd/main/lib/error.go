package lib

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func WriteError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	response := ErrorResponse{Message: message}
	json.NewEncoder(w).Encode(response)
}
