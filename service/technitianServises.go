package service

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
)

func GetTechnician(id string) (schemas.Technician, error) {
	db := config.GetDB()
	var technician schemas.Technician
	err := db.First(&technician, id).Error
	return technician, err
}

func GetTechnicians() ([]schemas.Technician, error) {
	db := config.GetDB()
	var technicians []schemas.Technician
	err := db.Find(&technicians).Error
	return technicians, err
}

func CreateTechnician(technician schemas.Technician) error {
	db := config.GetDB()
	err := db.Create(&technician).Error
	return err
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
	technicianUpdated.Name = technician.Name
	technicianUpdated.Password = technician.Password
	technicianUpdated.Email = technician.Email
	technicianUpdated.Phone = technician.Phone
	if err := db.Save(&technicianUpdated).Error; err != nil {
		return err
	}
	return nil
}
