package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Description string  `gorm:"size:255;not null"`
	Price       float64 `gorm:"not null"`
	Orders      []Order `gorm:"many2many:order_services;"`
}

type ServiceResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Orders      []Order   `json:"orders"`
}
