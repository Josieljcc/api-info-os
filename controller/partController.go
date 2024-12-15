package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/Josieljcc/api-info-os/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Get Parts
// @Description Get Parts
// @Tags Part
// @Accept  json
// @Produce  json
// @Router /parts [get]
func GetPartsController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	partsResponse, err := service.GetParts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, partsResponse)
}

// @Summary Get Part
// @Description Get Part
// @Tags Part
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /parts/{id} [get]
func GetPartController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	id := c.Params.ByName("id")
	partResponse, err := service.GetPart(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, partResponse)
}

// @Summary Create Part
// @Description Create Part
// @Tags Part
// @Accept  json
// @Param   part    body  schemas.PartCreate  true "Part"
// @Param authorization header string true "Bearer Authorization"
// @Produce  json
// @Router /parts [post]
func CreatePartController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	var part schemas.Part
	if err := c.ShouldBindJSON(&part); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	partResponse, err := service.CreatePart(part)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, partResponse)
}

// @Summary Update Part
// @Description Update Part
// @Tags Part
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /parts/{id} [put]
func UpdatePartController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	id := c.Param("id")
	var part schemas.Part
	if err := c.ShouldBindJSON(&part); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	partResponse, err := service.UpdatePart(part, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, partResponse)
}

// @Summary Delete Part
// @Description Delete Part
// @Tags Part
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /parts/{id} [delete]
func DeletePartController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	id := c.Params.ByName("id")
	if err := service.DeletePart(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Part deleted",
	})
}
