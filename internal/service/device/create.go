package device

import (
	"cms-service/internal/model"
	"context"
	"github.com/google/uuid"
	"log"
)

func (s *service) Create(ctx context.Context, model *model.CreateUserDTO) (string, error) {
	deviceUUID, err := uuid.NewUUID()
	if err != nil {
		log.Printf("Error creating device uuid %v", err)
		return "", err
	}

	err = s.deviceRepository.Create(ctx, deviceUUID.String(), model)
	if err != nil {
		log.Printf("Error creating device %v", err)
		return "", err
	}

	return deviceUUID.String(), err
}
