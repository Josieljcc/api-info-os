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

	v1.POST("/login", controller.LoginController)

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

	v1.GET("/part/:id", controller.GetPartController)
	v1.POST("/part", controller.CreatePartController)
	v1.DELETE("/part/:id", controller.DeletePartController)
	v1.PUT("/part/:id", controller.UpdatePartController)
	v1.GET("/parts", controller.GetPartsController)

	v1.GET("/order/:id", controller.GetOrderController)
	v1.POST("/order", controller.CreateOrderController)
	v1.GET("/orders", controller.GetOrdersController)
	v1.DELETE("/order/:id", controller.DeleteOrderController)
	v1.PUT("/order/:id", controller.UpdateOrderController)

	v1.GET("/service/:id", controller.GetServiceController)
	v1.POST("/service", controller.CreateServiceController)
	v1.GET("/services", controller.GetServicesController)
	v1.PUT("/service/:id", controller.UpdateServiceController)
	v1.DELETE("/service/:id", controller.DeleteServiceController)
}
