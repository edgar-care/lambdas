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

	if patient.OnboardingInfoID == "" || patient.OnboardingHealthID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Onboarding not started",
		}, 400)
		return
	}

	info, err := services.GetInfoById(patient.OnboardingInfoID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id does not correspond to an info part",
		}, 400)
		return
	}
	health, err := services.GetHealthById(patient.OnboardingHealthID)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id does not correspond to a health part",
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"patient":           patient,
		"onboarding_info":   info,
		"onboarding_health": health,
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
		"patients": patient,
	}, 201)
}
