package schemas

import (
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name     string
	Password string
	Email    string
	Address  string
	Phone    string
}
