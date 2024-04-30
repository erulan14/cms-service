package geozone

import (
	"cms-service/internal/repository"
	def "cms-service/internal/service"
)

var _ def.GeozoneService = (*service)(nil)

type service struct {
	geozoneRepository repository.GeozoneRepository
}

func NewService(geozoneRepository repository.GeozoneRepository) *service {
	return &service{geozoneRepository: geozoneRepository}
}
