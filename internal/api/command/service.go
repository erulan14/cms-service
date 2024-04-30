package command

import (
	"cms-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Implementation struct {
	commandService service.CommandService
	engine         *gin.Engine
}

func NewImplementation(engine *gin.Engine, commandService service.CommandService) *Implementation {
	impl := &Implementation{commandService: commandService}

	group := engine.Group("/command")
	group.GET("/", impl.List)
	group.GET("/:id", impl.Get)
	group.POST("/", impl.Create)
	group.PUT("/:id", impl.Update)
	group.DELETE("/:id", impl.Delete)
	impl.engine = engine

	return impl
}
