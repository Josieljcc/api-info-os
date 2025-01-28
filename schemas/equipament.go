package schemas

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name            string `gorm:"size:100;not null" json:"name"`
	Description     string `gorm:"size:255" json:"description"`
	EquipamentModel string `gorm:"size:50" json:"model"`
	SerialNumber    string `gorm:"size:100; uniqueIndex" json:"serialNumber"`
	ClientID        uint   `gorm:"not null" json:"client_id"`
	Client          Client `gorm:"foreignkey:ClientID" json:"client"`
}

type EquipmentRegister struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	EquipamentModel string `json:"model"`
	SerialNumber    string `json:"serialNumber"`
}

type EquipmentResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	EquipamentModel string `json:"model"`
	SerialNumber    string `json:"serialNumber"`
}

func (e Equipment) ToResponse() EquipmentResponse {
	return EquipmentResponse{
		ID:              e.ID,
		Name:            e.Name,
		Description:     e.Description,
		EquipamentModel: e.EquipamentModel,
		SerialNumber:    e.SerialNumber,
	}
}
