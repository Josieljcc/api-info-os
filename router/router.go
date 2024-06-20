package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	r := gin.Default()
	r.Use(cors.Default())
	initializeRoutes(r)
	r.Run(":8080")
}
