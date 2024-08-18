package lib

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/edgarlib"
)

func WriteResponse(w http.ResponseWriter, body interface{}, status int) {
	bytes, err := json.Marshal(body)
	edgarlib.CheckError(err)
	w.WriteHeader(status)
	_, _ = w.Write(bytes)
}
