package handlers

import (
	"encoding/json"
	authlib "github.com/edgar-care/edgarlib/v2/auth"
	"net/http"

	"github.com/edgar-care/document/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/v2/document"
	"github.com/go-chi/chi/v5"
)

func HandleUpdate(w http.ResponseWriter, r *http.Request) {

	ownerID := authlib.AuthMiddlewarePatient(w, r)
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

	document := edgarlib.UpdateDocument(input, IdDocument)
	if document.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": document.Err.Error(),
		}, document.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"update documents": document.Document,
		"message":          "Document name change",
	}, http.StatusCreated)
}
