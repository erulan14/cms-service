package geozone

import (
	"cms-service/internal/model"
	"context"
)

func (s *service) List(ctx context.Context) ([]model.Geozone, error) {
	geozones, err := s.geozoneRepository.List(ctx)
	if err != nil {
		return nil, err
	}
	return geozones, nil
}
