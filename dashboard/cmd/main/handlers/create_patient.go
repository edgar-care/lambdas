package handlers

import (
	"encoding/json"
	"github.com/edgar-care/dashboard/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/dashboard"
	"net/http"
)

func CreatePatient(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	var input edgarlib.CreatePatientInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	patient := edgarlib.CreatePatientFormDoctor(input, doctorID)
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
			"id":                patient.MedicalInfo.ID,
			"name":              patient.MedicalInfo.Name,
			"firstname":         patient.MedicalInfo.Firstname,
			"birthdate":         patient.MedicalInfo.Birthdate,
			"sex":               patient.MedicalInfo.Sex,
			"height":            patient.MedicalInfo.Height,
			"weight":            patient.MedicalInfo.Weight,
			"primary_doctor_id": patient.MedicalInfo.PrimaryDoctorID,
			"onboarding_status": patient.MedicalInfo.OnboardingStatus,
			"medical_antecedents": func() []map[string]interface{} {
				// Convert antecedent diseases to the desired format
				var diseases []map[string]interface{}
				for _, disease := range patient.AnteDiseasesWithTreatments {
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

	lib.WriteResponse(w, response, patient.Code)
}
