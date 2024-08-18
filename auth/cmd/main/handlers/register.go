package handlers

import (
	"encoding/json"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
)

func Register(w http.ResponseWriter, req *http.Request) {
	t := chi.URLParam(req, "type")

	var resp authlib.LoginResponse

	if t == "d" {
		var input authlib.DoctorInput

		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		account, err := authlib.RegisterDoctor(input.Email, input.Password, input.Name, input.Firstname, input.Address)
		if err != nil {
			lib.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.DeviceConnectMiddleware(w, req, account.ID)
		device := utils.GetCurrentUserDevice(w, req, account.ID)

		resp = authlib.Login(authlib.LoginInput{
			Email:    input.Email,
			Password: input.Password,
		}, "d", device.ID)
	} else if t == "a" {
		var input authlib.AdminInput
		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		response := authlib.RegisterAndLoginAdmin(input.Email, input.Password, input.Name, input.LastName, input.Token)
		resp = authlib.LoginResponse{
			Token: response.Token,
			Code:  response.Code,
			Err:   response.Err,
		}
	} else {
		var input authlib.PatientInput

		err := json.NewDecoder(req.Body).Decode(&input)
		lib.CheckError(err)

		account, err := authlib.RegisterPatient(input.Email, input.Password)
		if err != nil {
			lib.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		utils.DeviceConnectMiddleware(w, req, account.ID)
		device := utils.GetCurrentUserDevice(w, req, account.ID)

		resp = authlib.Login(authlib.LoginInput{
			Email:    input.Email,
			Password: input.Password,
		}, "p", device.ID)
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
