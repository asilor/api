package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var db *mongo.Database

func InitDB(uri, database string) {
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error initializing MongoDB connection: %v", err)
	}

	db = client.Database(database)

	log.Println("MongoDB connection successfully established")
}

func CloseDB() {
	if err := db.Client().Disconnect(context.Background()); err != nil {
		log.Fatalf("Error disconnecting MongoDB connection: %v", err)
	}

	log.Println("MongoDB connection successfully closed")
}

func GetCollection(collection string) *mongo.Collection {
	return db.Collection(collection)
}