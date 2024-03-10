package handlers

import (
	edgarlib "github.com/edgar-care/edgarlib/dashboard"
	"net/http"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	"github.com/go-chi/chi/v5"
)

func DeletePatientHandler(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	PatientID := chi.URLParam(req, "id")
	patient := edgarlib.DeletePatient(PatientID, doctorID)

	if patient.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": patient.Err.Error(),
		}, patient.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"doctor": patient.UpdatedDoctor,
	}, 201)

}
