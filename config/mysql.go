package config

import (
	"os"

	"github.com/Josieljcc/api-info-os/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initMySQL() (*gorm.DB, error) {
	logger := GetLogger("[mysql] ")
	dsn := os.Getenv("URL_DATABASE")
	if dsn == "" {
		logger.Errorf("URL_DATABASE is not set")
		return nil, new(os.PathError)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect to database: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(
		&schemas.Client{},
		&schemas.Technician{},
		&schemas.Service{},
		&schemas.Part{},
		&schemas.Order{},
	)
	if err != nil {
		logger.Errorf("failed to migrate database: %v", err)
		return nil, err
	}
	return db, nil
}
