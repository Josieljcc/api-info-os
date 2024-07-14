package router

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/controller"
	"github.com/Josieljcc/api-info-os/middleware"
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

	authRoutes := v1.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())

	authRoutes.GET("/client/:id", controller.GetClientController)
	authRoutes.POST("/client", controller.CreateClientController)
	authRoutes.DELETE("/client/:id", controller.DeleteClientController)
	authRoutes.PUT("/client/:id", controller.UpdateClientController)
	authRoutes.GET("/clients", controller.GetClientsController)

	authRoutes.GET("/technician/:id", controller.GetTechnicianController)
	authRoutes.POST("/technician", controller.CreateTechnicianController)
	authRoutes.DELETE("/technician/:id", controller.DeleteTechnicianController)
	authRoutes.PUT("/technician/:id", controller.UpdateTechnicianController)
	authRoutes.GET("/technicians", controller.GetTechniciansController)

	authRoutes.GET("/part/:id", controller.GetPartController)
	authRoutes.POST("/part", controller.CreatePartController)
	authRoutes.DELETE("/part/:id", controller.DeletePartController)
	authRoutes.PUT("/part/:id", controller.UpdatePartController)
	authRoutes.GET("/parts", controller.GetPartsController)

	authRoutes.GET("/order/:id", controller.GetOrderController)
	authRoutes.POST("/order", controller.CreateOrderController)
	authRoutes.GET("/orders", controller.GetOrdersController)
	authRoutes.DELETE("/order/:id", controller.DeleteOrderController)
	authRoutes.PUT("/order/:id", controller.UpdateOrderController)

	authRoutes.GET("/service/:id", controller.GetServiceController)
	authRoutes.POST("/service", controller.CreateServiceController)
	authRoutes.GET("/services", controller.GetServicesController)
	authRoutes.PUT("/service/:id", controller.UpdateServiceController)
	authRoutes.DELETE("/service/:id", controller.DeleteServiceController)
}
