package app

import (
	"cms-service/internal/api/geozone"
	"cms-service/internal/repository"
	geozoneRepository "cms-service/internal/repository/geozone"
	"cms-service/internal/service"
	geozoneService "cms-service/internal/service/geozone"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type geozoneServiceProvider struct {
	geozoneRepository repository.GeozoneRepository
	geozoneService    service.GeozoneService
	geozoneImpl       *geozone.Implementation
	db                *pgxpool.Pool
}

func newGeozoneServiceProvider(db *pgxpool.Pool) *geozoneServiceProvider {
	var service geozoneServiceProvider

	service.db = db

	return &service
}

func (s *geozoneServiceProvider) GeozoneRepository() repository.GeozoneRepository {
	if s.geozoneRepository == nil {
		s.geozoneRepository = geozoneRepository.NewRepository(s.db)
	}
	return s.geozoneRepository
}

func (s *geozoneServiceProvider) GeozoneService() service.GeozoneService {
	if s.geozoneService == nil {
		s.geozoneService = geozoneService.NewService(
			s.GeozoneRepository(),
		)
	}

	return s.geozoneService
}

func (s *geozoneServiceProvider) GeozoneImpl(engine *gin.Engine) *geozone.Implementation {
	if s.geozoneImpl == nil {
		s.geozoneImpl = geozone.NewImplementation(engine, s.GeozoneService())
	}

	return s.geozoneImpl
}
