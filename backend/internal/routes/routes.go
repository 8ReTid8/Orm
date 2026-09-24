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
	authHandler *handlers.AuthHandler,
	periodHandler *handlers.PeriodHandler,
	summaryHandler *handlers.SummaryHandler,
) {
	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/summary", summaryHandler.GetSummary)
		protected.GET("/summary/comparison", summaryHandler.GetComparison)
		protected.GET("/period/years", periodHandler.GetAvailableYears)

		protected.GET("/categories", handlers.GetCategories)

		//Account
		protected.POST("/accounts", accountHandler.CreateAccount)
		protected.GET("/accounts", accountHandler.GetAccounts)
		protected.DELETE("/accounts/:id", accountHandler.DeleteAccount)
		// protected.PATCH("/accounts/:id", accountHandler.UpdateAccount)

		//Transaction
		protected.POST("/transactions", transactionHandler.CreateTransaction)
		protected.GET("/transactions", transactionHandler.GetTransactions)
		protected.GET("/accounts/:id/transactions",transactionHandler.GetAccountTransactions,)
		protected.PATCH("/transactions/:id", transactionHandler.UpdateTransaction)
		protected.DELETE("/transactions/:id", transactionHandler.DeleteTransaction)

		//Budget
		protected.POST("/budgets", budgetHandler.CreateBudget)
		protected.GET("/budgets", budgetHandler.GetBudgets)
		protected.GET("/budgets/:id", budgetHandler.GetBudgetDetail)
		protected.PATCH("/budgets/:id", budgetHandler.UpdateBudget)
		protected.DELETE("/budgets/:id", budgetHandler.DeleteBudget)
	}
}
