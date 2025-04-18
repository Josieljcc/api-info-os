package schemas

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Name           string `gorm:"size:100;not null" json:"name"`
	Description    string `gorm:"size:255" json:"description"`
	EquipmentModel string `gorm:"size:50" json:"model"`
	SerialNumber   string `gorm:"size:100; uniqueIndex" json:"serialNumber"`
	ClientID       uint   `gorm:"not null" json:"clientID"`
	Client         Client `gorm:"foreignkey:ClientID" json:"client"`
}

type EquipmentRegister struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	EquipmentModel string `json:"model"`
	SerialNumber   string `json:"serialNumber"`
	ClientID       uint   `json:"clientID"`
}

type EquipmentResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	EquipmentModel string `json:"model"`
	SerialNumber   string `json:"serialNumber"`
	TotalPages     int    `json:"totalPages"`
	Page           int    `json:"page"`
	PageSize       int    `json:"pageSize"`
}

func (e Equipment) ToResponse() EquipmentResponse {
	return EquipmentResponse{
		ID:             e.ID,
		Name:           e.Name,
		Description:    e.Description,
		EquipmentModel: e.EquipmentModel,
		SerialNumber:   e.SerialNumber,
	}
}
