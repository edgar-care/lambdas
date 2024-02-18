package handlers

import (
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/edgar-care/document/cmd/main/lib"
	edgarlib "github.com/edgar-care/edgarlib/document"
	"github.com/go-chi/chi/v5"
)

const (
	maxFileSize = 10 << 20 // 10 MB max size
	maxBodySize = 1 << 20
)

var allowedExtensions = map[string]bool{
	".pdf":  true,
	".doc":  true,
	".png":  true,
	".odt":  true,
	".jpeg": true,
	".docx": true,
	".odtx": true,
}

func isValidFileExtension(filename string) bool {
	ext := filepath.Ext(filename)
	return allowedExtensions[ext]
}

func HandleUpload(w http.ResponseWriter, r *http.Request) {
	ownerID := lib.AuthMiddleware(w, r)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	// Parse the form data
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, header, err := r.FormFile("document")
	if err != nil {
		http.Error(w, "Failed to get document from request: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check if the file extension is valid
	if !isValidFileExtension(header.Filename) {
		http.Error(w, "Invalid file type. Only PDF, DOC, PNG, and ODT are allowed.", http.StatusBadRequest)
		return
	}

	// Extract values from form data
	documentType := r.FormValue("documentType")
	category := r.FormValue("category")
	isFavorite, err := strconv.ParseBool(r.FormValue("isFavorite"))
	if err != nil {
		http.Error(w, "Invalid is_favorite value", http.StatusBadRequest)
		return
	}

	// Use the UploadToS3 function to upload the file to S3
	downloadURL, err := lib.UploadToS3(file, header.Filename)
	if err != nil {
		http.Error(w, "Failed to upload document to S3: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate the S3 download URL
	document := edgarlib.UploadDocumentInput{
		OwnerID:      ownerID,
		DocumentType: documentType,
		Category:     category,
		IsFavorite:   isFavorite,
		Name:         header.Filename,
		DownloadURL:  downloadURL,
	}

	// Call CreateDocument to store the document in the external system
	createdDocument := edgarlib.CreateDocument(document, ownerID)
	if createdDocument.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": createdDocument.Err.Error(),
		}, createdDocument.Code)
		return
	}

	lib.WriteResponse(w, map[string]interface{}{
		"upload":  createdDocument.Document,
		"message": "Document created successfully",
	}, http.StatusCreated)
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	ownerID := lib.AuthMiddleware(w, r)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	IdDocumement := chi.URLParam(r, "id")

	delete := edgarlib.DeleteDocument(IdDocumement, ownerID)
	if delete.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": delete.Err.Error(),
		}, delete.Code)
		return
	}

	// Return the document details in the response
	lib.WriteResponse(w, map[string]interface{}{
		"delete":  delete,
		"message": "Document has been delete",
	}, http.StatusCreated)
}

func UploadFromDoctor(w http.ResponseWriter, r *http.Request) {
	ownerID := lib.AuthMiddlewareDoctor(w, r)
	if ownerID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return
	}

	// Parse the form data
	if err := r.ParseMultipartForm(maxFileSize); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve the file from the form data
	file, header, err := r.FormFile("document")
	if err != nil {
		http.Error(w, "Failed to get document from request: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Check if the file extension is valid
	if !isValidFileExtension(header.Filename) {
		http.Error(w, "Invalid file type. Only PDF, DOC, PNG, and ODT are allowed.", http.StatusBadRequest)
		return
	}

	// Extract values from form data
	patientID := r.FormValue("patient_id")
	documentType := r.FormValue("documentType")
	category := r.FormValue("category")
	isFavorite, err := strconv.ParseBool(r.FormValue("isFavorite"))
	if err != nil {
		http.Error(w, "Invalid is_favorite value", http.StatusBadRequest)
		return
	}

	// Use the UploadToS3 function to upload the file to S3
	downloadURL, err := lib.UploadToS3(file, header.Filename)
	if err != nil {
		http.Error(w, "Failed to upload document to S3: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate the S3 download URL
	document := edgarlib.UploadDocumentInput{
		OwnerID:      patientID,
		DocumentType: documentType,
		Category:     category,
		IsFavorite:   isFavorite,
		Name:         header.Filename,
		DownloadURL:  downloadURL,
	}

	// Call CreateDocument to store the document in the external system
	createdDocument := edgarlib.CreateDocument(document, patientID)
	if createdDocument.Err != nil {
		lib.WriteResponse(w, map[string]string{
			"message": createdDocument.Err.Error(),
		}, createdDocument.Code)
		return
	}

	// Return the document details in the response
	lib.WriteResponse(w, map[string]interface{}{
		"upload":  createdDocument.Document,
		"message": "Document created successfully",
	}, http.StatusCreated)
}
