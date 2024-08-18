package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
)

func GetDevice(w http.ResponseWriter, req *http.Request) {

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

	var devices edgarlib.GetDevicesConnectResponse
	page := req.URL.Query().Get("page")
	size := req.URL.Query().Get("size")
	if page == "" && size == "" {
		devices = edgarlib.GetDeviceConnect(ownerID.ID, 0, 0)
	} else {
		number_page, err1 := strconv.Atoi(page)
		number_size, err2 := strconv.Atoi(size)
		if err1 != nil || err2 != nil {
			devices = edgarlib.GetDeviceConnect(ownerID.ID, 0, 0)
		} else {
			devices = edgarlib.GetDeviceConnect(ownerID.ID, number_page, number_size)
		}
	}
	if devices.Err != nil {
		lib.WriteError(w, devices.Code, devices.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"devices": devices.DevicesConnect,
	}, devices.Code)
}
