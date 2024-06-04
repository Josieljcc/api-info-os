package service

import (
	"log"

	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/schemas"
)

func GetClients() ([]schemas.Client, error) {
	db := config.GetDB()
	var clients []schemas.Client
	if err := db.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func GetClient(id string) (schemas.Client, error) {
	db := config.GetDB()
	var client schemas.Client
	if err := db.First(&client, id).Error; err != nil {
		return schemas.Client{}, err
	}
	return client, nil
}

func CreateClient(client schemas.Client) (schemas.Client, error) {
	db := config.GetDB()
	if err := db.Create(&client).Error; err != nil {
		return schemas.Client{}, err
	}
	return client, nil
}

func UpdateClient(client schemas.Client, id string) error {
	db := config.GetDB()
	var clientUpdated schemas.Client
	err := db.First(&clientUpdated, id).Error
	if err != nil {
		return err
	}
	clientUpdated.Name = client.Name
	clientUpdated.Password = client.Password
	clientUpdated.Email = client.Email
	clientUpdated.Address = client.Address
	clientUpdated.Phone = client.Phone
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
	log.Printf("Client %s deleted", id)
	return nil
}
