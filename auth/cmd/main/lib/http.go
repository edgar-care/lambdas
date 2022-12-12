package lib

import (
    "encoding/json"
    "net/http"
)

func WriteResponse(w http.ResponseWriter, body interface{}, status int) {
    bytes, err := json.Marshal(body)
    CheckError(err)
    _, err = w.Write(bytes)
    CheckError(err)
    w.WriteHeader(status)
}
