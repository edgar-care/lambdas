package handlers

import (
	"net/http"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/dashboard"
	"github.com/go-chi/chi/v5"
)

func GetPatientId(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	t := chi.URLParam(req, "id")

	patient := edgarlib.GetPatientById(t, doctorID)
	if patient.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": patient.Err.Error(),
		}, patient.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"patient":      patient.Patient,
		"medical_info": patient.MedicalInfo,
	}, 200)
}

func GetPatients(w http.ResponseWriter, req *http.Request) {

	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	patient := edgarlib.GetPatients(doctorID)

	if patient.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": patient.Err.Error(),
		}, patient.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"patients": patient.PatientsInfo,
	}, 200)
}
