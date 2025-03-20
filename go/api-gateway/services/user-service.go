package services

import (
	"api-gateway/models"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type UserService interface {
	GetAll() []models.User
	GetInfo() models.User
	SignUp() models.User
}

type userService struct {
	users []models.User
}

func GetCollection() *mongo.Collection {
	return getInstance().Database(os.Getenv("DB_NAME")).Collection("user")
}

// GetAll implements UserService.
func (u *userService) GetAll(ctx *gin.Context) []models.User {
	var results []models.User
	query := GetCollection()
	cursor, err := query.Find(ctx, bson.D{}, options.Find().SetLimit(50))
	if err != nil {
		panic(err)
	}
	cursor.All(ctx, &results)
	return results
}

// GetInfo implements UserService.
func (u *userService) GetInfo(ctx *gin.Context, id primitive.ObjectID) models.User {
	var result models.User
	query := GetCollection()
	cursor, err := query.Find(ctx, bson.D{{Key: "_id", Value: id}}, options.Find().SetLimit(1))
	if err != nil {
		panic(err)
	}
	cursor.All(ctx, &result)
	return result
}

// SignUp implements UserService.
func (u *userService) SignUp() models.User {
	panic("unimplemented")
}

func New() UserService {
	return &userService{}
}
