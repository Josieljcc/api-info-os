package main

import (
	"github.com/Josieljcc/api-info-os/config"
	"github.com/Josieljcc/api-info-os/docs"
	"github.com/Josieljcc/api-info-os/router"
	"github.com/joho/godotenv"
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
	docs.SwaggerInfo.Title = "API-Info-OS"
	docs.SwaggerInfo.Description = "this api is for control services orders in a maintenance shop."
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	err := godotenv.Load()
	if err != nil {
		println("Error loading .env file")
		return
	}
	err = config.Init()
	if err != nil {
		println("Error loading config file")
		return
	}

	router.Initialize()
}
