package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/medical_folder"
)

func AddMedicalInfo(w http.ResponseWriter, req *http.Request) {
	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input edgarlib.CreateMedicalInfoInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	medical := edgarlib.CreateMedicalInfo(input, patientID)
	if medical.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medical.Err.Error(),
		}, medical.Code)
		return
	}

	response := map[string]interface{}{
		"medical_folder": map[string]interface{}{
			"id":                medical.MedicalInfo.ID,
			"name":              medical.MedicalInfo.Name,
			"firstname":         medical.MedicalInfo.Firstname,
			"birthdate":         medical.MedicalInfo.Birthdate,
			"sex":               medical.MedicalInfo.Sex,
			"height":            medical.MedicalInfo.Height,
			"weight":            medical.MedicalInfo.Weight,
			"primary_doctor_id": medical.MedicalInfo.PrimaryDoctorID,
			"onboarding_status": medical.MedicalInfo.OnboardingStatus,
			"medical_antecedents": func() []map[string]interface{} {
				// Convert antecedent diseases to the desired format
				var diseases []map[string]interface{}
				for _, disease := range medical.AnteDiseasesWithTreatments {
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

	lib.WriteResponse(w, response, medical.Code)
}
