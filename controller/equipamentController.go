package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/gin-gonic/gin"
)

// GetEquipmentbyId godoc
// @Summary Get a equipment by id
// @Description Get a equipment by id
// @Tags Equipment
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param id path string true "Equipment ID"
// @Success 200 {object} schemas.EquipmentResponse
// @Router /equipment/{id} [get]
func GetEquipmentbyIdController(c *gin.Context) {
	id := c.Param("id")
	equipment, err := service.GetEquipment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, equipment)
}

// GetEquipmentsController godoc
// @Summary Get all equipments
// @Description Get all equipments
// @Tags Equipment
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Success 200 {object} []schemas.EquipmentResponse
// @Router /equipment [get]
func GetEquipmentsController(c *gin.Context) {
	equipments, err := service.GetEquipments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, equipments)
}

// CreateEquipmentController godoc
// @Summary Create a equipment
// @Description Create a equipment
// @Tags Equipment
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param equipment body schemas.EquipmentRegister true "Equipment"
// @Success 200 {object} schemas.EquipmentResponse
// @Router /equipment [post]
func CreateEquipmentController(c *gin.Context) {
	var equipment schemas.Equipment
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	equipmentResponse, err := service.CreateEquipment(equipment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, equipmentResponse)
}

// UpdateEquipmentController godoc
// @Summary Update a equipment
// @Description Update a equipment
// @Tags Equipment
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param id path string true "Equipment ID"
// @Param equipment body schemas.EquipmentRegister true "Equipment"
// @Success 200 {object} schemas.EquipmentResponse
// @Router /equipment/{id} [put]
func UpdateEquipmentController(c *gin.Context) {
	var equipment schemas.Equipment
	id := c.Param("id")
	if err := c.ShouldBindJSON(&equipment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := service.UpdateEquipment(equipment, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Equipment updated successfully",
	})
}

// DeleteEquipmentController godoc
// @Summary Delete a equipment
// @Description Delete a equipment
// @Tags Equipment
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param id path string true "Equipment ID"
// @Success 200 {object} schemas.EquipmentResponse
// @Router /equipment/{id} [delete]
func DeleteEquipmentController(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteEquipment(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Equipment deleted successfully",
	})
}
