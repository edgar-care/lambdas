package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/double_auth"
)

func AddDoubleAutEmail(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
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

	doubleAuthResponse := edgarlib.CreateDoubleAuthEmail(input, patientID)

	if doubleAuthResponse.Err != nil {
		lib.WriteError(w, doubleAuthResponse.Code, doubleAuthResponse.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"double_auth": doubleAuthResponse.DoubleAuth,
	}, 201)
}
