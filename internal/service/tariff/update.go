package tariff

import (
	"cms-service/internal/model"
	"context"
)

func (s *service) Update(ctx context.Context, uuid string, model *model.UpdateTariffDTO) error {
	err := s.tariffRepository.Update(ctx, uuid, model)
	if err != nil {
		return err
	}

	return err
}
