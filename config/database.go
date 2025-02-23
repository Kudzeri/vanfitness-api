package config

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var UserCollection *mongo.Collection
var ProfileCollection *mongo.Collection

func ConnectDB() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}
	UserCollection = client.Database(os.Getenv("DATABASE_NAME")).Collection("users")
	ProfileCollection = client.Database(os.Getenv("DATABASE_NAME")).Collection("profiles")
}