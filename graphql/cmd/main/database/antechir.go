package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertAnteChir(antechir *models.AnteChirCreateInput) (*models.AnteChir, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteChir").InsertOne(ctx, antechir)

	entity := models.AnteChir{
		ID:              result.InsertedID.(primitive.ObjectID),
		Name:            antechir.Name,
		Localisation:    antechir.Localisation,
		InducedSymptoms: antechir.InducedSymptoms,
	}
	return &entity, err
}

func (db *DB) GetAnteChirs() (*[]models.AnteChir, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.AnteChir

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteChir").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetAnteChirByID(id string) (*models.AnteChir, error) {
	ctx := context.Background()
	var result models.AnteChir
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteChir").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateAnteChir(antechir *models.AnteChirUpdateInput) (*models.AnteChir, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(antechir.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetAnteChirByID(antechir.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, antechir, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteChir").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteAnteChir(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteChir").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
