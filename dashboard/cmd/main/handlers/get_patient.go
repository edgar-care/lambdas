package handlers

import (
	"net/http"

	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/dashboard"
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

	patient := edgarlib.GetPatientById(t, doctorID)
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
			"id":                patient.PatientInfo.MedicalInfo.ID,
			"name":              patient.PatientInfo.MedicalInfo.Name,
			"firstname":         patient.PatientInfo.MedicalInfo.Firstname,
			"birthdate":         patient.PatientInfo.MedicalInfo.Birthdate,
			"sex":               patient.PatientInfo.MedicalInfo.Sex,
			"height":            patient.PatientInfo.MedicalInfo.Height,
			"weight":            patient.PatientInfo.MedicalInfo.Weight,
			"primary_doctor_id": patient.PatientInfo.MedicalInfo.PrimaryDoctorID,
			"onboarding_status": patient.PatientInfo.MedicalInfo.OnboardingStatus,
			"medical_antecedents": func() []map[string]interface{} {
				// Convert antecedent diseases to the desired format
				var diseases []map[string]interface{}
				for _, disease := range patient.PatientInfo.Antedisease {
					d := map[string]interface{}{
						"id":   disease.AnteDisease.ID,
						"name": disease.AnteDisease.Name,
						"medicines": func() []map[string]interface{} {
							var medicines []map[string]interface{}
							for _, treatment := range disease.Treatments {
								medicine := map[string]interface{}{
									"id":          treatment.ID,
									"medicine_id": treatment.MedicineID,
									"period":      treatment.Period,
									"day":         treatment.Day,
									"quantity":    treatment.Quantity,
								}
								medicines = append(medicines, medicine)
							}
							return medicines
						}(),
						"still_relevant": disease.AnteDisease.StillRelevant,
					}
					diseases = append(diseases, d)
				}
				return diseases
			}(),
		},
	}, 200)
}

func GetPatients(w http.ResponseWriter, req *http.Request) {

	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	patient := edgarlib.GetPatients(doctorID)

	if patient.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": patient.Err.Error(),
		}, patient.Code)
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
				"id":                patientInfo.MedicalInfo.ID,
				"name":              patientInfo.MedicalInfo.Name,
				"firstname":         patientInfo.MedicalInfo.Firstname,
				"birthdate":         patientInfo.MedicalInfo.Birthdate,
				"sex":               patientInfo.MedicalInfo.Sex,
				"height":            patientInfo.MedicalInfo.Height,
				"weight":            patientInfo.MedicalInfo.Weight,
				"primary_doctor_id": patientInfo.MedicalInfo.PrimaryDoctorID,
				"onboarding_status": patientInfo.MedicalInfo.OnboardingStatus,
				"medical_antecedents": func() []map[string]interface{} {
					var diseases []map[string]interface{}
					for _, disease := range patientInfo.Antedisease {
						d := map[string]interface{}{
							"id":   disease.AnteDisease.ID,
							"name": disease.AnteDisease.Name,
							"medicines": func() []map[string]interface{} {
								var medicines []map[string]interface{}
								for _, treatment := range disease.Treatments {
									medicine := map[string]interface{}{
										"id":          treatment.ID,
										"medicine_id": treatment.MedicineID,
										"period":      treatment.Period,
										"day":         treatment.Day,
										"quantity":    treatment.Quantity,
									}
									medicines = append(medicines, medicine)
								}
								return medicines
							}(),
							"still_relevant": disease.AnteDisease.StillRelevant,
						}
						diseases = append(diseases, d)
					}
					return diseases
				}(),
			},
		}

		patientsResponse = append(patientsResponse, patientResponse)
	}
	lib.WriteResponse(w, map[string]interface{}{"patients": patientsResponse}, 200)

}
