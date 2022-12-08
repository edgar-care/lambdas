package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertUser(user *models.UserCreateInput) (*models.User, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("User").InsertOne(ctx, user)

	entity := models.User{
		ID:       result.InsertedID.(primitive.ObjectID),
		Email:    user.Email,
		Password: user.Password,
		Name:     user.Name,
		Age:      user.Age,
	}
	return &entity, err
}

func (db *DB) FindUsers() (*[]models.User, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.User

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("User").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) FindUserByID(id string) (*models.User, error) {
	ctx := context.Background()
	var result models.User
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("User").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateUser(user *models.UserUpdateInput) (*models.User, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(user.ID)
	filter := bson.M{"_id": objId}
	replacement, err := db.FindUserByID(user.ID)

	copier.CopyWithOption(replacement, user, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("User").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteUser(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("User").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
