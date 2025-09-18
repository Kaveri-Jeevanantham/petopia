package db

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// ConnectMongoDB establishes connection to MongoDB
func ConnectMongoDB() (*mongo.Client, error) {
	// Build MongoDB connection string from environment variables
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		// Build URI from individual components for better security
		mongoHost := os.Getenv("MONGO_HOST")
		if mongoHost == "" {
			// In Docker environment, use the service name as the hostname
			if os.Getenv("DOCKER_ENV") == "true" {
				mongoHost = "mongodb"
			} else {
				mongoHost = "localhost"
			}
		}

		mongoPort := os.Getenv("MONGO_PORT")
		if mongoPort == "" {
			mongoPort = "27017"
		}

		mongoUsername := os.Getenv("MONGO_USERNAME")
		if mongoUsername == "" {
			mongoUsername = "admin"
		}

		mongoPassword := os.Getenv("MONGO_PASSWORD")
		if mongoPassword == "" {
			mongoPassword = "P@ssw0rd"
		}

		// Construct URI with proper URL encoding for security
		mongoURI = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			url.QueryEscape(mongoUsername), url.QueryEscape(mongoPassword), mongoHost, mongoPort)
	}

	// Set client options with security and performance configurations
	clientOptions := options.Client().
		ApplyURI(mongoURI).
		SetMaxPoolSize(100).                        // Connection pool size
		SetMinPoolSize(5).                          // Minimum connections
		SetMaxConnIdleTime(30 * time.Second).       // Idle connection timeout
		SetServerSelectionTimeout(5 * time.Second). // Server selection timeout
		SetSocketTimeout(30 * time.Second)          // Socket timeout

	// Set connection timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Test the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	MongoClient = client
	log.Println("Successfully connected to MongoDB!")
	return client, nil
}

// GetMongoDatabase returns a database instance
func GetMongoDatabase() *mongo.Database {
	if MongoClient == nil {
		log.Fatal("MongoDB client not initialized")
	}

	dbName := os.Getenv("MONGO_DB_NAME")
	if dbName == "" {
		dbName = "petopia"
	}

	return MongoClient.Database(dbName)
}

// GetMongoCollection returns a collection instance
func GetMongoCollection(collectionName string) *mongo.Collection {
	db := GetMongoDatabase()
	return db.Collection(collectionName)
}

// DisconnectMongoDB closes the MongoDB connection
func DisconnectMongoDB() error {
	if MongoClient == nil {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return MongoClient.Disconnect(ctx)
}

// VerifyMongoConnection checks if the MongoDB connection is valid
func VerifyMongoConnection() bool {
	if MongoClient == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := MongoClient.Ping(ctx, nil)
	return err == nil
}

// ConnectMongoDBWithRetry attempts to connect with retries
func ConnectMongoDBWithRetry(maxRetries int) (*mongo.Client, error) {
	var err error
	var client *mongo.Client

	for i := 0; i < maxRetries; i++ {
		client, err = ConnectMongoDB()
		if err == nil {
			return client, nil
		}

		log.Printf("MongoDB connection attempt %d failed: %v. Retrying in 2 seconds...", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to MongoDB after %d attempts: %v", maxRetries, err)
}
