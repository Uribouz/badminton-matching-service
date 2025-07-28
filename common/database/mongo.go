package database

import (
	"badminton-service/config"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client   *mongo.Client
	Database *mongo.Database
}

func NewMongoDB(config config.MongoConfig) (*Database, error) {
	// Set client options with production-ready settings
	clientOptions := options.Client().
		ApplyURI(config.URI).
		SetMaxPoolSize(100).
		SetMinPoolSize(5).
		SetMaxConnIdleTime(30 * time.Second).
		SetServerSelectionTimeout(5 * time.Second).
		SetSocketTimeout(0).
		SetConnectTimeout(10 * time.Second).
		SetHeartbeatInterval(10 * time.Second)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping to verify connection
	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	log.Printf("Successfully connected to MongoDB at %s", config.URI)
	return &Database{client: client, Database: client.Database(config.Database)}, nil
}

func (db *Database) Disconnect(ctx context.Context) error {
	if db.client == nil {
		return nil
	}

	if err := db.client.Disconnect(ctx); err != nil {
		log.Printf("Failed to disconnect from MongoDB: %v", err)
		return err
	}

	log.Println("Successfully disconnected from MongoDB")
	return nil
}

func (db *Database) IsConnected(ctx context.Context) bool {
	if db.client == nil {
		return false
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	return db.client.Ping(ctx, nil) == nil
}

func (db *Database) Reconnect(config config.MongoConfig) error {
	if db.IsConnected(context.Background()) {
		return nil
	}

	newDB, err := NewMongoDB(config)
	if err != nil {
		return err
	}

	oldClient := db.client
	db.client = newDB.client
	db.Database = newDB.Database

	if oldClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		oldClient.Disconnect(ctx)
	}

	return nil
}
