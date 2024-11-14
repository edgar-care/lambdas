package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgar_auth "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type ResetPasswordInput struct {
	NewPassword string `json:"new_password"`
}

func ResetPassword(w http.ResponseWriter, req *http.Request) {
	var input ResetPasswordInput

	accountype := chi.URLParam(req, "type")

	uuid := req.URL.Query().Get("uuid")
	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	resp := edgar_auth.ResetPassword(input.NewPassword, uuid, accountype)
	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}
	lib.WriteResponse(w, map[string]string{}, resp.Code)
}
