package service

import (
	"cms-service/internal/model"
	"context"
)

type DeviceService interface {
	Create(ctx context.Context, model *model.CreateDeviceDTO) (string, error)
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

type GeozoneService interface {
	Create(ctx context.Context, model *model.CreateGeozoneDTO) (string, error)
	Get(ctx context.Context, uuid string) (*model.Geozone, error)
	List(ctx context.Context) ([]model.Geozone, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, uuid string, model *model.UpdateGeozoneDTO) error
}

type TariffService interface {
	Create(ctx context.Context, model *model.CreateTariffDTO) (string, error)
	Get(ctx context.Context, uuid string) (*model.Tariff, error)
	List(ctx context.Context) ([]model.Tariff, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, uuid string, model *model.UpdateTariffDTO) error
}
