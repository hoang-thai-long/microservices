package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"Name"`
	Phone    string             `bson:"Phone"`
	Email    string             `bson:"Email"`
	Born     primitive.DateTime `bson:"Born"`
	CreateAt primitive.DateTime `bson:"CreateAt"`
	UpdateAt primitive.DateTime `bson:"UpdateAt"`
}
