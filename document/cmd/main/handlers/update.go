package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/document/cmd/main/lib"
	"github.com/edgar-care/document/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) {

	ownerID := lib.AuthMiddleware(w, r)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	IdDocument := chi.URLParam(r, "id")

	var input services.DocumentInput
	err := json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	document, err := services.UpdateDocument(IdDocument, input)
	if err != nil {
		http.Error(w, "Failed to update document: "+err.Error(), http.StatusInternalServerError)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"update documents": document,
		"message":          "Document name change",
	}, http.StatusCreated)
}
