package router

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()

	config := cors.Config{
		AllowOrigins:     []string{"https://info-os.shop, http://localhost:3000, http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}

	r.Use(cors.New(config))
	initializeRoutes(r)
	r.Run(":8080")
}
