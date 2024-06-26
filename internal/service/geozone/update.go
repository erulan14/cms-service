package geozone

import (
	"cms-service/internal/model"
	"context"
	"time"
)

func (s *service) Update(ctx context.Context, uuid string, model *model.UpdateGeozoneDTO) error {
	model.UpdatedAt = time.Now()

	err := s.geozoneRepository.Update(ctx, uuid, model)
	if err != nil {
		return err
	}
	return err
}
