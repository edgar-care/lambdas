package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	"github.com/edgar-care/auth/cmd/main/services"
	edgarlib "github.com/edgar-care/edgarlib/auth"
	"github.com/go-chi/chi/v5"
)

func Register(w http.ResponseWriter, req *http.Request) {
	t := chi.URLParam(req, "type")

	var token string
	var code int
	var err error

	if t == "d" {
		var input edgarlib.DoctorInput
		err = json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		token, code, err = services.RegisterAndLoginDoctor(input)
	} else if t == "a" {
		var input services.AdminInput
		err = json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		token, code, err = services.RegisterAndLoginAdmin(input)
	} else {
		var input services.PatientInput

		err = json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		token, code, err = services.RegisterAndLoginPatient(input)
	}
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": err.Error(),
		}, code)
		return
	}
	lib.WriteResponse(w, map[string]string{
		"token": token,
	}, code)
}
