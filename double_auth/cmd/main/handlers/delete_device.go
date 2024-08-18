package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
)

func DeleteDevice(w http.ResponseWriter, req *http.Request) {

	ownerID := authlib.AuthMiddlewareAccount(w, req)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	deviceID := chi.URLParam(req, "id")

	if deviceID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "ID is not defined",
		}, 400)
		return
	}

	deletedDevice := edgarlib.DeleteDeviceConnect(deviceID, ownerID)
	if deletedDevice.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": deletedDevice.Err.Error(),
		}, deletedDevice.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"message": "Device deleted successfully",
		"deleted": deletedDevice.Deleted,
	}, http.StatusOK)
}
