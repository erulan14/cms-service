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
	device.GET("/:id", impl.Get)
	device.POST("/", impl.Create)
	device.PUT("/:id", impl.Update)
	device.DELETE("/", impl.Delete)
	impl.engine = engine

	return impl
}
