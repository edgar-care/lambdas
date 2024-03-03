package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/appointment"
	"github.com/go-chi/chi/v5"
)

type UpdateRdvInput struct {
	ID string `json:"id"`
}

func ModifRdv(w http.ResponseWriter, req *http.Request) {
	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	id_appointment := chi.URLParam(req, "id")

	if id_appointment == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "ID is not defined",
		}, 400)
		return
	}

	// ======================================================= //
	var new_appointment UpdateRdvInput
	err := json.NewDecoder(req.Body).Decode(&new_appointment)

	lib.CheckError(err)
	appointment := edgarlib.EditRdv(new_appointment.ID, id_appointment, patientID)

	if appointment.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": appointment.Err.Error(),
		}, appointment.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": appointment.Rdv,
	}, 201)
}
