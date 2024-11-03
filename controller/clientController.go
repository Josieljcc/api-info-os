package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/gin-gonic/gin"
)

// CreateClient godoc
// @Summary Create a new client
// @Description Create a new client
// @Tags Client
// @Accept json
// @Produce json
// @Param client body schemas.ClientRegister true "Client"
// @Success 200 {object} schemas.ClientResponse
// @Router /register/client [post]
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

// GetClient godoc
// @Summary Get a client
// @Description Get a client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "Client ID"
// @Success 200 {object} schemas.ClientResponse
// @Router /clients/{id} [get]
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

// GetClients godoc
// @Summary Get all clients
// @Description Get all clients
// @Tags Client
// @Accept json
// @Produce json
// @Success 200 {object} []schemas.ClientResponse
// @Router /clients [get]
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

// UpdateClient godoc
// @Summary Update a client
// @Description Update a client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "Client ID"
// @Param client body schemas.ClientRegister true "Client"
// @Success 200 {object} schemas.ClientResponse
// @Router /clients/{id} [put]
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

// DeleteClient godoc
// @Summary Delete a client
// @Description Delete a client
// @Tags Client
// @Accept json
// @Produce json
// @Param id path string true "Client ID"
// @Success 200 {object} schemas.ClientResponse
// @Router /clients/{id} [delete]
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
