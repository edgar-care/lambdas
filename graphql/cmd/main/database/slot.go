package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertSlot(slot *models.SlotCreateInput) (*models.Slot, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Slot").InsertOne(ctx, slot)
	entity := models.Slot{
		ID:            result.InsertedID.(primitive.ObjectID),
		DoctorID:      slot.DoctorID,
		AppointmentID: slot.AppointmentID,
		StartDate:     slot.StartDate,
		EndDate:       slot.EndDate,
	}
	return &entity, err
}

func (db *DB) GetDoctorSlot(doctor_id string) (*[]models.Slot, error) {
	ctx := context.Background()
	var results []models.Slot

	filter := bson.M{"doctor_id": doctor_id}

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Slot").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetSlot(doctor_id string) (*[]models.Slot, error) {
	ctx := context.Background()
	var results []models.Slot

	filter := bson.M{"doctor_id": doctor_id, "appointment_id": nil}

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Slot").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetSlotByID(id string) (*models.Slot, error) {
	ctx := context.Background()
	var result models.Slot
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Slot").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateSlot(slot *models.SlotUpdateInput) (*models.Slot, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(slot.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetSlotByID(slot.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, slot, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Slot").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteSlot(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Slot").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
