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

func GetAllDocument(w http.ResponseWriter, req *http.Request) {
	patientID := lib.AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, 401)
		return
	}

	document, err := services.GetAll(patientID)

	if err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": "Invalid input: " + err.Error(),
		}, 400)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"document": document,
	}, 200)
}
