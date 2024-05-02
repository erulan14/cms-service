package tariff

import (
	"cms-service/internal/model"
	"context"
	"github.com/google/uuid"
	"log"
)

func (s *service) Create(ctx context.Context, model *model.CreateTariffDTO) (string, error) {
	tariffUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Error creating tariff uuid %v", err)
		return "", err
	}

	err = s.tariffRepository.Create(ctx, tariffUUID.String(), model)
	if err != nil {
		log.Printf("Error creating tariff %v", err)
		return "", err
	}

	return tariffUUID.String(), err
}
