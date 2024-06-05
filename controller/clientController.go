package controller

import (
	"net/http"

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
	clientResponse, err := service.CreateClient(client)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(http.StatusOK, clientResponse)
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
	c.JSON(http.StatusOK, client)
}

func GetClientsController(c *gin.Context) {
	clientsResponse, err := service.GetClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, clientsResponse)
}

func UpdateClientController(c *gin.Context) {
	var client schemas.Client
	id := c.Param("id")
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := service.UpdateClient(client, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"client": "client updated",
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
