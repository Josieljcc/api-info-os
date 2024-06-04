package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	db, err = initMySQL()
	if err != nil {
		return fmt.Errorf("failed to init database: %w", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	if logger == nil {
		logger = NewLogger(p)
	}
	return logger
}
