package schemas

import (
	"gorm.io/gorm"
)

type Technician struct {
	gorm.Model
	Name     string  `gorm:"size:100;not null"`
	Password string  `gorm:"size:100;not null; min: 6"`
	Email    string  `gorm:"size:100;not null;uniqueIndex"`
	Phone    string  `gorm:"size:15"`
	Orders   []Order `gorm:"foreignkey:TechnicianID"`
}

type TechnicianResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Phone  string  `json:"phone"`
	Orders []Order `json:"orders"`
}

func (t Technician) ToResponse() TechnicianResponse {
	return TechnicianResponse{
		ID:     t.ID,
		Name:   t.Name,
		Email:  t.Email,
		Phone:  t.Phone,
		Orders: t.Orders,
	}
}
