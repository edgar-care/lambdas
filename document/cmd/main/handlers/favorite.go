package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/edgar-care/document/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/document"
	"github.com/go-chi/chi/v5"
)

func HandleFavorite(w http.ResponseWriter, r *http.Request) {

	ownerID := authlib.AuthMiddlewarePatient(w, r)
	if ownerID.Code == 409 || ownerID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": ownerID.Err.Error(),
		}, ownerID.Code)
		return
	}
	if ownerID.ID == "" {
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

	input.IsFavorite = true

	favorite := edgarlib.Updatefavorite(IdDocument, input.IsFavorite, ownerID.ID)
	if favorite.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": favorite.Err.Error(),
		}, favorite.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"update documents": favorite.Document,
		"message":          "Document status favorite edited",
	}, http.StatusCreated)
}

func RemoveFavorite(w http.ResponseWriter, r *http.Request) {

	ownerID := authlib.AuthMiddlewarePatient(w, r)
	if ownerID.Code == 409 || ownerID.Code == 401 {
		lib.WriteResponse(w, map[string]string{
			"message": ownerID.Err.Error(),
		}, ownerID.Code)
		return
	}
	if ownerID.ID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	IdDocument := chi.URLParam(r, "id")

	var input edgarlib.CreateDocumentInput
	err := json.NewDecoder(r.Body).Decode(&input)
	lib.CheckError(err)

	input.IsFavorite = false

	favorite := edgarlib.Updatefavorite(IdDocument, input.IsFavorite, ownerID.ID)
	if favorite.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": favorite.Err.Error(),
		}, favorite.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"update documents": favorite.Document,
		"message":          "Document status favorite deleted",
	}, http.StatusCreated)
}
