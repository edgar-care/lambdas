package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	edgarlib "github.com/edgar-care/edgarlib/v2/medicament"
	"github.com/edgar-care/medicament/cmd/main/lib"
)

func GetMedicament(w http.ResponseWriter, req *http.Request) {

	t := chi.URLParam(req, "id")

	medicament := edgarlib.GetMedicamentById(t)
	if medicament.Err != nil {
		lib.WriteError(w, medicament.Code, medicament.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"medicament": medicament.Medicine,
	}, 201)
}

func GetMedicaments(w http.ResponseWriter, req *http.Request) {

	var medicaments edgarlib.GetMedicamentsResponse
	page := req.URL.Query().Get("page")
	size := req.URL.Query().Get("size")
	if page == "" && size == "" {
		medicaments = edgarlib.GetMedicaments(0, 0)
	} else {
		number_page, err1 := strconv.Atoi(page)
		number_size, err2 := strconv.Atoi(size)
		if err1 != nil || err2 != nil {
			medicaments = edgarlib.GetMedicaments(0, 0)
		} else {
			medicaments = edgarlib.GetMedicaments(number_page, number_size)
		}
	}
	if medicaments.Err != nil {
		lib.WriteError(w, medicaments.Code, medicaments.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"medicament": medicaments.Medicines,
	}, 201)
}
