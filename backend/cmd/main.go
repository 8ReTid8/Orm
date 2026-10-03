package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"backend/internal/database"
	"backend/internal/handlers"
	"backend/internal/repositories"
	"backend/internal/routes"
	"backend/internal/services"
)

func main() {
	// local dev only — บน production อ่าน env จาก platform แทน
	godotenv.Load()
	// err := godotenv.Load()
	// 	if err != nil {
	// 		log.Fatal("Error loading .env file")
	// 	}
	database.ConnectDB()

	//REPO
	userRepo := repositories.NewUserRepository(database.DB)
	refreshTokenRepo := repositories.NewRefreshTokenRepository(database.DB)
	accountRepo := repositories.NewAccountRepository(database.DB)
	transactionRepo := repositories.NewTransactionRepository(database.DB)
	budgetRepo := repositories.NewBudgetRepository(database.DB)

	go func() {
		for {
			time.Sleep(24 * time.Hour)
			if err := refreshTokenRepo.DeleteExpired(context.Background()); err != nil {
				log.Printf("cleanup expired refresh tokens failed: %v", err)
			} else {
				log.Println("cleanup expired refresh tokens: done")
			}
		}
	}()

	//SERVICE
	emailService := services.NewEmailService(
		services.EmailConfig{
			Host:        os.Getenv("SMTP_HOST"),
			Port:        os.Getenv("SMTP_PORT"),
			Username:    os.Getenv("SMTP_USERNAME"),
			Password:    os.Getenv("SMTP_PASSWORD"),
			From:        os.Getenv("SMTP_FROM"),
			FrontendURL: os.Getenv("FRONTEND_URL"),
		},
	)
	// authService := services.NewAuthService(
	// 	userRepo,
	// )
	authService := services.NewAuthService(
		userRepo,
		refreshTokenRepo,
		emailService,
	)
	accountService := services.NewAccountService(
		database.DB,
		accountRepo,
		transactionRepo,
	)
	transactionService := services.NewTransactionService(
		database.DB,
		accountRepo,
		transactionRepo,
	)
	budgetService := services.NewBudgetService(
		budgetRepo,
		accountRepo,
		transactionRepo,
	)
	summaryService := services.NewSummaryService(
		transactionRepo,
	)

	//HANDLER
	authHandler := handlers.NewAuthHandler(
		authService,
	)
	transactionHandler := handlers.NewTransactionHandler(
		transactionService,
	)
	accountHandler := handlers.NewAccountHandler(
		accountService,
	)
	budgetHandler := handlers.NewBudgetHandler(
		budgetService,
	)
	periodHandler := handlers.NewPeriodHandler(
		budgetRepo,
		transactionRepo,
	)
	summaryHandler := handlers.NewSummaryHandler(
		summaryService,
	)
	router := gin.Default()

	router.Use(func(c *gin.Context) {

		allowedOrigin := os.Getenv("FRONTEND_URL")
		if allowedOrigin == "" {
			allowedOrigin = "http://localhost:5173"
		}

		c.Writer.Header().Set(
			"Access-Control-Allow-Origin",
			// "http://localhost:5173",
			// "http://192.168.1.102:5173",
			allowedOrigin,
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

	routes.SetupRoutes(router, transactionHandler, accountHandler, budgetHandler, authHandler, periodHandler, summaryHandler)

	// if err := router.Run(":8080"); err != nil {
	if err := router.Run(":10000"); err != nil {
		log.Fatal(err)
	}
}
