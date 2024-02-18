package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/slot"
)

func CreateSlot(w http.ResponseWriter, req *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input edgarlib.CreateSlotInput

	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	slot := edgarlib.CreateSlot(input, doctorID)
	if slot.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": slot.Err.Error(),
		}, slot.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": slot.Rdv,
	}, 201)
}
