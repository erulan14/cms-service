package tariff

import (
	"cms-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Implementation struct {
	tariffService service.TariffService
	engine        *gin.Engine
}

func NewImplementation(engine *gin.Engine, tariffService service.TariffService) *Implementation {
	impl := &Implementation{tariffService: tariffService}

	group := engine.Group("/tariff")
	group.GET("/", impl.List)
	group.GET("/:id", impl.Get)
	group.POST("/", impl.Create)
	group.PUT("/:id", impl.Update)
	group.DELETE("/:id", impl.Delete)
	impl.engine = engine

	return impl
}
