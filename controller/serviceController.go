package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/gin-gonic/gin"
)

func GetServicesController(c *gin.Context) {
	servicesResponse, err := service.GetServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, servicesResponse)
}

func CreateServiceController(c *gin.Context) {
	var serviceBody schemas.Service
	if err := c.ShouldBindJSON(&serviceBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	serviceResponse, err := service.CreateService(serviceBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, serviceResponse)
}

func GetServiceController(c *gin.Context) {
	id := c.Param("id")
	service, err := service.GetService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, service)
}

func UpdateServiceController(c *gin.Context) {
	id := c.Param("id")
	var serviceUpdated schemas.Service
	if err := c.ShouldBindJSON(&serviceUpdated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	serviceResponse, err := service.UpdateService(serviceUpdated, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, serviceResponse)
}

func DeleteServiceController(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteService(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Service deleted",
	})
}
