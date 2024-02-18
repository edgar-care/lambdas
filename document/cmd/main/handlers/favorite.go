package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/document/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/document"
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

	document := edgarlib.GetDocument(IdDocument)
	if document.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": document.Err.Error(),
		}, document.Code)
		return
	}

	var input edgarlib.CreateDocumentInput
	err := json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	// Invert the value of IsFavorite
	input.IsFavorite = !document.Document.IsFavorite

	favorite := edgarlib.Updatefavorite(IdDocument, input.IsFavorite)
	if favorite.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": favorite.Err.Error(),
		}, favorite.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"update documents": favorite,
		"message":          "Document status favorite edited",
	}, http.StatusCreated)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {

	// =================================== //
	ownerID := lib.AuthMiddleware(w, r)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	IdDocument := chi.URLParam(r, "id")

	var input edgarlib.CreateDocumentInput
	err := json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	// Invert the value of IsFavorite
	input.IsFavorite = false

	favorite := edgarlib.Updatefavorite(IdDocument, input.IsFavorite)
	if favorite.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": favorite.Err.Error(),
		}, favorite.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"update documents": favorite,
		"message":          "Document status favorite deleted",
	}, http.StatusCreated)
}
