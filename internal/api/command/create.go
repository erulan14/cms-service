package command

import (
	"cms-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) Create(c *gin.Context) {
	var dto model.CreateCommandDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uuid, err := i.commandService.Create(c.Request.Context(), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"uuid": uuid})
}
