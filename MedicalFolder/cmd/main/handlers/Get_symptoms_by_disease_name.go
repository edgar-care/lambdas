package handlers

import (
	"github.com/edgar-care/MedicalFolder/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/disease"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetSymptomsByDiseaseName(w http.ResponseWriter, req *http.Request) {

	t := chi.URLParam(req, "name")

	disease := edgarlib.GetSymptomsByDiseaseName(t)
	if disease.Err != nil {
		lib.WriteError(w, disease.Code, disease.Err.Error())
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"symptoms": disease.Symptoms,
	}, 200)
}
