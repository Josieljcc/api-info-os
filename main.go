package main

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logger.Errorf("godotenv.Load() error: %v", err)
		return
	}
	logger = config.GetLogger("main")
	err = config.Init()
	if err != nil {
		logger.Errorf("config.Init() error: %v", err)
		return
	}

	router.Initialize()
}
