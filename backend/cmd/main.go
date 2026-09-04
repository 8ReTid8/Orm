package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/services"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.ConnectDB()
	accountRepo := repositories.NewAccountRepository(
		database.DB,
	)
	accountService := services.NewAccountService(
		accountRepo,
	)

	accountHandler := handlers.NewAccountHandler(
		accountService,
	)
	transactionRepo := repositories.NewTransactionRepository(
		database.DB,
	)

	transactionService := services.NewTransactionService(
		database.DB,
		accountRepo,
		transactionRepo,
	)

	transactionHandler := handlers.NewTransactionHandler(
		transactionService,
	)
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set(
			"Access-Control-Allow-Origin",
			"http://localhost:5173",
			// "http://192.168.1.102:5173",

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

	routes.SetupRoutes(router, transactionHandler, accountHandler)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
