package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/gin-gonic/gin"
)

// @Summary Get Technicians
// @Description Get Technicians
// @Tags Technician
// @Accept  json
// @Produce  json
// @Router /technicians [get]
// @Success 200 {object} schemas.TechnicianResponse
// @Param page query string false "Page number"
// @Param pageSize query string false "Page size"
func GetTechniciansController(c *gin.Context) {
	technicians, err := service.GetTechnicians(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := gin.H{
		"technicians": technicians,
		"totalPages":  c.GetInt("totalPages"),
		"page":        c.GetInt("page"),
		"pageSize":    c.GetInt("pageSize"),
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Technician
// @Description Get Technician
// @Tags Technician
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /technicians/{id} [get]
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

// @Summary Create Technician
// @Description Create Technician
// @Tags Technician
// @Accept  json
// @Produce  json
// @Router /register/technician [post]
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

// @Summary Update Technician
// @Description Update Technician
// @Tags Technician
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /technicians/{id} [put]
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

// @Summary Delete Technician
// @Description Delete Technician
// @Tags Technician
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /technicians/{id} [delete]
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
