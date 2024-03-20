package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgar_auth "github.com/edgar-care/edgarlib/auth"
	"net/http"
)

type CreatePatientInput struct {
	Email string `json:"email"`
}

func CreatePatientAccount(w http.ResponseWriter, req *http.Request) {
	var input CreatePatientInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	resp := edgar_auth.CreatePatientAccount(input.Email)

	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
	}

	lib.WriteResponse(w, map[string]string{
		"id": resp.Id,
	}, resp.Code)
}
