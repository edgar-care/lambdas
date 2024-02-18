package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/appointment"
)

type RdvInput struct {
	IDPatient string `json:"id_patient"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

func CreateRdv(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input RdvInput

	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	rdv := edgarlib.CreateRdv(input.IDPatient, doctorID, input.StartDate, input.EndDate)

	if rdv.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": rdv.Err.Error(),
		}, rdv.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": rdv.Rdv,
	}, 201)
}
