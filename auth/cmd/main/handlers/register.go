package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/edgar-care/auth/cmd/main/lib"
	"github.com/edgar-care/auth/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func Register(w http.ResponseWriter, req *http.Request) {
	t := chi.URLParam(req, "type")

	var token string

	if t == "d" {
		var input services.DoctorInput
		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		input.Password = lib.HashPassword(input.Password)
		doctor, err := services.CreateDoctor(input)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Unable to create account: " + err.Error(),
			}, http.StatusBadRequest)
			return
		}
		token, _ = lib.CreateToken(map[string]interface{}{
			"doctor": doctor,
		})
	} else if t == "a" {
		var input services.AdminInput
		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		if lib.VerifyToken(input.Token) == false {
			lib.WriteResponse(w, map[string]string{
				"message": "Unable to create account: Invalid Token",
			}, http.StatusBadRequest)
			return
		}

		input.Password = lib.HashPassword(input.Password)
		admin, err := services.CreateAdmin(input)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Unable to create account: " + err.Error(),
			}, http.StatusBadRequest)
			return
		}
		token, _ = lib.CreateToken(map[string]interface{}{
			"admin": admin,
		})
	} else {
		var input services.PatientInput

		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		input.Password = lib.HashPassword(input.Password)
		patient, err := services.CreatePatient(input)
		if err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": "Unable to create account: " + strings.ToLower(err.Error()[9:]),
			}, 400)
			return
		}
		token, _ = lib.CreateToken(map[string]interface{}{
			"patient": patient,
		})
	}

	lib.WriteResponse(w, map[string]string{
		"token": token,
	}, 200)
}
