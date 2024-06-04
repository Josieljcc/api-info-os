package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/gin-gonic/gin"
)

func CreateClientController(c *gin.Context) {
	var client schemas.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var db = config.GetDB()
	if err := db.Create(&client).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"client": client,
	})
}

func GetClientController(c *gin.Context) {
	id := c.Params.ByName("id")
	client, err := service.GetClient(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"client": client,
	})
}

func GetClientsController(c *gin.Context) {
	clients, err := service.GetClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"clients": clients,
	})
}

func EditClientController(c *gin.Context) {
	var client schemas.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	client, err := service.UpdateClient(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"client": client,
	})
}

func DeleteClientController(c *gin.Context) {
	id := c.Params.ByName("id")
	if err := service.DeleteClient(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Client deleted",
	})
}
