package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/document/cmd/main/lib"
	"github.com/edgar-care/document/cmd/main/services"
	"github.com/go-chi/chi/v5"
)

func HandleFavorite(w http.ResponseWriter, r *http.Request) {

	// =================================== //
	ownerID := lib.AuthMiddleware(w, r)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	IdDocument := chi.URLParam(r, "id")

	document, err := services.GetDocument(IdDocument)
	if err != nil {
		http.Error(w, "Failed to create document: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var input services.DocumentInput
	err = json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	// Invert the value of IsFavorite
	input.IsFavorite = !document.IsFavorite

	_, err = services.UpdateFavoriteById(IdDocument, input.IsFavorite)
	if err != nil {
		http.Error(w, "Failed to update document: "+err.Error(), http.StatusInternalServerError)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"update documents": input.IsFavorite,
		"message":          "Document status favorite edited",
	}, http.StatusCreated)
}
