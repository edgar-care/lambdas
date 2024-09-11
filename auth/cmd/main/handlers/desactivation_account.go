package handlers

import (
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
)

func DisableAccount(w http.ResponseWriter, req *http.Request) {

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

	enable_account := edgarlib.ModifyStatusAccount(accountID.ID, false)

	if enable_account.Err != nil {
		lib.WriteError(w, enable_account.Code, enable_account.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"account_status": false,
	}, 201)
}
