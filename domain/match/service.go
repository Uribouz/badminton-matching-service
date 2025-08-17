package match

import (
	"badminton-service/common/logger"
	"context"
)

type Service struct{ repo Repository }

func NewService(repo Repository) Service {
	return Service{repo}
}

func (s Service) GetMatch(eventId string, courtNo int, dateTime string) (Match, error) {
	logger.Log.Infof("GetMatch: eventId=%v, courtNo=%v, dateTime=%v", eventId, courtNo, dateTime)

	ctx := context.Background()
	match, err := s.repo.GetMatch(ctx, eventId, courtNo, dateTime)
	if err != nil {
		return Match{}, err
	}

	if match == nil {
		return Match{}, nil
	}

	return *match, nil
}

func (s Service) SaveMatch(match Match) error {
	logger.Log.Infof("SaveMatch: %v", match)

	ctx := context.Background()
	return s.repo.SaveMatch(ctx, &match)
}