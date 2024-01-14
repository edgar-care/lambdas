package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/appointments/cmd/main/lib"
	"github.com/edgar-care/appointments/cmd/main/services"
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
	_, err = services.UpdateRdv("", idAppointment, &input.Reason)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"reason": input.Reason,
	}, 200)
}
