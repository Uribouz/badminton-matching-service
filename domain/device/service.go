package device

import (
	"badminton-service/common/logger"
	"context"
)

type Service struct{ repo Repository }

func NewService(repo Repository) Service {
	return Service{repo}
}

func (s Service) GetDevice(id string) (Device, error) {
	logger.Log.Infof("GetDevice: %v", id)

	ctx := context.Background()
	device, err := s.repo.GetDevice(ctx, id)
	if err != nil {
		return Device{}, err
	}

	if device == nil {
		return Device{}, nil
	}

	return *device, nil
}

func (s Service) SaveDevice(device Device) error {
	logger.Log.Infof("SaveDevice: %v", device)

	ctx := context.Background()
	return s.repo.SaveDevice(ctx, &device)
}
