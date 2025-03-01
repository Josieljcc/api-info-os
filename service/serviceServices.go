package service

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) ([]schemas.ServiceResponse, error) {
	db := config.GetDB()
	var services []schemas.Service
	if err := db.Scopes(Paginate(c)).Preload("Orders").Find(&services).Error; err != nil {
		return nil, err
	}
	var servicesResponse []schemas.ServiceResponse
	for _, service := range services {
		servicesResponse = append(servicesResponse, service.ToResponse())
	}

	c.Next()
	return servicesResponse, nil
}

func GetService(id string) (schemas.ServiceResponse, error) {
	db := config.GetDB()
	var service schemas.Service
	if err := db.Preload("Orders").First(&service, id).Error; err != nil {
		return schemas.ServiceResponse{}, err
	}
	return service.ToResponse(), nil
}

func CreateService(service schemas.Service) (schemas.ServiceResponse, error) {
	db := config.GetDB()
	if err := db.Create(&service).Error; err != nil {
		return schemas.ServiceResponse{}, err
	}
	return service.ToResponse(), nil
}

func UpdateService(service schemas.Service, id string) (schemas.ServiceResponse, error) {
	db := config.GetDB()
	var serviceUpdated schemas.Service
	err := db.Preload("Orders").First(&serviceUpdated, id).Error
	if err != nil {
		return schemas.ServiceResponse{}, err
	}
	if service.Description != "" {
		serviceUpdated.Description = service.Description
	}
	if service.Price != 0 {
		serviceUpdated.Price = service.Price
	}
	if service.Time != 0 {
		serviceUpdated.Time = service.Time
	}
	if err := db.Save(&serviceUpdated).Error; err != nil {
		return schemas.ServiceResponse{}, err
	}
	return serviceUpdated.ToResponse(), nil
}

func DeleteService(id string) error {
	db := config.GetDB()
	if err := db.Delete(&schemas.Service{}, id).Error; err != nil {
		return err
	}
	return nil
}
