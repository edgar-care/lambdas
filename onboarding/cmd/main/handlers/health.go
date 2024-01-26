package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/onboarding/cmd/main/lib"
	"github.com/edgar-care/onboarding/cmd/main/services"
)



func Health(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input services.HealthInput
	var updatePatient services.PatientInput
	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	health, err := services.CreateHealth(input)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get Info Health (check if you share all information) " + err.Error(),
		}, 400)
		return
	}

	updatePatient = services.PatientInput{
		Id: patientID,
		OnboardingHealthID: health.Id,
	}
	patient, err := services.AddOnboardingHealthID(updatePatient)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"health": health,
		"patient": patient,
	}, 201)
}