package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/onboarding/cmd/main/lib"
	"github.com/edgar-care/onboarding/cmd/main/services"
)



func GetMedicalInformation(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	patient, err := services.GetPatientById(patientID)
	lib.CheckError(err)
	if err != nil {
        lib.WriteResponse(w, map[string]string{
            "message": "Id not correspond to a patient",
        }, 400)
        return
	}

	if patient.OnboardingInfoID == "" {
		lib.WriteResponse(w, map[string]interface{}{
            "patient_health": nil,
			"patient_info": nil,
        }, 400)
		return
	}


	info, err := services.GetInfoById(patient.OnboardingInfoID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get Info (check if you share all information) " + err.Error(),
		}, 400)
		return
	}

	if patient.OnboardingHealthID == "" {
		lib.WriteResponse(w, map[string]interface{}{
            "patient_health": nil,
			"patient_info": info,
        }, 400)
		return
	}

	health, err := services.GetHealthById(patient.OnboardingHealthID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable to get Health (check if you share all information) " + err.Error(),
		}, 400)
		return
	}


	lib.WriteResponse(w, map[string]interface{}{
		"patient_health": health,
		"patient_info": info,
	}, 200)
}

func ModifyFolderMedical(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	patient, err := services.GetPatientById(patientID)
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
		"patient_info": info,
	}, 200)
}