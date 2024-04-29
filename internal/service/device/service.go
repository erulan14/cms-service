package device

import (
	"cms-service/internal/repository"
	def "cms-service/internal/service"
)

var _ def.DeviceService = (*service)(nil)

type service struct {
	deviceRepository repository.DeviceRepository
}

func NewService(deviceRepository repository.DeviceRepository) *service {
	return &service{deviceRepository: deviceRepository}
}
