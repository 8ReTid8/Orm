package handlers

import (
	"fmt"
	"net/http"
	

	"backend/internal/database"
	
	"backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getUserID(c *gin.Context) (uint, bool) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return 0, false
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user ID",
		})
		return 0, false
	}

	return userID, true
}

func findUserAccount(userID uint, accountID uint) (*models.Account, error) {

	var account models.Account

	if err := database.DB.
		Where(
			"id = ? AND user_id = ?",
			accountID,
			userID,
		).
		First(&account).Error; err != nil {

		return nil, err
	}

	return &account, nil
}

func addTransactionToBalance(
	tx *gorm.DB,
	accountID uint,
	transactionType string,
	amount float64,
) error {

	var expression string

	switch transactionType {
	case "income":
		expression = "balance + ?"

	case "expense":
		expression = "balance - ?"

	default:
		return fmt.Errorf("ประเภท transaction ไม่ถูกต้อง")
	}

	return tx.Model(&models.Account{}).
		Where("id = ?", accountID).
		Update(
			"balance",
			gorm.Expr(expression, amount),
		).Error
}

