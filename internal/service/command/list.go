package command

import (
	"cms-service/internal/model"
	"context"
)

func (s *service) List(ctx context.Context) ([]model.Command, error) {
	devices, err := s.commandRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return devices, nil
}
