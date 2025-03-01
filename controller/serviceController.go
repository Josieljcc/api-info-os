package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/Josieljcc/api-info-os/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Get Services
// @Description Get Services
// @Tags Service
// @Accept  json
// @Produce  json
// @Router /services [get]
// @Success 200 {object} schemas.ServiceResponse
// @Param page query string false "Page number"
// @Param pageSize query string false "Page size"
func GetServicesController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	servicesResponse, err := service.GetServices(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := gin.H{
		"services":   servicesResponse,
		"totalPages": c.GetInt("totalPages"),
		"page":       c.GetInt("page"),
		"pageSize":   c.GetInt("pageSize"),
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Create Service
// @Description Create Service
// @Tags Service
// @Accept  json
// @Produce  json
// @Router /services [post]
func CreateServiceController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

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

// @Summary Get Service
// @Description Get Service
// @Tags Service
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /services/{id} [get]
func GetServiceController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

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

// @Summary Update Service
// @Description Update Service
// @Tags Service
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /services/{id} [put]
func UpdateServiceController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

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

// @Summary Delete Service
// @Description Delete Service
// @Tags Service
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /services/{id} [delete]
func DeleteServiceController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

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
