package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()
	initializeRoutes(r)
	r.Run(":8080")
}
