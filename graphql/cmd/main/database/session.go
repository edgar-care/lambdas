package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertSession(session *models.SessionCreateInput) (*models.Session, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Session").InsertOne(ctx, session)
	entity := models.Session{
		ID:           result.InsertedID.(primitive.ObjectID),
		Symptoms:     session.Symptoms,
		Age:          session.Age,
		Height:       session.Height,
		Weight:       session.Weight,
		Sex:          session.Sex,
		LastQuestion: session.LastQuestion,
	}
	return &entity, err
}

func (db *DB) GetSessions() (*[]models.Session, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Session

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Session").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetSessionByID(id string) (*models.Session, error) {
	ctx := context.Background()
	var result models.Session
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Session").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateSession(session *models.SessionUpdateInput) (*models.Session, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(session.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetSessionByID(session.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, session, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Session").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteSession(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Session").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
