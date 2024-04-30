package geozone

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) List(c *gin.Context) {
	devices, err := i.geozoneService.List(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, devices)
	}
}
