package handlers

import (
	"github.com/go-chi/chi/v5"
	"net/http"

	"github.com/edgar-care/double_auth/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/double_auth"
)

func DisableDoubleAuth(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	double_auth := chi.URLParam(req, "id")

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

	deletedDevice := edgarlib.RemoveDoubleAuthMethod(double_auth, patientID)
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
