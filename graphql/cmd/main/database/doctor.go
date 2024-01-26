package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertDoctor(doctor *models.DoctorCreateInput) (*models.Doctor, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Doctor").InsertOne(ctx, doctor)

	entity := models.Doctor{
		ID:       result.InsertedID.(primitive.ObjectID),
		Email:    doctor.Email,
		Password: doctor.Password,
	}
	return &entity, err
}

func (db *DB) GetDoctors() (*[]models.Doctor, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Doctor

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Doctor").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetDoctorByID(id string) (*models.Doctor, error) {
	ctx := context.Background()
	var result models.Doctor
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Doctor").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) GetDoctorByEmail(email string) (*models.Doctor, error) {
	ctx := context.Background()
	var result models.Doctor

	filter := bson.M{"email": email}

	err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Doctor").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateDoctor(doctor *models.DoctorUpdateInput) (*models.Doctor, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(doctor.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetDoctorByID(doctor.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, doctor, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Doctor").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteDoctor(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Doctor").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
