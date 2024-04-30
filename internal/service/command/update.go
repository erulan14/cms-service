package command

import (
	"cms-service/internal/model"
	"context"
)

func (s *service) Update(ctx context.Context, uuid string, model *model.UpdateCommandDTO) error {
	err := s.commandRepository.Update(ctx, uuid, model)
	if err != nil {
		return err
	}

	return err
}
