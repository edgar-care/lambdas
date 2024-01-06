package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertDemoAccount(demoAccount *models.DemoAccountCreateInput) (*models.DemoAccount, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("DemoAccount").InsertOne(ctx, demoAccount)

	entity := models.DemoAccount{
		ID:       result.InsertedID.(primitive.ObjectID),
		Email:    demoAccount.Email,
		Password: demoAccount.Password,
	}
	return &entity, err
}

func (db *DB) GetDemoAccounts() (*[]models.DemoAccount, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.DemoAccount

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("DemoAccount").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetDemoAccountByID(id string) (*models.DemoAccount, error) {
	ctx := context.Background()
	var result models.DemoAccount
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("DemoAccount").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) GetDemoAccountByEmail(email string) (*models.DemoAccount, error) {
	ctx := context.Background()
	var result models.DemoAccount

	filter := bson.M{"email": email}

	err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("DemoAccount").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateDemoAccount(demoAccount *models.DemoAccountUpdateInput) (*models.DemoAccount, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(demoAccount.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetDemoAccountByID(demoAccount.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, demoAccount, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("DemoAccount").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteDemoAccount(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("DemoAccount").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
