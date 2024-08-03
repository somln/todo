package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo/models"
)

func ValidateTodoStatus(c *gin.Context, status models.TodoStatus) bool {
	if err := status.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}
