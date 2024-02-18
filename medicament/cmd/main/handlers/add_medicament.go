package handlers

import (
	"encoding/json"
	"net/http"

	edgarlib "github.com/edgar-care/edgarlib/medicament"
	"github.com/edgar-care/medicament/cmd/main/lib"
)

func Addmedicament(w http.ResponseWriter, req *http.Request) {

	var input edgarlib.CreateMedicamentInput

	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		lib.WriteError(w, http.StatusBadRequest, "Invalid JSON input")
		return
	}

	medicamentResponse := edgarlib.CreateMedicament(input)

	if medicamentResponse.Err != nil {
		lib.WriteError(w, medicamentResponse.Code, medicamentResponse.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"medicament": medicamentResponse.Medicament,
	}, 201)
}
