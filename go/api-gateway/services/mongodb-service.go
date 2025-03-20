package services

import (
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var lock = &sync.Mutex{}

var singleInstance *mongo.Client

func ConnectDb() *mongo.Client {
	url := os.Getenv("MONGODB_URL")
	var clientOptions = options.Client().ApplyURI(url)
	client, _ := mongo.Connect(clientOptions)
	return client
}

func getInstance() *mongo.Client {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			singleInstance = ConnectDb()
		}
	}
	return singleInstance
}

type MongoService interface {
	Client() *mongo.Client
}

func Client() mongo.Client {
	return *getInstance()
}
