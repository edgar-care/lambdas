package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/dashboard"
	"github.com/go-chi/chi/v5"
)

func GetPatientId(w http.ResponseWriter, req *http.Request) {
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

	t := chi.URLParam(req, "id")

	patient := edgarlib.GetPatientById(t, doctorID.ID)
	if patient.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": patient.Err.Error(),
		}, patient.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"id":                      patient.PatientInfo.ID,
		"email":                   patient.PatientInfo.Email,
		"treatment_follow_up_ids": patient.PatientInfo.TreatmentFollowUp,
		"document_ids":            patient.PatientInfo.DocumentsIds,
		"rendez_vous_ids":         patient.PatientInfo.RendezVousIds,
		"medical_folder": map[string]interface{}{
			"id":                         patient.PatientInfo.MedicalInfo.ID,
			"name":                       patient.PatientInfo.MedicalInfo.Name,
			"firstname":                  patient.PatientInfo.MedicalInfo.Firstname,
			"birthdate":                  patient.PatientInfo.MedicalInfo.Birthdate,
			"sex":                        patient.PatientInfo.MedicalInfo.Sex,
			"height":                     patient.PatientInfo.MedicalInfo.Height,
			"weight":                     patient.PatientInfo.MedicalInfo.Weight,
			"primary_doctor_id":          patient.PatientInfo.MedicalInfo.PrimaryDoctorID,
			"family_members_med_info_id": patient.PatientInfo.MedicalInfo.FamilyMembersMedInfoID,
			"onboarding_status":          patient.PatientInfo.MedicalInfo.OnboardingStatus,
			"medical_antecedents":        patient.PatientInfo.Antedisease,
		},
	}, 200)
}

func GetPatients(w http.ResponseWriter, req *http.Request) {

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

	patient := edgarlib.GetPatients(doctorID.ID)

	if patient.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": patient.Err.Error(),
		}, 400)
		return
	}

	var patientsResponse []map[string]interface{}

	for _, patientInfo := range patient.PatientsInfo {
		patientResponse := map[string]interface{}{
			"id":                      patientInfo.ID,
			"email":                   patientInfo.Email,
			"treatment_follow_up_ids": patientInfo.TreatmentFollowUp,
			"document_ids":            patientInfo.DocumentsIds,
			"rendez_vous_ids":         patientInfo.RendezVousIds,
			"medical_folder": map[string]interface{}{
				"id":                         patientInfo.MedicalInfo.ID,
				"name":                       patientInfo.MedicalInfo.Name,
				"firstname":                  patientInfo.MedicalInfo.Firstname,
				"birthdate":                  patientInfo.MedicalInfo.Birthdate,
				"sex":                        patientInfo.MedicalInfo.Sex,
				"height":                     patientInfo.MedicalInfo.Height,
				"weight":                     patientInfo.MedicalInfo.Weight,
				"primary_doctor_id":          patientInfo.MedicalInfo.PrimaryDoctorID,
				"family_members_med_info_id": patientInfo.MedicalInfo.FamilyMembersMedInfoID,
				"onboarding_status":          patientInfo.MedicalInfo.OnboardingStatus,
				"medical_antecedents":        patientInfo.Antedisease,
			},
		}

		patientsResponse = append(patientsResponse, patientResponse)
	}
	lib.WriteResponse(w, map[string]interface{}{"patients": patientsResponse}, 200)

}
