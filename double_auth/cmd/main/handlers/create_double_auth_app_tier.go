package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
)

func AddDoubleAuthAppTier(w http.ResponseWriter, req *http.Request) {

	ownerID := authlib.AuthMiddlewareAccount(w, req)
	if ownerID.Code == 409 || ownerID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": ownerID.Err.Error(),
		}, ownerID.Code)
		return
	}
	if ownerID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	doubleAuthResponse := edgarlib.CreateDoubleAuthAppTier(ownerID.ID)
	if doubleAuthResponse.Err != nil {
		lib.WriteError(w, doubleAuthResponse.Code, doubleAuthResponse.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"base32":      doubleAuthResponse.TotpInfo.Secret,
		"otpauth_url": doubleAuthResponse.TotpInfo.Url,
	}, doubleAuthResponse.Code)
}
