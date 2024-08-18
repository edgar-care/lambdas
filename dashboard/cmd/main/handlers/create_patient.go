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

	patient := edgarlib.CreatePatientFormDoctor(input, doctorID.ID)
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

	lib.WriteResponse(w, response, patient.Code)
}
