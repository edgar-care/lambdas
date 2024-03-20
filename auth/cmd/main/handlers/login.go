package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	edgar_auth "github.com/edgar-care/edgarlib/auth"
	"github.com/go-chi/chi/v5"
)

func Login(w http.ResponseWriter, req *http.Request) {
	var input edgar_auth.LoginInput

	t := chi.URLParam(req, "type")

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	resp := edgar_auth.Login(input, t)

	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"token": resp.Token,
	}, resp.Code)
}
