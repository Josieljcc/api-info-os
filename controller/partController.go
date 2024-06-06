package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/gin-gonic/gin"
)

func GetPartsController(c *gin.Context) {
	partsResponse, err := service.GetParts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, partsResponse)
}

func GetPartController(c *gin.Context) {
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

func CreatePartController(c *gin.Context) {
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

func UpdatePartController(c *gin.Context) {
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

func DeletePartController(c *gin.Context) {
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
