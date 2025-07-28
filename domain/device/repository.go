package device

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetDevice(ctx context.Context, id string) (*Device, error)
	SaveDevice(ctx context.Context, device *Device) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(database *mongo.Database) Repository {
	return &MongoRepository{
		collection: database.Collection("devices"),
	}
}

func (r *MongoRepository) GetDevice(ctx context.Context, id string) (*Device, error) {
	var device Device
	filter := bson.M{"id": id}

	err := r.collection.FindOne(ctx, filter).Decode(&device)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &device, nil
}

func (r *MongoRepository) SaveDevice(ctx context.Context, device *Device) error {
	filter := bson.M{"id": device.Id}
	update := bson.M{"$set": device}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)

	return err
}
