package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Login2faBackupCode(w http.ResponseWriter, req *http.Request) {
	var accountId string
	var password string
	deviceId := "deviceId"

	t := chi.URLParam(req, "type")

	var input edgarlib.Login2faSaveCodeInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	//if !checkerAccount(input.Email, input.Password) {
	//	lib.WriteError(w, http.StatusBadRequest, "Invalid email")
	//	return
	//}
	if t == "p" {
		patient, err := graphql.GetPatientByEmail(input.Email)
		if err != nil {
			lib.WriteError(w, http.StatusBadRequest, "Email does not correspond to a valid Patient")
			return
		}
		password = patient.Password
		accountId = patient.ID
	}
	if t == "d" {
		doctor, err := graphql.GetDoctorByEmail(input.Email)
		if err != nil {
			lib.WriteError(w, http.StatusBadRequest, "Email does not correspond to a valid Doctor")
			return
		}
		password = doctor.Password
		accountId = doctor.ID
	}

	passwordCheck := utils.CheckPassword(input.Password, password)
	if !passwordCheck {
		lib.WriteError(w, http.StatusUnauthorized, "Email and password mismatch")
		return
	}

	utils.DeviceConnectMiddleware(w, req, accountId)
	deviceId = utils.GetCurrentUserDevice(w, req, accountId).ID

	logSaveCode := edgarlib.Login2faSaveCode(input, deviceId, accountId)
	if logSaveCode.Err != nil {
		lib.WriteError(w, logSaveCode.Code, logSaveCode.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"token": logSaveCode.Token,
	}, logSaveCode.Code)
}
