package tariff

import (
	"cms-service/internal/model"
	"context"
	"log"
)

func (s *service) Get(ctx context.Context, uuid string) (*model.Tariff, error) {
	tariff, err := s.tariffRepository.Get(ctx, uuid)
	if err != nil {
		log.Printf("tariffRepository.Get error: %v", err)
		return nil, err
	}

	if tariff == nil {
		log.Printf("tariff not found")
		return nil, model.ErrNotFound
	}

	return tariff, err
}
