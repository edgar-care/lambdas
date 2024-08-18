package handlers

import (
	"github.com/edgar-care/auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"
)

type CreateSaveCodeResponse struct {
	SaveCode model.SaveCode
	Code     int
	Err      error
}

func CreateBackupCode(w http.ResponseWriter, req *http.Request) {

	accountID := edgarlib.AuthMiddlewareAccount(w, req)
	if accountID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	backupCode := edgarlib.CreateBackupCodes(accountID, req)

	if backupCode.Err != nil {
		lib.WriteError(w, backupCode.Code, backupCode.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"double_auth": backupCode.SaveCode.Code,
	}, backupCode.Code)
}
