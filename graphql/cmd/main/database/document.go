package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertDocument(document *models.DocumentCreateInput) (*models.Document, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Document").InsertOne(ctx, document)
	entity := models.Document{
		ID:           result.InsertedID.(primitive.ObjectID),
		OwnerID:      document.OwnerID,
		Name:         document.Name,
		DocumentType: document.DocumentType,
		Category:     document.Category,
		IsFavorite:   document.IsFavorite,
		DownloadURL:  document.DownloadURL,
	}
	return &entity, err
}

func (db *DB) GetDocuments() (*[]models.Document, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Document

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Document").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetDocumentByID(id string) (*models.Document, error) {
	ctx := context.Background()
	var result models.Document
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Document").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateDocument(document *models.DocumentUpdateInput) (*models.Document, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(document.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetDocumentByID(document.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, document, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Document").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteDocument(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Document").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
