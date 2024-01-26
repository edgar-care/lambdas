package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Alert struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name,omitempty"`
	Sex      *string            `bson:"sex,omitempty"`
	Height   *int32             `bson:"height,omitempty"`
	Weight   *int32             `bson:"weight,omitempty"`
	Symptoms []string           `bson:"symptoms,omitempty"`
	Comment  string             `bson:"comment,omitempty"`
}

type AlertCreateInput struct {
	Name     string   `bson:"name,omitempty"`
	Sex      *string  `bson:"sex,omitempty"`
	Height   *int32   `bson:"height,omitempty"`
	Weight   *int32   `bson:"weight,omitempty"`
	Symptoms []string `bson:"symptoms,omitempty"`
	Comment  string   `bson:"comment,omitempty"`
}

type AlertUpdateInput struct {
	ID       string    `bson:"_id"`
	Name     *string   `bson:"name,omitempty"`
	Sex      *string   `bson:"sex,omitempty"`
	Height   *int32    `bson:"height,omitempty"`
	Weight   *int32    `bson:"weight,omitempty"`
	Symptoms *[]string `bson:"symptoms,omitempty"`
	Comment  *string   `bson:"comment,omitempty"`
}
