package geozone

import (
	"cms-service/internal/model"
	"context"
	"log"
)

func (s *service) Get(ctx context.Context, uuid string) (*model.Geozone, error) {
	geozone, err := s.geozoneRepository.Get(ctx, uuid)
	if err != nil {
		log.Printf("GeozoneRepository.Get error: %v", err)
		return nil, err
	}

	if geozone == nil {
		log.Printf("Geozone not found")
		return nil, model.ErrNotFound
	}

	return geozone, err
}
