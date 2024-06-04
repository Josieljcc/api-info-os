package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Technician struct {
	gorm.Model
	Name     string  `gorm:"size:100;not null"`
	Password string  `gorm:"size:100;not null"`
	Email    string  `gorm:"size:100;not null;uniqueIndex"`
	Phone    string  `gorm:"size:15"`
	Orders   []Order `gorm:"foreignkey:TechnicianID"`
}

type TechnicianResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Orders    []Order   `json:"orders"`
}
