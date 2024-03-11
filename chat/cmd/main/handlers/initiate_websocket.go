package handlers

import (
	"encoding/json"
	"github.com/edgar-care/chat/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/http"
	"github.com/edgar-care/edgarlib/redis"
	"net/http"
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

	DoctorID := lib.AuthMiddlewareDoctor(input.Payload.AuthToken)
	PatientID := lib.AuthMiddleware(input.Payload.AuthToken)

	if DoctorID == "" && PatientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	id := DoctorID + PatientID

	_, err = redis.SetKey(id, input.ConnectionId, nil)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	_, err = redis.SetKey(input.ConnectionId, id, nil)
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

	userId, err := redis.GetKey(input.ConnectionId)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	_, err = redis.DeleteKey(input.ConnectionId)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}
	_, err = redis.DeleteKey(userId)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	edgarlib.WriteResponse(w, "Connection disconnected", 200)
}
