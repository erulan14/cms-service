package repository

import (
	"cms-service/internal/model"
	"context"
)

type DeviceRepository interface {
	Create(ctx context.Context, deviceUUID string, model *model.CreateUserDTO) error
	Get(ctx context.Context, deviceUUID string) (*model.Device, error)
	List(ctx context.Context) ([]model.Device, error)
	Update(ctx context.Context, deviceUUID string, model *model.UpdateDeviceDTO) error
	Delete(ctx context.Context, deviceUUID string) error
}
