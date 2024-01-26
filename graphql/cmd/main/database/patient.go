package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertPatient(patient *models.PatientCreateInput) (*models.Patient, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Patient").InsertOne(ctx, patient)

	entity := models.Patient{
		ID:       result.InsertedID.(primitive.ObjectID),
		Email:    patient.Email,
		Password: patient.Password,
	}
	return &entity, err
}

func (db *DB) GetPatients() (*[]models.Patient, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Patient

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Patient").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetPatientByID(id string) (*models.Patient, error) {
	ctx := context.Background()
	var result models.Patient
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Patient").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) GetPatientByEmail(email string) (*models.Patient, error) {
	ctx := context.Background()
	var result models.Patient

	filter := bson.M{"email": email}

	err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Patient").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdatePatient(patient *models.PatientUpdateInput) (*models.Patient, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(patient.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetPatientByID(patient.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, patient, copier.Option{IgnoreEmpty: true})
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Patient").ReplaceOne(ctx, filter, replacement)
	return replacement, err
}

func (db *DB) DeletePatient(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Patient").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
