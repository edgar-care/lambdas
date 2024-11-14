package handlers

import (
	"encoding/json"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"

	"github.com/edgar-care/edgarlib/v2/graphql"

	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/medical_folder"
)

func GetMedicalInformation(w http.ResponseWriter, req *http.Request) {

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

	medicalInfo := edgarlib.GetMedicalFolder(patientID.ID)
	if medicalInfo.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medicalInfo.Err.Error(),
		}, medicalInfo.Code)
		return
	}

	response := map[string]interface{}{
		"medical_folder": map[string]interface{}{
			"id":                         medicalInfo.MedicalInfo.ID,
			"name":                       medicalInfo.MedicalInfo.Name,
			"firstname":                  medicalInfo.MedicalInfo.Firstname,
			"birthdate":                  medicalInfo.MedicalInfo.Birthdate,
			"sex":                        medicalInfo.MedicalInfo.Sex,
			"height":                     medicalInfo.MedicalInfo.Height,
			"weight":                     medicalInfo.MedicalInfo.Weight,
			"primary_doctor_id":          medicalInfo.MedicalInfo.PrimaryDoctorID,
			"family_members_med_info_id": medicalInfo.MedicalInfo.FamilyMembersMedInfoID,
			"onboarding_status":          medicalInfo.MedicalInfo.OnboardingStatus,
			"medical_antecedents":        medicalInfo.MedicalAntecedents,
		},
	}
	lib.WriteResponse(w, response, medicalInfo.Code)

}

func ModifyFolderMedical(w http.ResponseWriter, req *http.Request) {

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

	t, err := graphql.GetPatientById(patientID.ID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input model.UpdateMedicalFolderInput
	err = json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	if t.MedicalInfoID == nil {
		lib.WriteResponse(w, map[string]string{"message": "medical folder not found"}, 404)
	}
	medicalFolder := edgarlib.UpdateMedicalFolderPatient(patientID.ID, input)
	if medicalFolder.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medicalFolder.Err.Error(),
		}, medicalFolder.Code)
		return
	}

	response := map[string]interface{}{
		"medical_folder": map[string]interface{}{
			"id":                         medicalFolder.MedicalFolder.ID,
			"name":                       medicalFolder.MedicalFolder.Name,
			"firstname":                  medicalFolder.MedicalFolder.Firstname,
			"birthdate":                  medicalFolder.MedicalFolder.Birthdate,
			"sex":                        medicalFolder.MedicalFolder.Sex,
			"height":                     medicalFolder.MedicalFolder.Height,
			"weight":                     medicalFolder.MedicalFolder.Weight,
			"primary_doctor_id":          medicalFolder.MedicalFolder.PrimaryDoctorID,
			"family_members_med_info_id": medicalFolder.MedicalFolder.FamilyMembersMedInfoID,
			"onboarding_status":          medicalFolder.MedicalFolder.OnboardingStatus,
			"medical_antecedents":        medicalFolder.MedicalFolder.AntecedentDiseaseIds,
		},
	}

	lib.WriteResponse(w, response, medicalFolder.Code)
}
