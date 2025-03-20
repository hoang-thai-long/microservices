package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `bson:"Title,omitempty"`
	Descriptions string             `bson:"Descriptions"`
}
