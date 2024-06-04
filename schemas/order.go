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
	ID         uint       `json:"id"`
	CreatedAt  string     `json:"createdAt"`
	UpdatedAt  string     `json:"updatedAt"`
	DeletedAt  string     `json:"deletedAt,omitempty"`
	Date       string     `json:"date"`
	Status     string     `json:"status"`
	Comment    string     `json:"comment"`
	Client     Client     `json:"client"`
	Technician Technician `json:"technician"`
	Services   []Service  `json:"services"`
	Parts      []Part     `json:"parts"`
}
