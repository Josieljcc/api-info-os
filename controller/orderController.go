package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/Josieljcc/api-info-os/utils"
	"github.com/gin-gonic/gin"
)

// @Summary Get Order
// @Description Get Order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Param authorization header string true "Bearer Authorization"
// @Router /orders/{id} [get]
func GetOrderController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	id := c.Param("id")
	order, err := service.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order.ToResponse())
}

// @Summary Get Orders
// @Description Get Orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Router /orders [get]
// @Success 200 {object} schemas.OrderResponse
// @Param page query string false "Page number"
// @Param pageSize query string false "Page size"
func GetOrdersController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	ordersResponse, err := service.GetOrders(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := gin.H{
		"orders":     ordersResponse,
		"totalPages": c.GetInt("totalPages"),
		"page":       c.GetInt("page"),
		"pageSize":   c.GetInt("pageSize"),
	}

	c.JSON(http.StatusOK, response)
}

// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept  json
// @Produce  json
// @Router /orders [post]
func CreateOrderController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	var order schemas.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	orderResponse, err := service.CreateOrder(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, orderResponse)
}

// @Summary Update Order
// @Description Update Order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /orders/{id} [put]
func UpdateOrderController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	id := c.Param("id")
	var order schemas.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	orderResponse, err := service.UpdateOrder(order, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orderResponse)
}

// @Summary Delete Order
// @Description Delete Order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /orders/{id} [delete]
func DeleteOrderController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	id := c.Param("id")
	err := service.DeleteOrder(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted",
	})
}

// @Summary Close Order
// @Description Close an order by setting status to 'completed' and adding closing date
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "Order ID"
// @Param authorization header string true "Bearer Authorization"
// @Success 200 {object} schemas.OrderResponse
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /orders/{id}/close [patch]
func CloseOrderController(c *gin.Context) {
	isAuthorized := utils.VerifyRole(c)
	if !isAuthorized {
		return
	}

	id := c.Param("id")
	orderResponse, err := service.CloseOrder(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, orderResponse)
}
