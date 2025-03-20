package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserRole struct {
	ID     primitive.ObjectID `bson:"ID,omitempty"`
	UserID primitive.ObjectID `bson:"UserID,omitempty"`
	RoleID primitive.ObjectID `bson:"RoleID,omitempty"`
}
