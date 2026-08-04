package main

import (
	"log"
	"net/http"
	
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"

	"backend/internal/database"
	"backend/internal/routes"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set(
			"Access-Control-Allow-Origin",
			"http://localhost:5173",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)
		c.Writer.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS",
		)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	})

	routes.SetupRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
