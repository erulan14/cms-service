package device

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (i *Implementation) List(c *gin.Context) {
	devices, err := i.deviceService.List(c.Request.Context())
	log.Println(devices)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, devices)
	}
}
