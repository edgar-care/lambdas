package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	edgarlib "github.com/edgar-care/edgarlib/medicament"
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
		"medicament": medicament.Medicament,
	}, 201)
}

func GetMedicaments(w http.ResponseWriter, req *http.Request) {

	medicaments := edgarlib.GetMedicaments()
	if medicaments.Err != nil {
		lib.WriteError(w, medicaments.Code, medicaments.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"medicament": medicaments.Medicaments,
	}, 201)
}
