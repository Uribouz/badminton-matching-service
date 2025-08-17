package partner

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetPartner(ctx context.Context, eventId, playerName string) (*Partner, error)
	SavePartner(ctx context.Context, partner *Partner) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(database *mongo.Database) Repository {
	return &MongoRepository{
		collection: database.Collection("partners"),
	}
}

func (r *MongoRepository) GetPartner(ctx context.Context, eventId, playerName string) (*Partner, error) {
	var partner Partner
	filter := bson.M{"event_id": eventId, "player_name": playerName}

	err := r.collection.FindOne(ctx, filter).Decode(&partner)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &partner, nil
}

func (r *MongoRepository) SavePartner(ctx context.Context, partner *Partner) error {
	filter := bson.M{"event_id": partner.EventId, "player_name": partner.PlayerName}
	update := bson.M{"$set": partner}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)

	return err
}