package device

import (
	"cms-service/internal/model"
	"context"
)

func (s *service) Update(ctx context.Context, uuid string, model *model.UpdateDeviceDTO) error {
	err := s.deviceRepository.Update(ctx, uuid, model)
	if err != nil {
		return err
	}

	return err
}
