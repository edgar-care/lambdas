package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/appointment"
)

func DeleteRdv(w http.ResponseWriter, req *http.Request) {
	patientID := authlib.AuthMiddlewarePatient(w, req)
	if patientID.Code == 409 || patientID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": patientID.Err.Error(),
		}, patientID.Code)
		return
	}

	if patientID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	id_appointment := chi.URLParam(req, "id")
	updateRdv := edgarlib.DeleteRdv(id_appointment, patientID.ID)

	if updateRdv.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": updateRdv.Err.Error(),
		}, updateRdv.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"updated_patient": updateRdv.UpdatedPatient,
	}, 201)
}
