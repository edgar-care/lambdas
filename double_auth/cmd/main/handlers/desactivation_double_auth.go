package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
)

func DisableDoubleAuth(w http.ResponseWriter, req *http.Request) {

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
	double_auth := chi.URLParam(req, "ENUM")

	if double_auth == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "ENUM is not defined",
		}, 400)
		return
	}

	validEnums := map[string]bool{
		"MOBILE":          true,
		"EMAIL":           true,
		"AUTHENTIFICATOR": true,
		"BACKUPCODE":      true,
	}

	if !validEnums[double_auth] {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid ENUM value",
		}, 400)
		return
	}

	deletedDevice := edgarlib.RemoveDoubleAuthMethod(double_auth, ownerID.ID)
	if deletedDevice != nil {
		lib.WriteResponse(w, map[string]string{
			"message": deletedDevice.Error(),
		}, 200)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"message": "Double auth deleted successfully",
	}, http.StatusOK)
}
