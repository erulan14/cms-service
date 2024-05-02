package app

import (
	"cms-service/internal/api/tariff"
	"cms-service/internal/repository"
	tariffRepository "cms-service/internal/repository/tariff"
	"cms-service/internal/service"
	tariffService "cms-service/internal/service/tariff"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type tariffServiceProvider struct {
	tariffRepository repository.TariffRepository
	tariffService    service.TariffService
	tariffImpl       *tariff.Implementation
	db               *pgxpool.Pool
}

func newTariffServiceProvider(db *pgxpool.Pool) *tariffServiceProvider {
	var service tariffServiceProvider

	service.db = db

	return &service
}

func (s *tariffServiceProvider) TariffRepository() repository.TariffRepository {
	if s.tariffRepository == nil {
		s.tariffRepository = tariffRepository.NewRepository(s.db)
	}
	return s.tariffRepository
}

func (s *tariffServiceProvider) TariffService() service.TariffService {
	if s.tariffService == nil {
		s.tariffService = tariffService.NewService(
			s.TariffRepository(),
		)
	}

	return s.tariffService
}

func (s *tariffServiceProvider) TariffImpl(engine *gin.Engine) *tariff.Implementation {
	if s.tariffImpl == nil {
		s.tariffImpl = tariff.NewImplementation(engine, s.TariffService())
	}

	return s.tariffImpl
}
