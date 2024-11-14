package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/medical_folder"
)

func AddMedicalInfo(w http.ResponseWriter, req *http.Request) {
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

	var input edgarlib.CreateNewMedicalInfoInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	medical := edgarlib.NewMedicalFolder(input, patientID.ID)
	if medical.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medical.Err.Error(),
		}, medical.Code)
		return
	}

	response := map[string]interface{}{
		"medical_folder": map[string]interface{}{
			"id":                         medical.MedicalInfo.ID,
			"name":                       medical.MedicalInfo.Name,
			"firstname":                  medical.MedicalInfo.Firstname,
			"birthdate":                  medical.MedicalInfo.Birthdate,
			"sex":                        medical.MedicalInfo.Sex,
			"height":                     medical.MedicalInfo.Height,
			"weight":                     medical.MedicalInfo.Weight,
			"primary_doctor_id":          medical.MedicalInfo.PrimaryDoctorID,
			"family_members_med_info_id": medical.MedicalInfo.FamilyMembersMedInfoID,
			"onboarding_status":          medical.MedicalInfo.OnboardingStatus,
			"medical_antecedents":        medical.MedicalAntecedents,
		},
	}

	lib.WriteResponse(w, response, medical.Code)
}
