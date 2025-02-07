package schemas

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	ID       uint    `gorm:"primaryKey" json:"id" auto:"increment"`
	Name     string  `gorm:"size:100;not null" json:"name"`
	Password string  `gorm:"size:100;not null; hash:bcrypt" json:"password"`
	Email    string  `gorm:"size:100;not null;uniqueIndex" json:"email"`
	Address  string  `gorm:"size:255" json:"address"`
	Phone    string  `gorm:"size:15" json:"phone"`
	Orders   []Order `gorm:"foreignkey:ClientID" json:"orders"`
}

type ClientResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type ClientRegister struct {
	Name     string `json:"name" example:"Josiel"`
	Email    string `json:"email" example:"josiel.jcc@hotmail.com"`
	Phone    string `json:"phone" example:"(11) 99999-9999"`
	Password string `json:"password" example:"123456"`
}

type ClientLoginResponse struct {
	Token string `json:"token"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ClientLogin struct {
	Email    string `json:"email" example:"josiel.jcc@hotmail.com"`
	Password string `json:"password" example:"123456"`
}

func (c Client) ToResponse() ClientResponse {
	return ClientResponse{
		ID:      c.ID,
		Name:    c.Name,
		Email:   c.Email,
		Address: c.Address,
		Phone:   c.Phone,
	}
}
