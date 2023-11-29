package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/slot/cmd/main/lib"
	"github.com/edgar-care/slot/cmd/main/services"
)

func CreateSlot(w http.ResponseWriter, req *http.Request) {

	doctorID := lib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	var input services.SlotInput

	err := json.NewDecoder(req.Body).Decode(&input)

	lib.CheckError(err)
	slot, err := services.CreateSlots(input, doctorID)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Unable  (check if you share all information) " + err.Error(),
		}, 400)
		return
	}

	var updatedDoctor services.DoctorInput
	doctor, err := services.GetDoctorById(doctorID)
	lib.CheckError(err)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Id not correspond to a doctor",
		}, 400)
		return
	}

	updatedDoctor = services.DoctorInput{
		Id:      doctorID,
		SlotIDs: append(doctor.SlotIDs, slot.Id),
	}
	updatDoctor, err := services.UpdateDoctor(updatedDoctor)
	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Update Failed " + err.Error(),
		}, 500)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"slot":   slot,
		"doctor": updatDoctor,
	}, 201)
}
