package player

import (
	"badminton-service/common/logger"
	"context"
)

type Service struct{ repo Repository }

func NewService(repo Repository) Service {
	return Service{repo}
}

func (s Service) GetPlayer(eventId, playerName string) (Player, error) {
	logger.Log.Infof("GetPlayer: eventId=%v, playerName=%v", eventId, playerName)

	ctx := context.Background()
	player, err := s.repo.GetPlayer(ctx, eventId, playerName)
	if err != nil {
		return Player{}, err
	}

	if player == nil {
		return Player{}, nil
	}

	return *player, nil
}

func (s Service) SavePlayer(player Player) error {
	logger.Log.Infof("SavePlayer: %v", player)

	ctx := context.Background()
	return s.repo.SavePlayer(ctx, &player)
}