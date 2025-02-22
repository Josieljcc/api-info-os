package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/Josieljcc/api-info-os/utils"
	"github.com/gin-gonic/gin"
)

// CreateClientController godoc
// @Summary Create a client
// @Description Create a client
// @Tags Client
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param client body schemas.ClientRegister true "Client"
// @Success 200 {object} schemas.ClientResponse
// @Router /client [post]
func CreateClientController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)

	if !isAuthorized {
		return
	}

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

// GetClientController godoc
// @Summary Get a client by id
// @Description Get a client by id
// @Tags Client
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param id path string true "Client ID"
// @Success 200 {object} schemas.ClientResponse
// @Router /client/{id} [get]
func GetClientController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)

	if !isAuthorized {
		return
	}

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

// GetClientsController godoc
// @Summary Get all clients
// @Description Get all clients
// @Tags Client
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param page query string false "Page number"
// @Param pageSize query string false "Page size"
// @Success 200 {object} schemas.CientResponseWithPagination
// @Router /client [get]
func GetClientsController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)

	if !isAuthorized {
		return
	}

	clientsResponse, err := service.GetClients(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := gin.H{
		"clients":    clientsResponse,
		"totalPages": c.GetInt("totalPages"),
		"page":       c.GetInt("page"),
		"pageSize":   c.GetInt("pageSize"),
	}

	c.JSON(http.StatusOK, response)
}

// UpdateClientController godoc
// @Summary Update a client
// @Description Update a client
// @Tags Client
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param id path string true "Client ID"
// @Param client body schemas.ClientRegister true "Client"
// @Success 200 {object} schemas.ClientResponse
func UpdateClientController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)

	if !isAuthorized {
		return
	}

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

// DeleteClientController godoc
// @Summary Delete a client
// @Description Delete a client
// @Tags Client
// @Accept json
// @Produce json
// @Param authorization header string true "Bearer Authorization"
// @Param id path string true "Client ID"
// @Success 200 {object} schemas.ClientResponse
// @Router /client/{id} [delete]
func DeleteClientController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)

	if !isAuthorized {
		return
	}

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
