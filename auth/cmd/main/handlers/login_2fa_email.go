package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"net/http"
)

func Login2faEmail(w http.ResponseWriter, req *http.Request) {
	var input authlib.Login2faEmailInput
	var accountId string

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	if !checkerAccount(input.Email, input.Password) {
		lib.WriteError(w, http.StatusBadRequest, "Invalid email")
		return
	}

	patient, patientErr := graphql.GetPatientByEmail(input.Email)
	if patientErr == nil {
		passwordCheck := authlib.CheckPassword(input.Password, patient.Password)
		if !passwordCheck {
			lib.WriteError(w, http.StatusUnauthorized, "Username and password mismatch")
			return
		} else {
			accountId = patient.ID
		}
	} else {
		doctor, doctorErr := graphql.GetDoctorByEmail(input.Email)
		if doctorErr == nil {
			passwordCheck := authlib.CheckPassword(input.Password, doctor.Password)
			if !passwordCheck {
				lib.WriteError(w, http.StatusUnauthorized, "Username and password mismatch")
				return
			} else {
				accountId = doctor.ID
			}
		} else {
			lib.WriteError(w, http.StatusBadRequest, "Email does not correspond to a valid patient or doctor")
		}
	}

	utils.DeviceConnectMiddleware(w, req, accountId)
	device := utils.GetCurrentUserDevice(w, req, accountId)

	logEmail := authlib.Login2faEmail(input, device.ID)
	if logEmail.Err != nil {
		lib.WriteError(w, logEmail.Code, logEmail.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"token": logEmail.Token,
	}, logEmail.Code)
}

func checkerAccount(email string, password string) bool {
	patient, patientErr := graphql.GetPatientByEmail(email)
	if patientErr != nil {
		doctor, doctorErr := graphql.GetDoctorByEmail(email)
		if doctorErr != nil {
			return false
		}
		checkingPassword := authlib.CheckPassword(password, doctor.Password)
		if checkingPassword != true {
			return false
		}
		return true
	}
	checkingPassword := authlib.CheckPassword(password, patient.Password)
	if checkingPassword != true {
		return false
	}
	return true
}
