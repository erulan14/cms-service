package app

import (
	"cms-service/internal/api/device"
	"cms-service/internal/repository"
	deviceRepository "cms-service/internal/repository/device"
	"cms-service/internal/service"
	deviceService "cms-service/internal/service/device"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type deviceServiceProvider struct {
	deviceRepository repository.DeviceRepository
	deviceService    service.DeviceService
	deviceImpl       *device.Implementation
	db               *pgxpool.Pool
}

func newDeviceServiceProvider(db *pgxpool.Pool) *deviceServiceProvider {
	var service deviceServiceProvider

	service.db = db

	return &service
}

func (s *deviceServiceProvider) DeviceRepository() repository.DeviceRepository {
	if s.deviceRepository == nil {
		s.deviceRepository = deviceRepository.NewRepository(s.db)
	}
	return s.deviceRepository
}

func (s *deviceServiceProvider) DeviceService() service.DeviceService {
	if s.deviceService == nil {
		s.deviceService = deviceService.NewService(
			s.DeviceRepository(),
		)
	}

	return s.deviceService
}

func (s *deviceServiceProvider) DeviceImpl(engine *gin.Engine) *device.Implementation {
	if s.deviceImpl == nil {
		s.deviceImpl = device.NewImplementation(engine, s.DeviceService())
	}

	return s.deviceImpl
}
