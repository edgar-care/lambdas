package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/slot"
)

func GetRdvDoctor(w http.ResponseWriter, req *http.Request) {

	t := chi.URLParam(req, "id")

	var rdv edgarlib.GetSlotsResponse
	page := req.URL.Query().Get("page")
	size := req.URL.Query().Get("size")
	if page == "" && size == "" {
		rdv = edgarlib.GetSlots(t, 0, 0)
	} else {
		number_page, err1 := strconv.Atoi(page)
		number_size, err2 := strconv.Atoi(size)
		if err1 != nil || err2 != nil {
			rdv = edgarlib.GetSlots(t, 0, 0)
		} else {
			rdv = edgarlib.GetSlots(t, number_page, number_size)
		}
	}

	if rdv.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": rdv.Err.Error(),
		}, rdv.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"rdv": rdv.Slots,
	}, 200)
}
