package service

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func GetOrders(c *gin.Context) ([]schemas.OrderResponse, error) {
	db := config.GetDB()
	var orders []schemas.Order

	query := db.Scopes(Paginate(c))

	if c.Query("clientID") != "" {
		query = query.Where("client_id = ?", c.Query("clientID"))
	}

	if c.Query("status") != "" {
		query = query.Where("status = ?", c.Query("status"))
	}

	if c.Query("technicianID") != "" {
		query = query.Where("technician_id = ?", c.Query("technicianID"))
	}

	err := query.Preload(clause.Associations).Find(&orders).Error

	if err != nil {
		return nil, err
	}

	var orderResponse []schemas.OrderResponse
	for _, order := range orders {
		orderResponse = append(orderResponse, order.ToResponse())
	}

	c.Next()
	return orderResponse, nil
}

func GetOrder(id string) (schemas.Order, error) {
	db := config.GetDB()
	var order schemas.Order
	if err := db.Preload(clause.Associations).First(&order, id).Error; err != nil {
		return schemas.Order{}, err
	}
	return order, nil
}

func CreateOrder(order schemas.Order) (schemas.OrderResponse, error) {
	db := config.GetDB()
	if err := db.Create(&order).Error; err != nil {
		return schemas.OrderResponse{}, err
	}
	return order.ToResponse(), nil
}

func UpdateOrder(order schemas.Order, id string) (schemas.OrderResponse, error) {
	db := config.GetDB()
	var orderUpdated schemas.Order
	err := db.Preload(clause.Associations).First(&orderUpdated, id).Error
	if err != nil {
		return schemas.OrderResponse{}, err
	}
	if order.Comment != "" {
		orderUpdated.Comment = order.Comment
	}
	if order.ClientID != "" {
		orderUpdated.ClientID = order.ClientID
	}
	if order.Status != "" {
		orderUpdated.Status = order.Status
	}
	if order.TechnicianID != "" {
		orderUpdated.TechnicianID = order.TechnicianID
	}
	if order.Date != "" {
		orderUpdated.Date = order.Date
	}
	if err := db.Save(&orderUpdated).Error; err != nil {
		return schemas.OrderResponse{}, err
	}
	return orderUpdated.ToResponse(), nil
}

func DeleteOrder(id string) error {
	db := config.GetDB()
	if err := db.Delete(&schemas.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}
