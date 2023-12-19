package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertTreatment(treatment *models.TreatmentCreateInput) (*models.Treatment, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Treatment").InsertOne(ctx, treatment)

	entity := models.Treatment{
		ID:          result.InsertedID.(primitive.ObjectID),
		Name:        treatment.Name,
		Disease:     treatment.Disease,
		Symptoms:    treatment.Symptoms,
		SideEffects: treatment.SideEffects,
	}
	return &entity, err
}

func (db *DB) GetTreatments() (*[]models.Treatment, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Treatment

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Treatment").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetTreatmentByID(id string) (*models.Treatment, error) {
	ctx := context.Background()
	var result models.Treatment
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Treatment").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateTreatment(treatment *models.TreatmentUpdateInput) (*models.Treatment, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(treatment.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetTreatmentByID(treatment.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, treatment, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Treatment").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteTreatment(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Treatment").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
