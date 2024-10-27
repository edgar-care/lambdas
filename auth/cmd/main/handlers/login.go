package handlers

import (
	"encoding/json"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
)

type CheckDoubleAuthResponse struct {
	Content map[string]interface{}
	Code    int
	Err     error
}

func Login(w http.ResponseWriter, req *http.Request) {
	var input authlib.LoginInput
	var accountId string

	t := chi.URLParam(req, "type")

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

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

	doubleAuthSent := getDoubleAuth(accountId, w, req)

	if doubleAuthSent.Err == nil && doubleAuthSent.Code != 200 {
		utils.DeviceConnectMiddleware(w, req, accountId)
		device := utils.GetCurrentUserDevice(w, req, accountId)

		resp := authlib.Login(input, t, device.ID)

		if resp.Err != nil {
			lib.WriteResponse(w, map[string]string{
				"message": resp.Err.Error(),
			}, resp.Code)
			return
		}

		lib.WriteResponse(w, map[string]interface{}{
			"token": resp.Token,
		}, resp.Code)
	} else {
		lib.WriteResponse(w, doubleAuthSent.Content, doubleAuthSent.Code)
	}

}

func getDoubleAuth(accountID string, w http.ResponseWriter, req *http.Request) CheckDoubleAuthResponse {
	device := utils.GetCurrentUserDevice(w, req, accountID)

	patientInfo, err := graphql.GetPatientById(accountID)
	if err == nil && patientInfo.DoubleAuthMethodsID != nil && *patientInfo.DoubleAuthMethodsID != "" {
		response, err := graphql.GetDoubleAuthById(*patientInfo.DoubleAuthMethodsID)
		if err != nil {
			return CheckDoubleAuthResponse{
				Content: map[string]interface{}{"message": "Unable to fetch double authentication methods."},
				Code:    http.StatusNotFound,
				Err:     err,
			}
		}
		return CheckDoubleAuthResponse{
			Content: map[string]interface{}{"2fa_methods": response.Methods, "device": device},
			Code:    200,
			Err:     nil,
		}
	}

	doctorInfo, err := graphql.GetDoctorById(accountID)
	if err == nil && doctorInfo.DoubleAuthMethodsID != nil && *doctorInfo.DoubleAuthMethodsID != "" {
		response, err := graphql.GetDoubleAuthById(*doctorInfo.DoubleAuthMethodsID)
		if err != nil {
			return CheckDoubleAuthResponse{
				Content: map[string]interface{}{"message": "Unable to fetch double authentication methods."},
				Code:    http.StatusNotFound,
				Err:     err,
			}
		}
		return CheckDoubleAuthResponse{
			Content: map[string]interface{}{"2fa_methods": response.Methods, "device": device},
			Code:    200,
			Err:     nil,
		}
	}

	return CheckDoubleAuthResponse{
		Content: map[string]interface{}{"message": "No double authentication methods found."},
		Code:    http.StatusNotFound,
		Err:     nil,
	}
}
