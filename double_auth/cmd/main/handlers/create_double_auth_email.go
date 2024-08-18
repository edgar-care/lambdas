package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
)

func AddDoubleAutEmail(w http.ResponseWriter, req *http.Request) {

	ownerID := authlib.AuthMiddlewareAccount(w, req)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input edgarlib.CreateDoubleAuthInput

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	doubleAuthResponse := edgarlib.CreateDoubleAuthEmail(input, ownerID)

	if doubleAuthResponse.Err != nil {
		lib.WriteError(w, doubleAuthResponse.Code, doubleAuthResponse.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"double_auth": doubleAuthResponse.DoubleAuth,
	}, 201)
}
