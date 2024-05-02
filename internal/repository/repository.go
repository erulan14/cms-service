package repository

import (
	"cms-service/internal/model"
	"context"
)

type DeviceRepository interface {
	Create(ctx context.Context, deviceUUID string, model *model.CreateDeviceDTO) error
	Get(ctx context.Context, deviceUUID string) (*model.Device, error)
	List(ctx context.Context) ([]model.Device, error)
	Update(ctx context.Context, deviceUUID string, model *model.UpdateDeviceDTO) error
	Delete(ctx context.Context, deviceUUID string) error
}

type CommandRepository interface {
	Create(ctx context.Context, deviceUUID string, model *model.CreateCommandDTO) error
	Get(ctx context.Context, deviceUUID string) (*model.Command, error)
	List(ctx context.Context) ([]model.Command, error)
	Update(ctx context.Context, deviceUUID string, model *model.UpdateCommandDTO) error
	Delete(ctx context.Context, deviceUUID string) error
}

type GeozoneRepository interface {
	Create(ctx context.Context, deviceUUID string, model *model.CreateGeozoneDTO) error
	Get(ctx context.Context, deviceUUID string) (*model.Geozone, error)
	List(ctx context.Context) ([]model.Geozone, error)
	Update(ctx context.Context, deviceUUID string, model *model.UpdateGeozoneDTO) error
	Delete(ctx context.Context, deviceUUID string) error
}

type TariffRepository interface {
	Create(ctx context.Context, deviceUUID string, model *model.CreateTariffDTO) error
	Get(ctx context.Context, deviceUUID string) (*model.Tariff, error)
	List(ctx context.Context) ([]model.Tariff, error)
	Update(ctx context.Context, deviceUUID string, model *model.UpdateTariffDTO) error
	Delete(ctx context.Context, deviceUUID string) error
}
