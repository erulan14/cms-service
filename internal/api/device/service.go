package device

import (
	"cms-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Implementation struct {
	deviceService service.DeviceService
	engine        *gin.Engine
}

func NewImplementation(engine *gin.Engine, deviceService service.DeviceService) *Implementation {
	impl := &Implementation{deviceService: deviceService}

	group := engine.Group("/device")
	group.GET("/", impl.List)
	group.GET("/:id", impl.Get)
	group.POST("/", impl.Create)
	group.PUT("/:id", impl.Update)
	group.DELETE("/:id", impl.Delete)
	impl.engine = engine

	return impl
}
