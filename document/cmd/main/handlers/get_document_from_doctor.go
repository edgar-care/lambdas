package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/go-chi/chi/v5"

	edgarlib "github.com/edgar-care/edgarlib/v2/document"
	response "github.com/edgar-care/edgarlib/v2/http"
)

func DownloadFromDoctor(w http.ResponseWriter, r *http.Request) {
	doctorID := authlib.AuthMiddlewareDoctor(w, r)
	if doctorID.Code == 409 || doctorID.Code == 401 {
		response.WriteResponse(w, map[string]string{
			"message": doctorID.Err.Error(),
		}, doctorID.Code)
		return
	}
	if doctorID.ID == "" {
		response.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	IdDocumement := chi.URLParam(r, "id")

	downloadDocument := edgarlib.GetDocument(IdDocumement)
	if downloadDocument.Err != nil {
		response.WriteResponse(w, map[string]string{
			"message": downloadDocument.Err.Error(),
		}, downloadDocument.Code)
		return
	}

	// Return the document details in the response
	response := map[string]interface{}{
		"download": downloadDocument.Document,
		"message":  "Document get succesfuly",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
