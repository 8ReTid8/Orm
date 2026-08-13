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
	// transaction := api.Group("/transactions")
	// transaction.Use(middleware.AuthMiddleware())
	// {
	// 	transaction.POST("", handlers.CreateTransaction)
	// }
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/banks", handlers.GetBanks)
		protected.GET("/categories", handlers.GetCategories)

		// protected.POST(
		// 	"/transactions",
		// 	handlers.CreateTransaction,
		// )
	}
}
