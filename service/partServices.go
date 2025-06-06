package service

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/gin-gonic/gin"
)

func GetParts(c *gin.Context) ([]schemas.PartResponse, error) {
	db := config.GetDB()
	var parts []schemas.Part

	query := db.Scopes(Paginate(c))

	if c.Query("name") != "" {
		query = query.Where("name LIKE ?", "%"+c.Query("name")+"%")
	}

	err := query.Find(&parts).Error

	if err != nil {
		return nil, err
	}

	var partsResponse []schemas.PartResponse
	for _, part := range parts {
		partsResponse = append(partsResponse, part.ToResponse())
	}

	c.Next()
	return partsResponse, nil
}

func GetPart(id string) (schemas.PartResponse, error) {
	db := config.GetDB()
	var part schemas.Part
	db.First(&part, id)
	return part.ToResponse(), nil
}

func CreatePart(part schemas.Part) (schemas.PartResponse, error) {
	db := config.GetDB()
	err := db.Create(&part).Error
	if err != nil {
		return schemas.PartResponse{}, err
	}
	return part.ToResponse(), nil
}

func UpdatePart(part schemas.Part, id string) (schemas.PartResponse, error) {
	db := config.GetDB()
	var partUpdated schemas.Part
	err := db.First(&partUpdated, id).Error
	if err != nil {
		return schemas.PartResponse{}, err
	}
	if part.Name != "" {
		partUpdated.Name = part.Name
	}
	if part.Description != "" {
		partUpdated.Description = part.Description
	}
	if part.Quantity != 0 {
		partUpdated.Quantity = part.Quantity
	}
	if part.Price != 0 {
		partUpdated.Price = part.Price
	}
	if err := db.Save(&partUpdated).Error; err != nil {
		return schemas.PartResponse{}, err
	}
	return partUpdated.ToResponse(), nil
}

func DeletePart(id string) error {
	db := config.GetDB()
	err := db.Delete(&schemas.Part{}, id).Error
	return err
}
