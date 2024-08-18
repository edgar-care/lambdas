package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
)

func GetDevice(w http.ResponseWriter, req *http.Request) {

	ownerID := authlib.AuthMiddlewareAccount(w, req)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	t := chi.URLParam(req, "id")

	device := edgarlib.GetDeviceConnectById(t)
	if device.Err != nil {
		lib.WriteError(w, device.Code, device.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"double_auth": device.DeviceConnect,
	}, device.Code)
}

func GetDevices(w http.ResponseWriter, req *http.Request) {

	ownerID := authlib.AuthMiddlewareAccount(w, req)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	devices := edgarlib.GetDeviceConnect(ownerID)
	if devices.Err != nil {
		lib.WriteError(w, devices.Code, devices.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"devices": devices.DevicesConnect,
	}, devices.Code)
}
