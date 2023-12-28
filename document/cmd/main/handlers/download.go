package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/document/cmd/main/lib"
	"github.com/edgar-care/document/cmd/main/services"
)

func HandleDownload(w http.ResponseWriter, r *http.Request) {
	ownerID := lib.AuthMiddleware(w, r)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	IdDocumement := chi.URLParam(r, "id")

	downloadDocument, err := services.GetDocument(IdDocumement)
	if err != nil {
		http.Error(w, "Failed to create document: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the document details in the response
	lib.WriteResponse(w, map[string]interface{}{
		"download": downloadDocument,
		"message":  "Document get succesfuly",
	}, http.StatusCreated)
}
