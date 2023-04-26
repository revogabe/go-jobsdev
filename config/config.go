package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db     *mongo.Database
	logger *Logger
)

func Init() error {
	var err error

	// Initialize MongoDB Atlas
	clientOptions := options.Client().ApplyURI(os.Getenv("CONNECTION_STRING"))
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return fmt.Errorf("failed to initialize MongoDB Atlas: %v", err)
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB Atlas: %v", err)
	}

	// Get the database
	db = client.Database("mydatabase")

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetMongoDB() *mongo.Database {
	return db
}
