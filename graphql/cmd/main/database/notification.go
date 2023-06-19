package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertNotification(notification *models.NotificationCreateInput) (*models.Notification, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Notification").InsertOne(ctx, notification)
	entity := models.Notification{
		ID:           result.InsertedID.(primitive.ObjectID),
		Token:     notification.Token,
		Title: notification.Title,
		Message: notification.Message,
	}
	return &entity, err
}

func (db *DB) GetNotifications() (*[]models.Notification, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Notification

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Notification").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetNotificationByID(id string) (*models.Notification, error) {
	ctx := context.Background()
	var result models.Notification
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Notification").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateNotification(notification *models.NotificationUpdateInput) (*models.Notification, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(notification.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetNotificationByID(notification.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, notification, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Notification").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteNotification(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Notification").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}