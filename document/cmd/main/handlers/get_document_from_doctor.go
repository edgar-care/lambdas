package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/document/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/document"
)

func DownloadFromDoctor(w http.ResponseWriter, r *http.Request) {
	doctorID := lib.AuthMiddlewareDoctor(w, r)
	if doctorID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	IdDocumement := chi.URLParam(r, "id")

	downloadDocument := edgarlib.GetDocument(IdDocumement)
	if downloadDocument.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": downloadDocument.Err.Error(),
		}, downloadDocument.Code)
		return
	}

	// Return the document details in the response
	lib.WriteResponse(w, map[string]interface{}{
		"download": downloadDocument.Document,
		"message":  "Document get succesfuly",
	}, http.StatusCreated)
}
