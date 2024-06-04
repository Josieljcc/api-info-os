package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string  `gorm:"size:100;not null" json:"name"`
	Password string  `gorm:"size:100;not null" json:"password"`
	Email    string  `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Address  string  `gorm:"size:255" json:"address"`
	Phone    string  `gorm:"size:15" json:"phone"`
	Orders   []Order `gorm:"foreignkey:ClientID" json:"orders"`
}

type ClientResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Name      string    `json:"name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
}
