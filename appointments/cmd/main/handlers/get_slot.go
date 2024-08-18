package handlers

import (
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"
	"strconv"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/slot"
	"github.com/go-chi/chi/v5"
)

func GetSlotId(w http.ResponseWriter, req *http.Request) {
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

	id := chi.URLParam(req, "id")

	slot := edgarlib.GetSlotById(id, doctorID.ID)

	if slot.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": slot.Err.Error(),
		}, slot.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"slot": slot.Slot,
	}, 200)
}

func GetSlots(w http.ResponseWriter, req *http.Request) {

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
	var slot edgarlib.GetSlotsResponse
	page := req.URL.Query().Get("page")
	size := req.URL.Query().Get("size")
	if page == "" && size == "" {
		slot = edgarlib.GetSlots(doctorID.ID, 0, 0)
	} else {
		number_page, err1 := strconv.Atoi(page)
		number_size, err2 := strconv.Atoi(size)
		if err1 != nil || err2 != nil {
			slot = edgarlib.GetSlots(doctorID.ID, 0, 0)
		} else {
			slot = edgarlib.GetSlots(doctorID.ID, number_page, number_size)
		}
	}

	if slot.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": slot.Err.Error(),
		}, slot.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"slot": slot.Slots,
	}, 200)
}
