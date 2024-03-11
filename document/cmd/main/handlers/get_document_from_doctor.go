package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/document/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/document"
	response "github.com/edgar-care/edgarlib/http"
)

func DownloadFromDoctor(w http.ResponseWriter, r *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, r)
	if doctorID == "" {
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
