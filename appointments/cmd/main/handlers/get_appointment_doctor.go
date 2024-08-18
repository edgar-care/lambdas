package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/appointments/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/slot"
)

func GetRdvDoctor(w http.ResponseWriter, req *http.Request) {

	t := chi.URLParam(req, "id")

	rdv := edgarlib.GetSlots(t)

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
