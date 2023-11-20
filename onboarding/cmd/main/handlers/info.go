package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/onboarding/cmd/main/lib"
	"github.com/edgar-care/onboarding/cmd/main/services"
)

func Info(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input services.InfoInput
	var updatePatient services.PatientInput
	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	info, err := services.CreateInfo(input)
	
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	updatePatient = services.PatientInput{
		Id: patientID,
		OnboardingInfoID: info.Id,
	}
	patient, err := services.AddOnboardingInfoID(updatePatient)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}


	lib.WriteResponse(w, map[string]interface{}{
		"info": info,
		"patient": patient,
	}, 201)
}