package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClientHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"client": "created client",
	})
}

func GetClientHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"client": "client data",
	})
}

func GetClientsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"clients": "client data",
	})
}

func EditClientHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"client": "edited client",
	})
}

func DeleteClientHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"client": "deleted client",
	})
}
