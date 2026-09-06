package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	transactionHandler *handlers.TransactionHandler,
	accountHandler *handlers.AccountHandler,
	budgetHandler *handlers.BudgetHandler,
) {
	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/categories", handlers.GetCategories)
		
		//Account
		// protected.POST("/accounts", handlers.CreateAccount)
		// protected.GET("/accounts", handlers.GetAccounts)
		protected.POST("/accounts", accountHandler.CreateAccount)
		protected.GET("/accounts", accountHandler.GetAccounts)
		// protected.PATCH("/accounts/:id", accountHandler.UpdateAccount)
		// protected.DELETE("/accounts/:id", accountHandler.DeleteAccount)
		//Transaction
		protected.POST("/transactions",transactionHandler.CreateTransaction)
		protected.GET("/transactions", transactionHandler.GetTransactions)
		protected.PATCH("/transactions/:id",transactionHandler.UpdateTransaction)
		protected.DELETE("/transactions/:id",transactionHandler.DeleteTransaction)

		//Budget
		protected.POST("/budgets", budgetHandler.CreateBudget)
		protected.GET("/budgets", budgetHandler.GetBudgets)
		protected.PATCH("/budgets/:id", handlers.UpdateBudget)
		protected.DELETE("/budgets/:id", handlers.DeleteBudget)
	}
}
