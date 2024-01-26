package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertAnteDisease(antedisease *models.AnteDiseaseCreateInput) (*models.AnteDisease, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteDisease").InsertOne(ctx, antedisease)

	entity := models.AnteDisease{
		ID:         result.InsertedID.(primitive.ObjectID),
		Name:       antedisease.Name,
		Chronicity: antedisease.Chronicity,
		Chir:       antedisease.Chir,
		Treatment:  antedisease.Treatment,
		Symptoms:   antedisease.Symptoms,
	}
	return &entity, err
}

func (db *DB) GetAnteDiseases() (*[]models.AnteDisease, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.AnteDisease

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteDisease").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetAnteDiseaseByID(id string) (*models.AnteDisease, error) {
	ctx := context.Background()
	var result models.AnteDisease
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteDisease").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateAnteDisease(antedisease *models.AnteDiseaseUpdateInput) (*models.AnteDisease, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(antedisease.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetAnteDiseaseByID(antedisease.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, antedisease, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteDisease").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteAnteDisease(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteDisease").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
