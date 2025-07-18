package router

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/controller"
	"github.com/Josieljcc/api-info-os/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1.POST("/login", controller.LoginController)
	v1.POST("/register/technician", controller.CreateTechnicianController)

	var clientPatch = "/client"
	var technicianPatch = "/technician"
	var partPatch = "/part"
	var orderPatch = "/order"
	var servicePatch = "/service"
	var equipmentPatch = "/equipment"

	authRoutes := v1.Group("/")
	authRoutes.Use(middleware.AuthMiddleware())

	authRoutes.GET(clientPatch+"/:id", controller.GetClientController)
	authRoutes.DELETE(clientPatch+"/:id", controller.DeleteClientController)
	authRoutes.PUT(clientPatch+"/:id", controller.UpdateClientController)
	authRoutes.GET(clientPatch, controller.GetClientsController)
	authRoutes.POST(clientPatch, controller.CreateClientController)

	authRoutes.GET(technicianPatch+"/:id", controller.GetTechnicianController)
	authRoutes.DELETE(technicianPatch+"/:id", controller.DeleteTechnicianController)
	authRoutes.PUT(technicianPatch+"/:id", controller.UpdateTechnicianController)
	authRoutes.GET(technicianPatch, controller.GetTechniciansController)

	authRoutes.GET(partPatch+"/:id", controller.GetPartController)
	authRoutes.POST(partPatch, controller.CreatePartController)
	authRoutes.DELETE(partPatch+"/:id", controller.DeletePartController)
	authRoutes.PUT(partPatch+"/:id", controller.UpdatePartController)
	authRoutes.GET(partPatch, controller.GetPartsController)

	authRoutes.GET(orderPatch+"/:id", controller.GetOrderController)
	authRoutes.POST(orderPatch, controller.CreateOrderController)
	authRoutes.DELETE(orderPatch+"/:id", controller.DeleteOrderController)
	authRoutes.PUT(orderPatch+"/:id", controller.UpdateOrderController)
	authRoutes.GET(orderPatch, controller.GetOrdersController)

	authRoutes.GET(servicePatch+"/:id", controller.GetServiceController)
	authRoutes.POST(servicePatch, controller.CreateServiceController)
	authRoutes.PUT(servicePatch+"/:id", controller.UpdateServiceController)
	authRoutes.DELETE(servicePatch+"/:id", controller.DeleteServiceController)
	authRoutes.GET(servicePatch, controller.GetServicesController)

	authRoutes.GET(equipmentPatch+"/:id", controller.GetEquipmentbyIdController)
	authRoutes.POST(equipmentPatch, controller.CreateEquipmentController)
	authRoutes.PUT(equipmentPatch+"/:id", controller.UpdateEquipmentController)
	authRoutes.DELETE(equipmentPatch+"/:id", controller.DeleteEquipmentController)
	authRoutes.GET(equipmentPatch, controller.GetEquipmentsController)

	authRoutes.GET("/backup", controller.BackupDBController)
	authRoutes.POST("/restore", controller.RestoreDBController)
}
