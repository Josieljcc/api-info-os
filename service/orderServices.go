package service

import (
	"errors"
	"time"

	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/gin-gonic/gin"
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

	err := query.Preload("Client").Preload("Technician").Preload("Services").Preload("Parts").Find(&orders).Error

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
	if err := db.Preload("Client").Preload("Technician").Preload("Services").Preload("Parts").First(&order, id).Error; err != nil {
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
	err := db.Preload("Client").Preload("Technician").Preload("Services").Preload("Parts").First(&orderUpdated, id).Error
	if err != nil {
		return schemas.OrderResponse{}, err
	}
	if order.Comment != "" {
		orderUpdated.Comment = order.Comment
	}
	if order.ClientID != 0 {
		orderUpdated.ClientID = order.ClientID
	}
	if order.Status != "" {
		// Valida se o status é válido
		if !order.Status.IsValid() {
			return schemas.OrderResponse{}, errors.New("invalid status")
		}
		orderUpdated.Status = order.Status
	}
	if order.TechnicianID != 0 {
		orderUpdated.TechnicianID = order.TechnicianID
	}
	if !order.OpeningDate.IsZero() {
		orderUpdated.OpeningDate = order.OpeningDate
	}
	if !order.ForecastDate.IsZero() {
		orderUpdated.ForecastDate = order.ForecastDate
	}
	if order.ClosingDate != nil {
		orderUpdated.ClosingDate = order.ClosingDate
	}
	if err := db.Save(&orderUpdated).Error; err != nil {
		return schemas.OrderResponse{}, err
	}
	return orderUpdated.ToResponse(), nil
}

func CloseOrder(id string) (schemas.OrderResponse, error) {
	db := config.GetDB()
	var order schemas.Order

	// Busca a ordem
	err := db.Preload("Client").Preload("Technician").Preload("Services").Preload("Parts").First(&order, id).Error
	if err != nil {
		return schemas.OrderResponse{}, err
	}

	// Verifica se a ordem já está fechada
	if order.Status == schemas.StatusCompleted {
		return schemas.OrderResponse{}, errors.New("ordem de serviço já está fechada")
	}

	// Atualiza o status para concluída e define a data de fechamento
	now := time.Now()
	order.Status = schemas.StatusCompleted
	order.ClosingDate = &now

	// Salva as alterações
	if err := db.Save(&order).Error; err != nil {
		return schemas.OrderResponse{}, err
	}

	return order.ToResponse(), nil
}

func DeleteOrder(id string) error {
	db := config.GetDB()
	if err := db.Delete(&schemas.Order{}, id).Error; err != nil {
		return err
	}
	return nil
}
