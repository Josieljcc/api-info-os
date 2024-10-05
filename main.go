package main

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/docs"
	"github.com/Josieljcc/api-info-os/router"
	"github.com/joho/godotenv"
)

var (
	logger *config.Logger
)

// @title API de Controle de Ordens de Serviço
// @version 1.0
// @description Esta é a API para controle de ordens de serviço em uma assistência de manutenção de computadores.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	docs.SwaggerInfo.Title = "API de Controle de Ordens de Serviço"
	docs.SwaggerInfo.Description = "Esta é a API para controle de ordens de serviço em uma assistência de manutenção de computadores."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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
