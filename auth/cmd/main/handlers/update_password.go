package handlers

import (
	"encoding/json"
	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"
)

func UpdatePassword(w http.ResponseWriter, req *http.Request) {
	accountID := edgarlib.AuthMiddlewareAccount(w, req)
	if accountID.Code == 409 || accountID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": accountID.Err.Error(),
		}, accountID.Code)
		return
	}
	if accountID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input edgarlib.UpdatePasswordInput

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	newPassword := edgarlib.UpdatePassword(input, accountID.ID)

	if newPassword.Err != nil {
		lib.WriteError(w, newPassword.Code, newPassword.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"password": "updated",
	}, newPassword.Code)

}
