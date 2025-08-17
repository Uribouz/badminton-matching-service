package player

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetPlayer(ctx context.Context, eventId, playerName string) (*Player, error)
	SavePlayer(ctx context.Context, player *Player) error
}

type MongoRepository struct {
	collection *mongo.Collection
}

func NewMongoRepository(database *mongo.Database) Repository {
	return &MongoRepository{
		collection: database.Collection("players"),
	}
}

func (r *MongoRepository) GetPlayer(ctx context.Context, eventId, playerName string) (*Player, error) {
	var player Player
	filter := bson.M{"event_id": eventId, "player_name": playerName}

	err := r.collection.FindOne(ctx, filter).Decode(&player)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &player, nil
}

func (r *MongoRepository) SavePlayer(ctx context.Context, player *Player) error {
	filter := bson.M{"event_id": player.EventId, "player_name": player.PlayerName}
	update := bson.M{"$set": player}

	opts := options.Update().SetUpsert(true)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)

	return err
}