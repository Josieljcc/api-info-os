package controller

import (
	"net/http"

	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/service"
	"github.com/gin-gonic/gin"
)

// @Summary Get Order
// @Description Get Order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   id    path  string  true "ID"
// @Router /orders/{id} [get]
func GetOrderController(c *gin.Context) {
	id := c.Param("id")
	order, err := service.GetOrder(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

// @Summary Get Orders
// @Description Get Orders
// @Tags Order
// @Accept  json
// @Produce  json
// @Router /orders [get]
func GetOrdersController(c *gin.Context) {
	ordersResponse, err := service.GetOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, ordersResponse)
}

// @Summary Create Order
// @Description Create Order
// @Tags Order
// @Accept  json
// @Produce  json
// @Param   order    body  schemas.Order  true "Order"
// @Router /orders [post]
func CreateOrderController(c *gin.Context) {
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
// @Param   order    body  schemas.Order  true "Order"
// @Router /orders/{id} [put]
func UpdateOrderController(c *gin.Context) {
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
