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

type DoubleAuth struct {
	Methods    []string
	DeviceInfo map[string]interface{}
}

type DoubleAuthResponse struct {
	Content DoubleAuth
	Err     error
}

// si patient et médecin avec la même email
func Login(w http.ResponseWriter, req *http.Request) {
	var input authlib.LoginInput
	var accountId string
	var password string
	var doubleAuthId *string
	deviceId := "deviceId"

	t := chi.URLParam(req, "type")

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	if t == "p" {
		patient, err := graphql.GetPatientByEmail(input.Email)
		if err != nil {
			lib.WriteError(w, http.StatusBadRequest, "Email does not correspond to a valid Patient")
			return
		}
		password = patient.Password
		accountId = patient.ID
		doubleAuthId = patient.DoubleAuthMethodsID
	}
	if t == "d" {
		doctor, err := graphql.GetDoctorByEmail(input.Email)
		if err != nil {
			lib.WriteError(w, http.StatusBadRequest, "Email does not correspond to a valid Doctor")
			return
		}
		password = doctor.Password
		accountId = doctor.ID
		doubleAuthId = doctor.DoubleAuthMethodsID
	}
	if t == "a" {
		admin, err := graphql.GetAdminByEmail(input.Email)
		if err != nil {
			lib.WriteError(w, http.StatusBadRequest, "Email does not correspond to a valid Doctor")
			return
		}
		password = admin.Password
		accountId = admin.ID
	}

	passwordCheck := authlib.CheckPassword(input.Password, password)
	if !passwordCheck {
		lib.WriteError(w, http.StatusUnauthorized, "Email and password mismatch")
		return
	}

	if t != "a" {
		if doubleAuthId != nil && *doubleAuthId != "" {
			doubleAuthSent := getDoubleAuth(accountId, *doubleAuthId, w, req)

			if doubleAuthSent.Err != nil {
				lib.WriteResponse(w, "Unable to fetch double authentication methods.", http.StatusNotFound)
				return
			}

			if len(doubleAuthSent.Content.Methods) != 0 {
				lib.WriteResponse(w, doubleAuthSent.Content, http.StatusOK)
				return
			}
		}

		utils.DeviceConnectMiddleware(w, req, accountId)
		deviceId = utils.GetCurrentUserDevice(w, req, accountId).ID
	}

	resp := authlib.Login(input, t, deviceId)

	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"token": resp.Token,
	}, resp.Code)
}

func getDoubleAuth(accountID string, doubleAuthId string, w http.ResponseWriter, req *http.Request) DoubleAuthResponse {
	device := utils.GetCurrentUserDevice(w, req, accountID)
	deviceInfo := map[string]interface{}{
		"os":       device.DeviceType,
		"browser":  device.Browser,
		"location": device.City,
	}

	response, err := graphql.GetDoubleAuthById(doubleAuthId)
	if err != nil {
		return DoubleAuthResponse{
			Err: err,
		}
	}

	return DoubleAuthResponse{
		Content: DoubleAuth{
			Methods:    response.Methods,
			DeviceInfo: deviceInfo,
		},
		Err: nil,
	}
}
