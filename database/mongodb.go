package database

import (
	"car-collection-manager/utils"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()
var CarsCollection *mongo.Collection
var Client = mongo.Client{}

func ConnectMongoDb() {
	uri := utils.GetEnvOrDefault("DBAAS_MONGODB_ENDPOINT", "")
	dbName := utils.GetEnvOrDefault("DBAAS_MONGODB_APP", "app")

	// Create a new MongoDB client
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	Client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB successfully")

	CarsCollection = Client.Database(dbName).Collection("cars")
}
