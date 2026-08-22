package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// protected.GET("/banks", handlers.GetBanks)
		protected.GET("/categories", handlers.GetCategories)
		protected.POST("/accounts", handlers.CreateAccount)
		protected.GET("/accounts", handlers.GetAccounts)
		protected.POST(
			"/transactions",
			handlers.CreateTransaction,
		)
		protected.GET("/transactions", handlers.GetTransactions)
		protected.PATCH(
			"/transactions/:id",
			handlers.UpdateTransaction,
		)
	}
}
