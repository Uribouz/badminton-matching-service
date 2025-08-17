package match

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetMatch(ctx context.Context, eventId string, courtNo int, dateTime string) (*Match, error)
	SaveMatch(ctx context.Context, match *Match) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(database *mongo.Database) Repository {
	return &MongoRepository{
		collection: database.Collection("matches"),
	}
}

func (r *MongoRepository) GetMatch(ctx context.Context, eventId string, courtNo int, dateTime string) (*Match, error) {
	var match Match
	filter := bson.M{"event_id": eventId, "court_no": courtNo, "date_time": dateTime}

	err := r.collection.FindOne(ctx, filter).Decode(&match)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &match, nil
}

func (r *MongoRepository) SaveMatch(ctx context.Context, match *Match) error {
	filter := bson.M{"event_id": match.EventId, "court_no": match.CourtNo, "date_time": match.DateTime}
	update := bson.M{"$set": match}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)

	return err
}