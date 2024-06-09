package schemas

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Description string  `gorm:"size:255;not null"`
	Price       float64 `gorm:"not null"`
	Time        int     `gorm:"size:255;not null"`
	Orders      []Order `gorm:"many2many:order_services;"`
}

type ServiceResponse struct {
	ID          uint    `json:"id"`
	Description string  `json:"description"`
	Time        int     `json:"time"`
	Price       float64 `json:"price"`
}

func (s Service) ToResponse() ServiceResponse {
	return ServiceResponse{
		ID:          s.ID,
		Description: s.Description,
		Price:       s.Price,
	}
}
