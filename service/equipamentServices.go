package service

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
)

func GetEquipments() ([]schemas.EquipmentResponse, error) {
	db := config.GetDB()
	var equipments []schemas.Equipment
	if err := db.Find(&equipments).Error; err != nil {
		return nil, err
	}
	var equipmentsResponse []schemas.EquipmentResponse
	for _, equipment := range equipments {
		equipmentsResponse = append(equipmentsResponse, equipment.ToResponse())
	}
	return equipmentsResponse, nil
}

func GetEquipment(id string) (schemas.EquipmentResponse, error) {
	db := config.GetDB()
	var equipment schemas.Equipment
	if err := db.First(&equipment, id).Error; err != nil {
		return schemas.EquipmentResponse{}, err
	}
	return equipment.ToResponse(), nil
}

func CreateEquipment(equipment schemas.Equipment) (schemas.EquipmentResponse, error) {
	db := config.GetDB()

	if err := db.Create(&equipment).Error; err != nil {
		return schemas.EquipmentResponse{}, err
	}

	return equipment.ToResponse(), nil
}

func UpdateEquipment(equipment schemas.Equipment, id string) error {
	db := config.GetDB()
	var equipmentUpdated schemas.Equipment
	err := db.First(&equipmentUpdated, id).Error
	if err != nil {
		return err
	}
	if equipment.Name != "" {
		equipmentUpdated.Name = equipment.Name
	}
	if equipment.Description != "" {
		equipmentUpdated.Description = equipment.Description
	}
	if equipment.EquipamentModel != "" {
		equipmentUpdated.EquipamentModel = equipment.EquipamentModel
	}
	if equipment.SerialNumber != "" {
		equipmentUpdated.SerialNumber = equipment.SerialNumber
	}
	if err := db.Save(&equipmentUpdated).Error; err != nil {
		return err
	}
	return nil
}

func DeleteEquipment(id string) error {
	db := config.GetDB()
	if err := db.Delete(&schemas.Equipment{}, id).Error; err != nil {
		return err
	}
	return nil
}
