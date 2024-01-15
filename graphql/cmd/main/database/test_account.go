package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertTestAccount(testAccount *models.TestAccountCreateInput) (*models.TestAccount, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("TestAccount").InsertOne(ctx, testAccount)

	entity := models.TestAccount{
		ID:       result.InsertedID.(primitive.ObjectID),
		Email:    testAccount.Email,
		Password: testAccount.Password,
	}
	return &entity, err
}

func (db *DB) GetTestAccounts() (*[]models.TestAccount, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.TestAccount

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("TestAccount").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetTestAccountByID(id string) (*models.TestAccount, error) {
	ctx := context.Background()
	var result models.TestAccount
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("TestAccount").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) GetTestAccountByEmail(email string) (*models.TestAccount, error) {
	ctx := context.Background()
	var result models.TestAccount

	filter := bson.M{"email": email}

	err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("TestAccount").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateTestAccount(testAccount *models.TestAccountUpdateInput) (*models.TestAccount, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(testAccount.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetTestAccountByID(testAccount.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, testAccount, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("TestAccount").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteTestAccount(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("TestAccount").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
