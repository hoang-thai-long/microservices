package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `bson:"UserName,omitempty"`
	Password string             `bson:"Password,omitempty"`
	TemPass  string             `bson:"TemPass"`
	CreateAt primitive.DateTime `bson:"CreateAt,omitempty"`
	UpdateAt primitive.DateTime `bson:"UpdateAt,omitempty"`
}
