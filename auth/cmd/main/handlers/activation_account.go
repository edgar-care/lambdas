package handlers

import (
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/edgar-care/auth/cmd/main/lib"
)

func EnableAccount(w http.ResponseWriter, req *http.Request) {
	token := edgarlib.GetBearerToken(req)
	accountID, _ := edgarlib.GetAuthenticatedAccount(token)
	if accountID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	enable_account := edgarlib.ModifyStatusAccount(accountID, true)

	if enable_account.Err != nil {
		lib.WriteError(w, enable_account.Code, enable_account.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"account_status": true,
	}, 201)
}
