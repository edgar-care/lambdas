package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/appointment"
)

type DeleteRdvInput struct {
	Reason string `json:"reason"`
}

func CancelRdv(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input DeleteRdvInput

	err := json.NewDecoder(req.Body).Decode(&input)
	lib.CheckError(err)

	idAppointment := chi.URLParam(req, "id")
	updateRdv := edgarlib.CancelRdv(idAppointment, input.Reason)

	if updateRdv.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": updateRdv.Err.Error(),
		}, updateRdv.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"reason": updateRdv.Reason,
	}, 200)
}
