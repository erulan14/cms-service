package device

import (
	"cms-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) Update(c *gin.Context) {
	uuid, _ := c.Params.Get("id")
	if uuid == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var dto model.UpdateDeviceDTO
	err := c.BindJSON(&dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = i.deviceService.Update(c.Request.Context(), uuid, &dto)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
