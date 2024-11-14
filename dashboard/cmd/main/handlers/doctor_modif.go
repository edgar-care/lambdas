package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	lib "github.com/edgar-care/dashboard/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	model "github.com/edgar-care/edgarlib/v2/graphql/model"
	edgarlib "github.com/edgar-care/edgarlib/v2/medical_folder"
)

func ModifyMedicalInfo(w http.ResponseWriter, req *http.Request) {

	doctorID := authlib.AuthMiddlewareDoctor(w, req)
	if doctorID.Code == 409 || doctorID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": doctorID.Err.Error(),
		}, doctorID.Code)
		return
	}
	if doctorID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	patientId := chi.URLParam(req, "id")

	var input model.UpdateMedicalFolderInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	medicalInfo := edgarlib.UpdateMedicalFolderPatient(patientId, input)
	if medicalInfo.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medicalInfo.Err.Error(),
		}, medicalInfo.Code)
		return
	}

	response := map[string]interface{}{
		"medical_folder": map[string]interface{}{
			"id":                         medicalInfo.MedicalFolder.ID,
			"name":                       medicalInfo.MedicalFolder.Name,
			"firstname":                  medicalInfo.MedicalFolder.Firstname,
			"birthdate":                  medicalInfo.MedicalFolder.Birthdate,
			"sex":                        medicalInfo.MedicalFolder.Sex,
			"height":                     medicalInfo.MedicalFolder.Height,
			"weight":                     medicalInfo.MedicalFolder.Weight,
			"primary_doctor_id":          medicalInfo.MedicalFolder.PrimaryDoctorID,
			"family_members_med_info_id": medicalInfo.MedicalFolder.FamilyMembersMedInfoID,
			"onboarding_status":          medicalInfo.MedicalFolder.OnboardingStatus,
			"medical_antecedents":        medicalInfo.MedicalFolder.AntecedentDiseaseIds,
		},
	}

	lib.WriteResponse(w, response, medicalInfo.Code)
}
