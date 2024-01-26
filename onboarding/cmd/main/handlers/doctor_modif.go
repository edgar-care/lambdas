package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/onboarding/cmd/main/lib"
	"github.com/edgar-care/onboarding/cmd/main/services"
)

func ModifyMedicalInfo(w http.ResponseWriter, req *http.Request) {

	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	patientId := chi.URLParam(req, "id")

	patient, err := services.GetPatientById(patientId)
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

	var input services.MedicalInfo
	err = json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)

	info, err := services.UpdateMedicalInfo(patient.OnboardingInfoID, input.Info)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get Info (check if you share all information) " + err.Error(),
		}, 400)
		return
	}

	health, err := services.UpdateMedicalHealth(patient.OnboardingHealthID, input.Health)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get Health (check if you share all information) " + err.Error(),
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"patient_health": health,
		"patient_info":   info,
	}, 200)
}
