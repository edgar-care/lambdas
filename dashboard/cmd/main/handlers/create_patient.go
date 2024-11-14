package handlers

import (
	"encoding/json"
	"github.com/edgar-care/dashboard/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/dashboard"
	"net/http"
)

func CreatePatient(w http.ResponseWriter, req *http.Request) {
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
		}, http.StatusUnauthorized)
		return
	}

	var input edgarlib.CreatePatientInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	patient := edgarlib.CreatePatientFromDoctor(doctorID.ID, input)
	if patient.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": patient.Err.Error(),
		}, patient.Code)
		return
	}

	response := map[string]interface{}{
		"id":      patient.Patient.ID,
		"patient": patient.Patient.Email,
		"medical_folder": map[string]interface{}{
			"id":                         patient.MedicalInfo.ID,
			"name":                       patient.MedicalInfo.Name,
			"firstname":                  patient.MedicalInfo.Firstname,
			"birthdate":                  patient.MedicalInfo.Birthdate,
			"sex":                        patient.MedicalInfo.Sex,
			"height":                     patient.MedicalInfo.Height,
			"weight":                     patient.MedicalInfo.Weight,
			"primary_doctor_id":          patient.MedicalInfo.PrimaryDoctorID,
			"family_members_med_info_id": patient.MedicalInfo.FamilyMembersMedInfoID,
			"onboarding_status":          patient.MedicalInfo.OnboardingStatus,
			"medical_antecedents":        patient.AnteDiseasesWithTreatments,
		},
	}

	lib.WriteResponse(w, response, patient.Code)
}
