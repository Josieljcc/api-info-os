package service

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/utils"
	"github.com/gin-gonic/gin"
)

func GetTechnicians(c *gin.Context) ([]schemas.TechnicianResponse, error) {
	db := config.GetDB()
	var technicians []schemas.TechnicianResponse
	err := db.Scopes(Paginate(c)).Find(&technicians).Error

	c.Next()
	return technicians, err
}

func GetTechnician(id string) (schemas.Technician, error) {
	db := config.GetDB()
	var technician schemas.Technician
	err := db.First(&technician, id).Error
	return technician, err
}

func GetTechnicianByEmail(email string) (schemas.Technician, error) {
	db := config.GetDB()
	var technician schemas.Technician
	err := db.Where("email = ?", email).First(&technician).Error
	return technician, err
}

func CreateTechnician(technician schemas.Technician) (technicianResponse schemas.TechnicianResponse, err error) {
	db := config.GetDB()
	hashedPassword, err := utils.HashPassword(technician.Password)
	if err != nil {
		return schemas.TechnicianResponse{}, err
	}
	technician.Password = string(hashedPassword)
	if err := db.Create(&technician).Error; err != nil {
		return schemas.TechnicianResponse{}, err
	}

	return technician.ToResponse(), nil
}

func DeleteTechnician(id string) error {
	db := config.GetDB()
	err := db.Delete(&schemas.Technician{}, id).Error
	return err
}

func UpdateTechnician(technician schemas.Technician, id string) error {
	db := config.GetDB()
	var technicianUpdated schemas.Technician
	err := db.First(&technicianUpdated, id).Error
	if err != nil {
		return err
	}
	if technician.Name != "" {
		technicianUpdated.Name = technician.Name
	}
	if technician.Password != "" {
		hashedPassword, err := utils.HashPassword(technician.Password)
		if err != nil {
			return err
		}
		technicianUpdated.Password = string(hashedPassword)
	}
	if technician.Email != "" {
		technicianUpdated.Email = technician.Email
	}
	if technician.Phone != "" {
		technicianUpdated.Phone = technician.Phone
	}

	if err := db.Save(&technicianUpdated).Error; err != nil {
		return err
	}
	return nil
}
