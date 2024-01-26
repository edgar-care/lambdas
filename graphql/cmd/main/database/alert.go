package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertAlert(alert *models.AlertCreateInput) (*models.Alert, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Alert").InsertOne(ctx, alert)

	entity := models.Alert{
		ID:       result.InsertedID.(primitive.ObjectID),
		Name:     alert.Name,
		Sex:      alert.Sex,
		Height:   alert.Height,
		Weight:   alert.Weight,
		Symptoms: alert.Symptoms,
		Comment:  alert.Comment,
	}
	return &entity, err
}

func (db *DB) GetAlerts() (*[]models.Alert, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Alert

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Alert").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetAlertByID(id string) (*models.Alert, error) {
	ctx := context.Background()
	var result models.Alert
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Alert").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateAlert(alert *models.AlertUpdateInput) (*models.Alert, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(alert.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetAlertByID(alert.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, alert, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Alert").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteAlert(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Alert").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
