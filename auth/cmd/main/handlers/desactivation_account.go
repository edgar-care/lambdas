package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/auth"
)

func DisableAccount(w http.ResponseWriter, req *http.Request) {

	authToken := lib.GetBearerToken(req)
	patientID := lib.AuthMiddleware(authToken)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input Enable_account

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	enable_account := edgarlib.ModifyStatusAccount(patientID, false)

	if enable_account.Err != nil {
		lib.WriteError(w, enable_account.Code, enable_account.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"account_status": false,
	}, 201)
}
