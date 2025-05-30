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
// @Router /part [get]
// @Success 200 {object} schemas.PartResponse
// @Param page query string false "Page number"
// @Param pageSize query string false "Page size"
func GetPartsController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	partsResponse, err := service.GetParts(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := gin.H{
		"parts":      partsResponse,
		"totalPages": c.GetInt("totalPages"),
		"page":       c.GetInt("page"),
		"pageSize":   c.GetInt("pageSize"),
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Get Part
// @Description Get Part
// @Tags Part
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /part/{id} [get]
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
// @Router /part [post]
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
// @Router /part/{id} [put]
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
// @Router /part/{id} [delete]
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
