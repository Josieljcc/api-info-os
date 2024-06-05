package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/gin-gonic/gin"
)

func GetTechnicianController(c *gin.Context) {
	id := c.Param("id")
	technician, err := service.GetTechnician(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, technician.ToResponse())
}

func CreateTechnicianController(c *gin.Context) {
	var technician schemas.Technician
	if err := c.ShouldBindJSON(&technician); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	technicianResponse, err := service.CreateTechnician(technician)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, technicianResponse)
}

func UpdateTechnicianController(c *gin.Context) {
	id := c.Param("id")
	var technician schemas.Technician
	if err := c.ShouldBindJSON(&technician); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := service.UpdateTechnician(technician, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "Technician updated",
	})
}

func DeleteTechnicianController(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteTechnician(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Technician deleted",
	})
}

func GetTechniciansController(c *gin.Context) {
	technicians, err := service.GetTechnicians()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var techniciansResponse []schemas.TechnicianResponse
	for _, technician := range technicians {
		techniciansResponse = append(techniciansResponse, technician.ToResponse())
	}
	c.JSON(http.StatusOK, techniciansResponse)
}
