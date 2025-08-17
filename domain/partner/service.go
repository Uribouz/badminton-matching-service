package partner

import (
	"badminton-service/common/logger"
	"context"
)

type Service struct{ repo Repository }

func NewService(repo Repository) Service {
	return Service{repo}
}

func (s Service) GetPartner(eventId, playerName string) (Partner, error) {
	logger.Log.Infof("GetPartner: eventId=%v, playerName=%v", eventId, playerName)

	ctx := context.Background()
	partner, err := s.repo.GetPartner(ctx, eventId, playerName)
	if err != nil {
		return Partner{}, err
	}

	if partner == nil {
		return Partner{}, nil
	}

	return *partner, nil
}

func (s Service) SavePartner(partner Partner) error {
	logger.Log.Infof("SavePartner: %v", partner)

	ctx := context.Background()
	return s.repo.SavePartner(ctx, &partner)
}