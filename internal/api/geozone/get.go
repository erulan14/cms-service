package geozone

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) Get(c *gin.Context) {
	uuid, _ := c.Params.Get("id")
	if uuid == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	get, err := i.geozoneService.Get(c.Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, get)
}
