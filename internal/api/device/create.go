package device

import (
	"cms-service/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (i *Implementation) Create(c *gin.Context) {
	var dto model.CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	create, err := i.deviceService.Create(c.Request.Context(), &dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"uuid": create})
}
