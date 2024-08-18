package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/treatment"
	"github.com/edgar-care/treatment/cmd/main/lib"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func DeleteTreatment(w http.ResponseWriter, req *http.Request) {

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

	t := chi.URLParam(req, "id")

	treatment := edgarlib.DeleteTreatment(t)

	if treatment.Err != nil {
		lib.WriteError(w, treatment.Code, treatment.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"delete": treatment.Deleted,
	}, treatment.Code)
}
