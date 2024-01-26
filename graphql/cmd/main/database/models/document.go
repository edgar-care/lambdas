package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID           primitive.ObjectID `bson:"_id"`
	OwnerID      string             `bson:"owner_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	DocumentType string             `bson:"document_type,omitempty"`
	Category     string             `bson:"category,omitempty"`
	IsFavorite   bool               `bson:"is_favorite,omitempty"`
	DownloadURL  string             `bson:"download_url,omitempty"`
}

type DocumentCreateInput struct {
	OwnerID      string `bson:"owner_id,omitempty"`
	Name         string `bson:"name,omitempty"`
	DocumentType string `bson:"document_type,omitempty"`
	Category     string `bson:"category,omitempty"`
	IsFavorite   bool   `bson:"is_favorite,omitempty"`
	DownloadURL  string `bson:"download_url,omitempty"`
}

type DocumentUpdateInput struct {
	ID           string  `bson:"_id"`
	OwnerID      *string `bson:"owner_id,omitempty"`
	Name         *string `bson:"name,omitempty"`
	DocumentType *string `bson:"document_type,omitempty"`
	Category     *string `bson:"category,omitempty"`
	IsFavorite   *bool   `bson:"is_favorite,omitempty"`
	DownloadURL  *string `bson:"download_url,omitempty"`
}
