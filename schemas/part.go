package schemas

import (
	"gorm.io/gorm"
)

type Part struct {
	gorm.Model
	Name        string  `gorm:"size:100;not null"`
	Description string  `gorm:"size:255"`
	Quantity    int     `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Orders      []Order `gorm:"many2many:order_parts;"`
}

type PartResponse struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Orders      []Order `json:"orders"`
	TotalPages  int     `json:"totalPages"`
	Page        int     `json:"page"`
	PageSize    int     `json:"pageSize"`
}

type PartCreate struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
}

func (p Part) ToResponse() PartResponse {
	return PartResponse{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		Quantity:    p.Quantity,
		Price:       p.Price,
		Orders:      p.Orders,
	}
}
