package router

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	v1.GET("/client", handler.GetClientHandler)
	v1.POST("/client", handler.CreateClientHandler)
	v1.DELETE("/client", handler.DeleteClientHandler)
	v1.PUT("/client", handler.EditClientHandler)
	v1.GET("/clients", handler.GetClientsHandler)
}
