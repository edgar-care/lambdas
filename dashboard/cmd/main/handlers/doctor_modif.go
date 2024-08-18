package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	lib "github.com/edgar-care/dashboard/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
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

	var input edgarlib.UpdateMedicalInfoInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	medicalInfo := edgarlib.UpdateMedicalFolderFromDoctor(input, patientId)
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
			"medical_antecedents": func() []map[string]interface{} {
				// Convert antecedent diseases to the desired format
				var diseases []map[string]interface{}
				for _, disease := range medicalInfo.AnteDiseasesWithTreatments {
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
									"start_date":  treatment.StartDate,
									"end_date":    treatment.EndDate,
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

	lib.WriteResponse(w, response, medicalInfo.Code)
}
