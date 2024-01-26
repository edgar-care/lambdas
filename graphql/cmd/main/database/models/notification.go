package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notification struct {
	ID           primitive.ObjectID `bson:"_id"`
	Token     	string           `bson:"Token,omitempty"`
	Title string             `bson:"title,omitempty"`
	Message string `bson:"message,omitempty"`

}

type NotificationCreateInput struct {
	Token     	string           `bson:"Token,omitempty"`
	Title string             `bson:"title,omitempty"`
	Message string `bson:"message,omitempty"`
}

type NotificationUpdateInput struct {
	ID           string    `bson:"_id"`
	Token     	*string           `bson:"Token,omitempty"`
	Title *string             `bson:"title,omitempty"`
	Message *string `bson:"message,omitempty"`
}