package handlers

import (
	"net/http"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	"github.com/edgar-care/dashboard/cmd/main/services"
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

	patient, err := services.GetPatientById(t)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id not correspond to a patient",
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"patient": patient,
	}, 201)
}

func GetPatients(w http.ResponseWriter, req *http.Request) {

	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	patient, err := services.GetAllPatientDoctor(doctorID)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"all patient": patient,
	}, 201)
}
