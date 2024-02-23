package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	edgar_auth "github.com/edgar-care/edgarlib/auth"
	"github.com/go-chi/chi/v5"
)

func Register(w http.ResponseWriter, req *http.Request) {
	t := chi.URLParam(req, "type")

	var resp edgar_auth.RegisterAndLoginResponse

	if t == "d" {
		var input edgar_auth.DoctorInput
		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		resp = edgar_auth.RegisterAndLoginDoctor(input.Email, input.Password, input.Name, input.Firstname, input.Address)
	} else if t == "a" {
		var input edgar_auth.AdminInput
		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		resp = edgar_auth.RegisterAndLoginAdmin(input.Email, input.Password, input.Name, input.LastName, input.Token)
	} else {
		var input edgar_auth.PatientInput

		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		resp = edgar_auth.RegisterAndLoginPatient(input.Email, input.Password)
	}
	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}
	lib.WriteResponse(w, map[string]string{
		"token": resp.Token,
	}, resp.Code)
}
