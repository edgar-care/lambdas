package handlers

import (
	"context"
	edgargraph "github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	edgarlib "github.com/edgar-care/edgarlib/treatment"
	"github.com/edgar-care/treatment/cmd/main/lib"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func GetTreatment(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	t := chi.URLParam(req, "id")

	treatment := edgarlib.GetTreatmentById(t, patientID)
	if treatment.Err != nil {
		lib.WriteError(w, treatment.Code, treatment.Err.Error())
		return
	}

	response := map[string]interface{}{
		"treatment": map[string]interface{}{
			"name":           treatment.Antedisease.Name,
			"still_relevant": treatment.Antedisease.StillRelevant,
			"treatment":      treatment.Treatment,
		},
	}

	lib.WriteResponse(w, response, treatment.Code)
}

func GetTreatments(w http.ResponseWriter, req *http.Request) {
	gqlClient := edgargraph.CreateClient()
	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	treatments := edgarlib.GetTreatments(patientID)
	if treatments.Err != nil {
		lib.WriteError(w, treatments.Code, treatments.Err.Error())
		return
	}

	treatmentMap := make(map[string][]model.Treatment)

	for _, antedisease := range treatments.Antedisease {
		associatedAnte, err := edgargraph.GetAnteDiseaseByID(context.Background(), gqlClient, antedisease.ID)
		if err != nil {
			log.Println("Failed to retrieve associated treatments for antedisease:", err)
			continue
		}

		for _, treatmentID := range associatedAnte.GetAnteDiseaseByID.Treatment_ids {
			associatedTreatment := edgarlib.GetTreatmentById(treatmentID, patientID)
			if associatedTreatment.Err != nil {
				log.Println("Failed to retrieve associated treatment with ID:", treatmentID, ":", associatedTreatment.Err)
				continue
			}

			treatmentMap[antedisease.ID] = append(treatmentMap[antedisease.ID], associatedTreatment.Treatment)
		}
	}

	response := make([]map[string]interface{}, 0)
	for _, antedisease := range treatments.Antedisease {
		treatmentData := map[string]interface{}{
			"antedisease": antedisease,
			"treatments":  treatmentMap[antedisease.ID],
		}
		response = append(response, treatmentData)
	}

	lib.WriteResponse(w, response, treatments.Code)
}
