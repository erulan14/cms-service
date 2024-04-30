package geozone

import (
	"cms-service/internal/model"
	"context"
	"github.com/google/uuid"
	"log"
	"time"
)

func (s *service) Create(ctx context.Context, model *model.CreateGeozoneDTO) (string, error) {
	GeozoneUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Error creating geozone uuid %v", err)
		return "", err
	}

	model.CreatedAt = time.Now()
	model.UpdatedAt = time.Now()

	err = s.geozoneRepository.Create(ctx, GeozoneUUID.String(), model)
	if err != nil {
		log.Printf("Error creating geozone %v", err)
		return "", err
	}

	return GeozoneUUID.String(), err
}
