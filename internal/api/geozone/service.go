package geozone

import (
	"cms-service/internal/service"
	"github.com/gin-gonic/gin"
)

type Implementation struct {
	geozoneService service.GeozoneService
	engine         *gin.Engine
}

func NewImplementation(engine *gin.Engine, geozoneService service.GeozoneService) *Implementation {
	impl := &Implementation{geozoneService: geozoneService}

	group := engine.Group("/geozone")
	group.GET("/", impl.List)
	group.GET("/:id", impl.Get)
	group.POST("/", impl.Create)
	group.PUT("/:id", impl.Update)
	group.DELETE("/:id", impl.Delete)
	impl.engine = engine

	return impl
}
