package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/double_auth"
)

func AddTrustDevice(w http.ResponseWriter, req *http.Request) {
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

	new_trust_device := edgarlib.AddTrustDevice(t, ownerID.ID)

	if new_trust_device.Err != nil {
		lib.WriteError(w, new_trust_device.Code, new_trust_device.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"trusted_device": new_trust_device.Patient.TrustDevices,
	}, 201)
}
