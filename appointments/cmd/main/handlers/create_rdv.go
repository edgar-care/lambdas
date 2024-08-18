package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/appointment"
)

type RdvInput struct {
	StartDate int `json:"start_date"`
	EndDate   int `json:"end_date"`
}

func CreateRdv(w http.ResponseWriter, req *http.Request) {
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
	var input RdvInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	rdv := edgarlib.CreateRdv("", doctorID.ID, input.StartDate, input.EndDate, "")

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
