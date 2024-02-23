package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgar_auth "github.com/edgar-care/edgarlib/auth"
	"net/http"
)

type MissingPasswordInput struct {
	Email string `json:"email"`
}

func MissingPassword(w http.ResponseWriter, req *http.Request) {
	var input MissingPasswordInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	resp := edgar_auth.MissingPassword(input.Email)
	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}
	lib.WriteResponse(w, map[string]string{}, resp.Code)
}
