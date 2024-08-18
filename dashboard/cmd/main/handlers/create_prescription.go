package handlers

import (
	"encoding/json"
	"github.com/edgar-care/dashboard/cmd/main/lib"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/ordonnance"
	"net/http"
)

func CreatePrescription(w http.ResponseWriter, req *http.Request) {
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

	var input edgarlib.CreateOrdonnaceInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	ordonnance := edgarlib.CreateOrdonnance(input, doctorID.ID)
	if ordonnance.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": ordonnance.Err.Error(),
		}, ordonnance.Code)
		return
	}
	lib.WriteResponse(w, map[string]interface{}{
		"prescription":     ordonnance.Ordonnance,
		"url_prescription": ordonnance.Url,
	}, 201)
}
