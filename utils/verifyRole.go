package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyRole(c *gin.Context) bool {
	role := c.GetString("role")
	if role == "client" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "You are not authorized to perform this action",
		})
		return false
	}
	return true
}
