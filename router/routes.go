package router

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/controller"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	v1.GET("/client/:id", controller.GetClientController)
	v1.POST("/client", controller.CreateClientController)
	v1.DELETE("/client/:id", controller.DeleteClientController)
	v1.PUT("/client", controller.EditClientController)
	v1.GET("/clients", controller.GetClientsController)
}
