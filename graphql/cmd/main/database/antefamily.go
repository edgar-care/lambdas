package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertAnteFamily(antefamily *models.AnteFamilyCreateInput) (*models.AnteFamily, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteFamily").InsertOne(ctx, antefamily)

	entity := models.AnteFamily{
		ID:      result.InsertedID.(primitive.ObjectID),
		Name:    antefamily.Name,
		Disease: antefamily.Disease,
	}
	return &entity, err
}

func (db *DB) GetAnteFamilies() (*[]models.AnteFamily, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.AnteFamily

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteFamily").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetAnteFamilyByID(id string) (*models.AnteFamily, error) {
	ctx := context.Background()
	var result models.AnteFamily
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteFamily").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateAnteFamily(antefamily *models.AnteFamilyUpdateInput) (*models.AnteFamily, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(antefamily.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetAnteFamilyByID(antefamily.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, antefamily, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteFamily").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteAnteFamily(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("AnteFamily").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
