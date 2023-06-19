package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertInfo(info *models.InfoCreateInput) (*models.Info, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Info").InsertOne(ctx, info)
	entity := models.Info{
		ID:           result.InsertedID.(primitive.ObjectID),
		Name:     info.Name,
		Age:          info.Age,
		Height:       info.Height,
		Weight:       info.Weight,
		Sexe:          info.Sexe,
		Surname: info.Surname,
	}
	return &entity, err
}

func (db *DB) GetInfos() (*[]models.Info, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Info

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Info").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetInfoByID(id string) (*models.Info, error) {
	ctx := context.Background()
	var result models.Info
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Info").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateInfo(info *models.InfoUpdateInput) (*models.Info, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(info.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetInfoByID(info.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, info, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Info").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteInfo(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Info").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}