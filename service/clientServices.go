package service

import (
	"log"

	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
	"github.com/Josieljcc/api-info-os/utils"
)

func GetClients() ([]schemas.ClientResponse, error) {
	db := config.GetDB()
	var clients []schemas.Client
	if err := db.Find(&clients).Error; err != nil {
		return nil, err
	}
	var clientsResponse []schemas.ClientResponse
	for _, client := range clients {
		clientsResponse = append(clientsResponse, client.ToResponse())
	}
	return clientsResponse, nil
}

func GetClientByEmail(email string) (schemas.Client, error) {
	db := config.GetDB()
	var client schemas.Client
	err := db.Where("email = ?", email).First(&client).Error
	return client, err
}

func GetClient(id string) (schemas.ClientResponse, error) {
	db := config.GetDB()
	var client schemas.Client
	if err := db.First(&client, id).Error; err != nil {
		return schemas.ClientResponse{}, err
	}
	return client.ToResponse(), nil
}

func CreateClient(client schemas.Client) (schemas.ClientResponse, error) {
	db := config.GetDB()
	hashedPassword, err := utils.HashPassword(client.Password)
	if err != nil {
		return schemas.ClientResponse{}, err
	}
	client.Password = string(hashedPassword)
	log.Printf("Hashed password: %s", string(hashedPassword))

	if err := db.Create(&client).Error; err != nil {
		return schemas.ClientResponse{}, err
	}

	return client.ToResponse(), nil
}

func UpdateClient(client schemas.Client, id string) error {
	db := config.GetDB()
	var clientUpdated schemas.Client
	err := db.First(&clientUpdated, id).Error
	if err != nil {
		return err
	}
	if client.Name != "" {
		clientUpdated.Name = client.Name
	}
	if client.Password != "" {
		hashedPassword, err := utils.HashPassword(client.Password)
		if err != nil {
			return err
		}
		clientUpdated.Password = string(hashedPassword)
	}
	if client.Email != "" {
		clientUpdated.Email = client.Email
	}
	if client.Address != "" {
		clientUpdated.Address = client.Address
	}
	if client.Phone != "" {
		clientUpdated.Phone = client.Phone
	}
	if err := db.Save(&clientUpdated).Error; err != nil {
		return err
	}
	return nil
}

func DeleteClient(id string) error {
	db := config.GetDB()
	if err := db.Delete(&schemas.Client{}, id).Error; err != nil {
		return err
	}
	return nil
}
