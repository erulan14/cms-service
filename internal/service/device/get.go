package device

import (
	"cms-service/internal/model"
	"context"
	"log"
)

func (s *service) Get(ctx context.Context, uuid string) (*model.Device, error) {
	device, err := s.deviceRepository.Get(ctx, uuid)
	if err != nil {
		log.Printf("deviceRepository.Get error: %v", err)
		return nil, err
	}

	if device == nil {
		log.Printf("device not found")
		return nil, model.ErrNotFound
	}

	return device, err
}
