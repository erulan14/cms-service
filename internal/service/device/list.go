package device

import (
	"cms-service/internal/model"
	"context"
)

func (s *service) List(ctx context.Context) ([]model.Device, error) {
	devices, err := s.deviceRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return devices, nil
}
