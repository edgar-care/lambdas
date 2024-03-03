package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/edgar-care/document/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/document"
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

func GetAllDocument(w http.ResponseWriter, req *http.Request) {
	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	document := edgarlib.GetDocuments(patientID)

	if document.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": document.Err.Error(),
		}, document.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"document": document.Documents,
	}, 200)
}
