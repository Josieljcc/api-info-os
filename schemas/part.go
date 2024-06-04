package schemas

import (
	"time"

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
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	Orders      []Order   `json:"orders"`
}
