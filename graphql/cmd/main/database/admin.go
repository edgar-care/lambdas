package database

import (
	"context"
	"os"

	"github.com/edgar-care/graphql/cmd/main/database/models"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) InsertAdmin(admin *models.AdminCreateInput) (*models.Admin, error) {
	ctx := context.Background()
	result, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Admin").InsertOne(ctx, admin)

	entity := models.Admin{
		ID:       result.InsertedID.(primitive.ObjectID),
		Email:    admin.Email,
		Password: admin.Password,
		Name:     admin.Name,
		LastName: admin.LastName,
	}
	return &entity, err
}

func (db *DB) GetAdmins() (*[]models.Admin, error) {
	ctx := context.Background()
	filter := bson.D{}
	var results []models.Admin

	cursor, err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Admin").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (db *DB) GetAdminByID(id string) (*models.Admin, error) {
	ctx := context.Background()
	var result models.Admin
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}

	err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Admin").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) GetAdminByEmail(email string) (*models.Admin, error) {
	ctx := context.Background()
	var result models.Admin

	filter := bson.M{"email": email}

	err := db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Admin").FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (db *DB) UpdateAdmin(admin *models.AdminUpdateInput) (*models.Admin, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(admin.ID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objId}
	replacement, err := db.GetAdminByID(admin.ID)
	if err != nil {
		return nil, err
	}

	copier.CopyWithOption(replacement, admin, copier.Option{IgnoreEmpty: true})

	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Admin").ReplaceOne(ctx, filter, replacement)

	return replacement, err
}

func (db *DB) DeleteAdmin(id string) (bool, error) {
	ctx := context.Background()
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, err
	}
	filter := bson.M{"_id": objId}
	_, err = db.client.Database(os.Getenv("DATABASE_NAME")).Collection("Admin").DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}
	return true, err

}
