package tariff

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) List(c *gin.Context) {
	tariffs, err := i.tariffService.List(c.Request.Context())

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, tariffs)
	}
}
