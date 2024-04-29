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

	device := engine.Group("/device")
	device.GET("/", impl.List)
	device.GET("/:id", impl.List)
	device.POST("/", impl.List)
	device.PUT("/:id", impl.List)
	device.DELETE("/", impl.List)
	impl.engine = engine

	return impl
}
