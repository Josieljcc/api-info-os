package schemas

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Date         string    `gorm:"not null"`
	Status       string    `gorm:"not null"`
	Comment      string    `gorm:"size:255"`
	ClientID     string    `gorm:"not null"`
	TechnicianID string    `gorm:"not null"`
	Services     []Service `gorm:"many2many:order_services;"`
	Parts        []Part    `gorm:"many2many:order_parts;"`
}

type OrderResponse struct {
	ID           uint   `json:"id"`
	Date         string `json:"date"`
	Status       string `json:"status"`
	Comment      string `json:"comment"`
	ClientID     string `json:"clientId"`
	TechnicianID string `json:"technicianId"`
}

func (o Order) ToResponse() OrderResponse {
	return OrderResponse{
		ID:           o.ID,
		Date:         o.Date,
		Status:       o.Status,
		Comment:      o.Comment,
		ClientID:     o.ClientID,
		TechnicianID: o.TechnicianID,
	}
}
