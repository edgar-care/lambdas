package handlers

import (
	"context"
	"encoding/json"
	"github.com/edgar-care/edgarlib/graphql"
	"net/http"

	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/medical_folder"
)

func GetMedicalInformation(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	medicalInfo := edgarlib.GetMedicalInfosById(patientID)
	if medicalInfo.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medicalInfo.Err.Error(),
		}, medicalInfo.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"medical_folder": medicalInfo.MedicalInfo,
	}, 200)
}

func ModifyFolderMedical(w http.ResponseWriter, req *http.Request) {

	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	//t := chi.URLParam(req, "id")
	t, err := graphql.GetPatientById(context.Background(), graphql.CreateClient(), patientID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input edgarlib.CreateMedicalInfoInput
	err = json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	medicalFolder := edgarlib.UpdateMedicalFolder(input, t.GetPatientById.Medical_info_id)
	if medicalFolder.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": medicalFolder.Err.Error(),
		}, medicalFolder.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"medifcal_folder": medicalFolder.MedicalInfo,
	}, 200)
}
