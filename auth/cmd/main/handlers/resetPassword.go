package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgar_auth "github.com/edgar-care/edgarlib/auth"
	"net/http"
)

type ResetPasswordInput struct {
	Email       string `json:"email"`
	NewPassword string `json:"new_password"`
}

func ResetPassword(w http.ResponseWriter, req *http.Request) {
	var input ResetPasswordInput
	uuid := req.URL.Query().Get("uuid")
	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	resp := edgar_auth.ResetPassword(input.Email, input.NewPassword, uuid)
	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}
	lib.WriteResponse(w, map[string]string{}, resp.Code)
}
