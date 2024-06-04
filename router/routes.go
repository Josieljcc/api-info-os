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
	v1.PUT("/client/:id", controller.UpdateClientController)
	v1.GET("/clients", controller.GetClientsController)

	v1.GET("/technician/:id", controller.GetTechnicianController)
	v1.POST("/technician", controller.CreateTechnicianController)
	v1.DELETE("/technician/:id", controller.DeleteTechnicianController)
	v1.PUT("/technician/:id", controller.UpdateTechnicianController)
	v1.GET("/technicians", controller.GetTechniciansController)

}
