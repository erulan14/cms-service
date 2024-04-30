package device

import (
	"cms-service/internal/repository"
	def "cms-service/internal/service"
)

var _ def.CommandService = (*service)(nil)

type service struct {
	commandRepository repository.CommandRepository
}

func NewService(commandRepository repository.CommandRepository) *service {
	return &service{commandRepository: commandRepository}
}
