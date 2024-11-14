package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgar_auth "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type MissingPasswordInput struct {
	Email string `json:"email"`
}

func MissingPassword(w http.ResponseWriter, req *http.Request) {
	var input MissingPasswordInput

	accountype := chi.URLParam(req, "type")

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	resp := edgar_auth.MissingPassword(input.Email, accountype)
	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}
	lib.WriteResponse(w, map[string]string{}, resp.Code)
}
