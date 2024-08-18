package handlers

import (
	"github.com/edgar-care/double_auth/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
	"net/http"
)

func GetDoubleAuth(w http.ResponseWriter, req *http.Request) {

	ownerID := authlib.AuthMiddlewareAccount(w, req)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	device := edgarlib.GetDoubleAuthById(ownerID)
	if device.Err != nil {
		lib.WriteError(w, device.Code, device.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"double_auth": device.DoubleAuth,
	}, device.Code)
}
