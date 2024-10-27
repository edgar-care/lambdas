package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/edgar-care/auth/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/double_auth"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/redis"
	"math/rand"
	"net/http"
)

type AskMobileConnectionInput struct {
	ConnectionId string                     `json:"connectionId"`
	Payload      PayloadAskMobileConnection `json:"payload"`
}

type PayloadAskMobileConnection struct {
	UUID     string `json:"uuid"`
	Type     string `json:"type"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseMobileConnectionInput struct {
	ConnectionId string                          `json:"connectionId"`
	Payload      PayloadResponseMobileConnection `json:"payload"`
}

type PayloadResponseMobileConnection struct {
	AuthToken string `json:"authToken"`
	UUID      string `json:"uuid"`
	Response  bool   `json:"response"`
}

func AskMobileConnection(w http.ResponseWriter, req *http.Request) {
	var input AskMobileConnectionInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	resp := authlib.Login(authlib.LoginInput{
		Email:    input.Payload.Email,
		Password: input.Payload.Password,
	}, input.Payload.Type, "tempLogin")
	if resp.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": resp.Err.Error(),
		}, resp.Code)
		return
	}
	accountID, _ := authlib.GetAuthenticatedAccount(resp.Token)

	infoKey := input.Payload.UUID + ":info"
	infoDevice, err := redis.GetUserInfoHash(infoKey)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
		return
	}

	data := map[string]interface{}{
		"action":   "ask_mobile_connection",
		"uuid":     input.Payload.UUID,
		"os":       infoDevice.OS,
		"browser":  infoDevice.Browser,
		"location": infoDevice.Location,
	}

	trustDevices := double_auth.GetTrustDeviceConnect(accountID)

	var trustDevicesId []string
	for _, trustDevice := range trustDevices.DevicesConnect {
		trustDevicesId = append(trustDevicesId, trustDevice.ID)
	}

	lib.BroadcastMessage(w, trustDevicesId, data)

	response := map[string]interface{}{
		"message": "Message sent",
	}

	lib.WriteResponse(w, response, 200)
}

func ResponseMobileConnection(w http.ResponseWriter, req *http.Request) {
	var input ResponseMobileConnectionInput
	var email string
	var code string

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	accountID := lib.AuthMiddleware(input.Payload.AuthToken)
	check_account := authlib.CheckAccountEnable(accountID)
	if check_account.Code == 409 {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authorized, this account is disable",
		}, 409)
		return
	}

	connectionId, err := redis.GetKey(input.Payload.UUID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
		return
	}

	if connectionId == "" {
		lib.WriteResponse(w, map[string]string{"message": "The connection has expired"}, 403)
	}

	if input.Payload.Response {
		accountId, accountType := authlib.GetAuthenticatedAccount(input.Payload.AuthToken)

		if accountType == "patient" {
			patient, err := graphql.GetPatientById(accountId)
			if err != nil {
				lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
				return
			}
			email = patient.Email
		} else {
			doctor, err := graphql.GetDoctorById(accountId)
			if err != nil {
				lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
				return
			}
			email = doctor.Email
		}

		code = fmt.Sprintf("%06d", rand.Intn(1000000))

		expire := 600
		_, err = redis.SetKey(email, code, &expire)
		if err != nil {
			lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
			return
		}
	}

	data := map[string]interface{}{
		"action":   "response_mobile_connection",
		"code":     code,
		"response": input.Payload.Response,
	}

	lib.BroadcastMessage(w, []string{input.Payload.UUID}, data)
	lib.DisconnectToConnection(w, connectionId)

	response := map[string]interface{}{
		"message": "Message sent",
	}

	lib.WriteResponse(w, response, 200)
}
