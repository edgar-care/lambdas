package handlers

import (
	"encoding/json"
	edgarlib "github.com/edgar-care/edgarlib/auth"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
)

type Enable_account struct {
	EnableAccount bool `json:"enable_account"`
}

func EnableAccount(w http.ResponseWriter, req *http.Request) {
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

	enable_account := edgarlib.ModifyStatusAccount(patientID, true)

	if enable_account.Err != nil {
		lib.WriteError(w, enable_account.Code, enable_account.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"account_status": true,
	}, 201)
}
