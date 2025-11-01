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

	if c.Query("clientName") != "" {
		query = query.Joins("Client").Where("clients.name LIKE ?", "%"+c.Query("clientName")+"%")
	}

	if c.Query("status") != "" {
		query = query.Where("status = ?", c.Query("status"))
	}

	if c.Query("technicianID") != "" {
		query = query.Where("technician_id = ?", c.Query("technicianID"))
	}
	if c.Query("openingDate") != "" {
		openingDate, err := parseQueryDate(c.Query("openingDate"))
		if err != nil {
			return nil, err
		}
		query = query.Where("opening_date = ?", openingDate)
	}

	if c.Query("openingStartDate") != "" && c.Query("openingEndDate") != "" {
		openingStartDate, err := parseQueryDate(c.Query("openingStartDate"))
		if err != nil {
			return nil, err
		}
		openingEndDate, err := parseQueryDate(c.Query("openingEndDate"))
		if err != nil {
			return nil, err
		}
		query = query.Where("opening_date BETWEEN ? AND ?", openingStartDate, openingEndDate)
	}

	if c.Query("forecastDate") != "" {
		forecastDate, err := parseQueryDate(c.Query("forecastDate"))
		if err != nil {
			return nil, err
		}
		query = query.Where("forecast_date = ?", forecastDate)
	}

	if c.Query("forecastStartDate") != "" && c.Query("forecastEndDate") != "" {
		forecastStartDate, err := parseQueryDate(c.Query("forecastStartDate"))
		if err != nil {
			return nil, err
		}
		forecastEndDate, err := parseQueryDate(c.Query("forecastEndDate"))
		if err != nil {
			return nil, err
		}
		query = query.Where("forecast_date BETWEEN ? AND ?", forecastStartDate, forecastEndDate)
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

	err := db.Preload("Client").Preload("Technician").Preload("Services").Preload("Parts").First(&order, id).Error
	if err != nil {
		return schemas.OrderResponse{}, err
	}

	if order.Status == schemas.StatusCompleted {
		return schemas.OrderResponse{}, errors.New("ordem de serviço já está fechada")
	}

	now := time.Now()
	order.Status = schemas.StatusCompleted
	order.ClosingDate = &now

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

func parseQueryDate(value string) (time.Time, error) {
	layouts := []string{
		"2006-01-02",
		time.RFC3339,
		time.RFC3339Nano,
	}

	var err error
	for _, layout := range layouts {
		var parsed time.Time
		parsed, err = time.Parse(layout, value)
		if err == nil {
			return parsed, nil
		}
	}

	return time.Time{}, err
}
