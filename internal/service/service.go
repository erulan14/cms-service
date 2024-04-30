package service

import (
	"cms-service/internal/model"
	"context"
)

type DeviceService interface {
	Create(ctx context.Context, model *model.CreateUserDTO) (string, error)
	Get(ctx context.Context, uuid string) (*model.Device, error)
	List(ctx context.Context) ([]model.Device, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, uuid string, model *model.UpdateDeviceDTO) error
}

type CommandService interface {
	Create(ctx context.Context, model *model.CreateCommandDTO) (string, error)
	Get(ctx context.Context, uuid string) (*model.Command, error)
	List(ctx context.Context) ([]model.Command, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, uuid string, model *model.UpdateCommandDTO) error
}
