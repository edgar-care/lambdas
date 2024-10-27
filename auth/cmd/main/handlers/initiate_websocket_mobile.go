package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	"github.com/edgar-care/edgarlib/v2/redis"
	"github.com/google/uuid"
	"net/http"
)

type ReadyInput struct {
	ConnectionId string `json:"connectionId"`
	Os           string `json:"os"`
	Browser      string `json:"browser"`
	Location     string `json:"location"`
}

type DisconnectInput struct {
	ConnectionId string `json:"connectionId"`
}

func ReadyLoginWeb(w http.ResponseWriter, req *http.Request) {
	var input ReadyInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	uuidLogin := uuid.New()

	expire := 600
	_, err = redis.SetKey(uuidLogin.String(), input.ConnectionId, &expire)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	_, err = redis.SetKey(input.ConnectionId, uuidLogin.String(), &expire)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	//stock OS + browser + location dans un redis
	_, err = redis.StoreUserInfoHash(uuidLogin.String()+":info", input.Os, input.Browser, input.Location, &expire)
	if err != nil {
		lib.WriteResponse(w, map[string]string{"message": err.Error()}, 500)
	}

	response := map[string]interface{}{
		"message": "Connection ready",
		"uuid":    uuidLogin.String(),
	}

	lib.WriteResponse(w, response, 200)
}
