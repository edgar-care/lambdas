package handlers

import (
	"encoding/json"
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

	medicalInfo := edgarlib.GetMedicalInfo(patientID.ID)
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

	var input edgarlib.UpdateMedicalInfoInput
	err = json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	if t.MedicalInfoID == nil {
		lib.WriteResponse(w, map[string]string{"message": "medical folder not found"}, 404)
	}
	medicalFolder := edgarlib.UpdateMedicalFolder(input, *t.MedicalInfoID)
	if medicalFolder.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medicalFolder.Err.Error(),
		}, medicalFolder.Code)
		return
	}

	response := map[string]interface{}{
		"medical_folder": map[string]interface{}{
			"id":                         medicalFolder.MedicalInfo.ID,
			"name":                       medicalFolder.MedicalInfo.Name,
			"firstname":                  medicalFolder.MedicalInfo.Firstname,
			"birthdate":                  medicalFolder.MedicalInfo.Birthdate,
			"sex":                        medicalFolder.MedicalInfo.Sex,
			"height":                     medicalFolder.MedicalInfo.Height,
			"weight":                     medicalFolder.MedicalInfo.Weight,
			"primary_doctor_id":          medicalFolder.MedicalInfo.PrimaryDoctorID,
			"family_members_med_info_id": medicalFolder.MedicalInfo.FamilyMembersMedInfoID,
			"onboarding_status":          medicalFolder.MedicalInfo.OnboardingStatus,
			"medical_antecedents": func() []map[string]interface{} {
				// Convert antecedent diseases to the desired format
				var diseases []map[string]interface{}
				for _, disease := range medicalFolder.AnteDiseasesWithTreatments {
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

	lib.WriteResponse(w, response, medicalFolder.Code)
}
