package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"net/http"
)

func Login2faThirdParty(w http.ResponseWriter, req *http.Request) {
	var input authlib.Login2faThirdPartyInput
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

	logThirdParty := authlib.Login2faThirdParty(input, device.ID)
	if logThirdParty.Err != nil {
		lib.WriteError(w, logThirdParty.Code, logThirdParty.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"token": logThirdParty.Token,
	}, logThirdParty.Code)
}
