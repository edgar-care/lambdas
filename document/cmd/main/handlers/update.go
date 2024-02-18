package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/edgar-care/document/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/document"
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
		"update documents": document,
		"message":          "Document name change",
	}, http.StatusCreated)
}
