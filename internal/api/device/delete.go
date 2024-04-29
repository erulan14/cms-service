package device

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) Delete(c *gin.Context) {
	uuid, _ := c.Params.Get("id")
	if uuid == "" {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err := i.deviceService.Delete(c.Request.Context(), uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
