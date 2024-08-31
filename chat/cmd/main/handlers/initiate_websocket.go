package handlers

import (
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/http"
	"github.com/edgar-care/edgarlib/v2/redis"
	"net/http"
)

type ReadyInput struct {
	ConnectionId string       `json:"connectionId"`
	Payload      PayloadReady `json:"payload"`
}

type PayloadReady struct {
	AuthToken string `json:"authToken"`
	DeviceID  string `json:"deviceId"`
}

type DisconnectInput struct {
	ConnectionId string `json:"connectionId"`
}

func Connection(w http.ResponseWriter, req *http.Request) {
	response := map[string]interface{}{
		"message": "Connected",
	}

	lib.WriteResponse(w, response, 200)
}

func Ready(w http.ResponseWriter, req *http.Request) {

	var input ReadyInput

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

	if accountID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	_, err = redis.SetKey(input.Payload.DeviceID, input.ConnectionId, nil)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	_, err = redis.SetKey(input.ConnectionId, input.Payload.DeviceID, nil)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	response := map[string]interface{}{
		"message": "Connection ready",
	}

	lib.WriteResponse(w, response, 200)
}

func Disconnect(w http.ResponseWriter, req *http.Request) {

	var input DisconnectInput
	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	deviceId, err := redis.GetKey(input.ConnectionId)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	_, err = redis.DeleteKey(input.ConnectionId)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}
	_, err = redis.DeleteKey(deviceId)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	edgarlib.WriteResponse(w, "Connection disconnected", 200)
}
