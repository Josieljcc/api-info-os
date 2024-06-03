package main

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	err := config.Init()
	if err != nil {
		logger.Errorf("config.Init() error: %v", err)
		return
	}

	router.Initialize()
}
