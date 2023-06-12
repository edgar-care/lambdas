package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	"github.com/edgar-care/auth/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func Login(w http.ResponseWriter, req *http.Request) {
	var input services.LoginInput

	t := chi.URLParam(req, "type")

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	resp, err := services.Login(input, t)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp,
		}, http.StatusBadRequest)
	}

	lib.WriteResponse(w, map[string]interface{}{
		"token": resp,
	}, http.StatusOK)
}
