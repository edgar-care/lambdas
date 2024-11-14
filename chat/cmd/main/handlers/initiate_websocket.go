package handlers

import (
	"encoding/base64"
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/http"
	"github.com/edgar-care/edgarlib/v2/redis"
	"net/http"
	"strings"
)

type ReadyInput struct {
	ConnectionId string       `json:"connectionId"`
	Payload      PayloadReady `json:"payload"`
}

type PayloadReady struct {
	AuthToken string `json:"authToken"`
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
	if check_account.Code == 409 || check_account.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": check_account.Err.Error(),
		}, check_account.Code)
		return
	}

	if accountID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	parts := strings.Split(input.Payload.AuthToken, ".")
	if len(parts) != 3 {
		return
	}

	decodedBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		lib.CheckError(err)
		return
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(decodedBytes, &jsonMap); err != nil {
		lib.CheckError(err)
		return
	}

	_, err = redis.SetKey(jsonMap["name_device"].(string), input.ConnectionId, nil)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	_, err = redis.SetKey(input.ConnectionId, jsonMap["name_device"].(string), nil)
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
