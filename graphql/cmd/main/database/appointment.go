package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertRdv(rdv *models.RdvCreateInput) (*models.Rdv, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Rdv").InsertOne(ctx, rdv)
	entity := models.Rdv{
		ID:        result.InsertedID.(primitive.ObjectID),
		DoctorID:  rdv.DoctorID,
		IdPatient: rdv.IdPatient,
		StartDate: rdv.StartDate,
		EndDate:   rdv.EndDate,
	}
	return &entity, err
}

func (db *DB) GetPatientRdv(id_patient string) (*[]models.Rdv, error) {
	ctx := context.Background()
	var results []models.Rdv

	filter := bson.M{"id_patient": id_patient}

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Rdv").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetDoctorRdv(doctor_id string) (*[]models.Rdv, error) {
	ctx := context.Background()
	var results []models.Rdv

	filter := bson.M{"doctor_id": doctor_id, "id_patient": nil}

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Rdv").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetRdvByID(id string) (*models.Rdv, error) {
	ctx := context.Background()
	var result models.Rdv
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Rdv").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateRdv(rdv *models.RdvUpdateInput) (*models.Rdv, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(rdv.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetRdvByID(rdv.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, rdv, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Rdv").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteRdv(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}

	filter := bson.M{"_id": objId}
	update := bson.M{
		"$unset": bson.M{
			"id_patient": 1,
		},
	}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Rdv").UpdateOne(ctx, filter, update)
	if err != nil {
		return false, err
	}
	return true, nil
}
