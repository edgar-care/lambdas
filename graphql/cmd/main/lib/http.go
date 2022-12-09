package lib

import "encoding/json"

type ErrorBody struct {
	Message string `json:"message"`
}

func ErrorMarshal(message string) []byte {
	var bytes []byte
	body := ErrorBody{
		Message: message,
	}
	bytes, _ = json.Marshal(body)
	return bytes
}
