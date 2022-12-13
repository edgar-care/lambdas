package lib

import (
    "encoding/json"
    "net/http"
)

func WriteResponse(w http.ResponseWriter, body interface{}, status int) {
    bytes, err := json.Marshal(body)
    CheckError(err)
    w.WriteHeader(status)
    _, _ = w.Write(bytes)
}
