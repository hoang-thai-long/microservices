package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type AccountUser struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	AccountID primitive.ObjectID `bson:"AccountID,omitempty"`
	UserID    primitive.ObjectID `bson:"UserID,omitempty"`
}
