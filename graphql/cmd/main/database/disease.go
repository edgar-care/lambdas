package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertDisease(disease *models.DiseaseCreateInput) (*models.Disease, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Disease").InsertOne(ctx, disease)

	entity := models.Disease{
		ID:       result.InsertedID.(primitive.ObjectID),
		Code:     disease.Code,
		Name:     disease.Name,
		Symptoms: disease.Symptoms,
		Advice:   disease.Advice,
	}
	return &entity, err
}

func (db *DB) GetDiseases() (*[]models.Disease, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Disease

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Disease").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetDiseaseByID(id string) (*models.Disease, error) {
	ctx := context.Background()
	var result models.Disease
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Disease").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateDisease(disease *models.DiseaseUpdateInput) (*models.Disease, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(disease.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetDiseaseByID(disease.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, disease, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Disease").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteDisease(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Disease").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
