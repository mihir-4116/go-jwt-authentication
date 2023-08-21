package main

import (
	"os"

	"github.com/gin-gonic/gin"
	// routes "go-jwt-authentication/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := gin.New()

	// router.AuthRoutes(router)
	// router.UserRoutes(router)

	router.GET("/api-1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Success": "Access granted for api-1"})
	})

	router.GET("/api-2", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"Success": "Access granted for api-2"})
	})

	router.Run(":" + port)
}