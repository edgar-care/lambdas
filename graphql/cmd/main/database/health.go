package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertHealth(health *models.HealthCreateInput) (*models.Health, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Health").InsertOne(ctx, health)
	entity := models.Health{
		ID:           result.InsertedID.(primitive.ObjectID),
		PatientsAllergies:     health.PatientsAllergies,
		PatientsIllness: health.PatientsIllness,
		PatientsPrimaryDoctor: health.PatientsPrimaryDoctor,
		PatientsTreatments: health.PatientsTreatments,
	}
	return &entity, err
}

func (db *DB) GetHealths() (*[]models.Health, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Health

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Health").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetHealthByID(id string) (*models.Health, error) {
	ctx := context.Background()
	var result models.Health
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Health").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateHealth(health *models.HealthUpdateInput) (*models.Health, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(health.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetHealthByID(health.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, health, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Health").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteHealth(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Health").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}