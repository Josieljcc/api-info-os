package schemas

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"size:255;not null"`
	Price       float64 `gorm:"not null"`
	Time        int     `gorm:"size:255;not null"`
	Orders      []Order `gorm:"many2many:order_services;"`
}

type ServiceResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Time        int     `json:"time"`
	Price       float64 `json:"price"`
	Orders      []Order `json:"orders"`
	TotalPages  int     `json:"totalPages"`
	Page        int     `json:"page"`
	PageSize    int     `json:"pageSize"`
}

func (s Service) ToResponse() ServiceResponse {
	return ServiceResponse{
		ID:          s.ID,
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
	}
}
