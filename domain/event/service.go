package event

import (
	"badminton-service/common/logger"
	"context"
)

type Service struct{ repo Repository }

func NewService(repo Repository) Service {
	return Service{repo}
}

func (s Service) GetEvent(eventId string) (Event, error) {
	logger.Log.Infof("GetEvent: %v", eventId)

	ctx := context.Background()
	event, err := s.repo.GetEvent(ctx, eventId)
	if err != nil {
		return Event{}, err
	}

	if event == nil {
		return Event{}, nil
	}

	return *event, nil
}

func (s Service) SaveEvent(event Event) error {
	logger.Log.Infof("SaveEvent: %v", event)

	ctx := context.Background()
	return s.repo.SaveEvent(ctx, &event)
}