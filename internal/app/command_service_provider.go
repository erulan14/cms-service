package app

import (
	"cms-service/internal/api/command"
	"cms-service/internal/repository"
	commandRepository "cms-service/internal/repository/command"
	"cms-service/internal/service"
	commandService "cms-service/internal/service/command"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type commandServiceProvider struct {
	commandRepository repository.CommandRepository
	commandService    service.CommandService
	commandImpl       *command.Implementation
	db                *pgxpool.Pool
}

func newCommandServiceProvider(db *pgxpool.Pool) *commandServiceProvider {
	var service commandServiceProvider

	service.db = db

	return &service
}

func (s *commandServiceProvider) CommandRepository() repository.CommandRepository {
	if s.commandRepository == nil {
		s.commandRepository = commandRepository.NewRepository(s.db)
	}
	return s.commandRepository
}

func (s *commandServiceProvider) CommandService() service.CommandService {
	if s.commandService == nil {
		s.commandService = commandService.NewService(
			s.CommandRepository(),
		)
	}

	return s.commandService
}

func (s *commandServiceProvider) CommandImpl(engine *gin.Engine) *command.Implementation {
	if s.commandImpl == nil {
		s.commandImpl = command.NewImplementation(engine, s.CommandService())
	}

	return s.commandImpl
}
