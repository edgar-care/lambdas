package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/slot"
)

func CreateSlot(w http.ResponseWriter, req *http.Request) {
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

	var input edgarlib.CreateSlotInput

	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	slot := edgarlib.CreateSlot(input, doctorID.ID)
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
