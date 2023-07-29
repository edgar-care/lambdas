package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertSymptom(symptom *models.SymptomCreateInput) (*models.Symptom, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Symptom").InsertOne(ctx, symptom)

	entity := models.Symptom{
		ID:       result.InsertedID.(primitive.ObjectID),
		Code:     symptom.Code,
		Symptom:  symptom.Symptom,
		Advice:   symptom.Advice,
		Question: symptom.Question,
	}
	return &entity, err
}

func (db *DB) GetSymptoms() (*[]models.Symptom, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Symptom

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Symptom").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetSymptomByID(id string) (*models.Symptom, error) {
	ctx := context.Background()
	var result models.Symptom
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Symptom").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateSymptom(symptom *models.SymptomUpdateInput) (*models.Symptom, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(symptom.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetSymptomByID(symptom.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, symptom, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Symptom").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteSymptom(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Symptom").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
