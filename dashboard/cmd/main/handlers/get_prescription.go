package handlers

import (
	"github.com/edgar-care/dashboard/cmd/main/lib"
	"github.com/go-chi/chi/v5"

	authlib "github.com/edgar-care/edgarlib/v2/auth"
	edgarlib "github.com/edgar-care/edgarlib/v2/ordonnance"
	"net/http"
)

func GetPrescriptionByID(w http.ResponseWriter, req *http.Request) {
	doctorID := authlib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}
	t := chi.URLParam(req, "id")

	ordonnance := edgarlib.GetOrdonnancebyID(t)
	if ordonnance.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": ordonnance.Err.Error(),
		}, ordonnance.Code)
		return
	}
	lib.WriteResponse(w, map[string]interface{}{
		"prescription":     ordonnance.Ordonnance,
		"url_prescription": ordonnance.Url,
	}, 200)
}

func GetPrescription(w http.ResponseWriter, req *http.Request) {
	doctorID := authlib.AuthMiddlewareDoctor(w, req)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	ordonnance := edgarlib.GetOrdonnancesDoctor(doctorID)
	if ordonnance.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": ordonnance.Err.Error(),
		}, ordonnance.Code)
		return
	}
	lib.WriteResponse(w, map[string]interface{}{
		"prescription": ordonnance.Ordonnance,
	}, 200)
}
