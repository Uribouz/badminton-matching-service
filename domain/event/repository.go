package event

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetEvent(ctx context.Context, eventId string) (*Event, error)
	SaveEvent(ctx context.Context, event *Event) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(database *mongo.Database) Repository {
	return &MongoRepository{
		collection: database.Collection("events"),
	}
}

func (r *MongoRepository) GetEvent(ctx context.Context, eventId string) (*Event, error) {
	var event Event
	filter := bson.M{"event_id": eventId}

	err := r.collection.FindOne(ctx, filter).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &event, nil
}

func (r *MongoRepository) SaveEvent(ctx context.Context, event *Event) error {
	filter := bson.M{"event_id": event.EventId}
	update := bson.M{"$set": event}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)

	return err
}