package tariff

import (
	"cms-service/internal/model"
	"context"
)

func (s *service) List(ctx context.Context) ([]model.Tariff, error) {
	tariffs, err := s.tariffRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return tariffs, nil
}
