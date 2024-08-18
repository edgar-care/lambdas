package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"net/http"
)

func Login2faBackupCode(w http.ResponseWriter, req *http.Request) {
	var accountId string

	var input edgarlib.Login2faSaveCodeInput
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
		passwordCheck := utils.CheckPassword(input.Password, patient.Password)
		if !passwordCheck {
			lib.WriteError(w, http.StatusUnauthorized, "Username and password mismatch")
			return
		} else {
			accountId = patient.ID
		}
	} else {
		doctor, doctorErr := graphql.GetDoctorByEmail(input.Email)
		if doctorErr == nil {
			passwordCheck := utils.CheckPassword(input.Password, doctor.Password)
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

	logSaveCode := edgarlib.Login2faSaveCode(input, device.ID)
	if logSaveCode.Err != nil {
		lib.WriteError(w, logSaveCode.Code, logSaveCode.Err.Error())
		return
	}

	utils.DeviceConnectMiddleware(w, req, logSaveCode.Token)

	lib.WriteResponse(w, map[string]interface{}{
		"token": logSaveCode.Token,
	}, logSaveCode.Code)
}
