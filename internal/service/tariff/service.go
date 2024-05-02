package tariff

import (
	"cms-service/internal/repository"
	def "cms-service/internal/service"
)

var _ def.TariffService = (*service)(nil)

type service struct {
	tariffRepository repository.TariffRepository
}

func NewService(tariffRepository repository.TariffRepository) *service {
	return &service{tariffRepository: tariffRepository}
}
