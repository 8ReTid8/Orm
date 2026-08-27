package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"backend/internal/database"
	"backend/internal/dto"
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

func parseTransactionForm(c *gin.Context) (dto.TransactionInput, error) {

	transactionType := c.PostForm("type")
	category := c.PostForm("category")
	title := c.PostForm("title")
	note := c.PostForm("note")

	if transactionType != "income" &&
		transactionType != "expense" {

		return dto.TransactionInput{},
			fmt.Errorf("ประเภท transaction ไม่ถูกต้อง")
	}

	if title == "" || category == "" {
		return dto.TransactionInput{},
			fmt.Errorf("กรุณากรอกข้อมูลให้ครบ")
	}

	amount, err := strconv.ParseFloat(
		c.PostForm("amount"),
		64,
	)

	if err != nil || amount <= 0 {
		return dto.TransactionInput{},
			fmt.Errorf("จำนวนเงินไม่ถูกต้อง")
	}

	accountID64, err := strconv.ParseUint(
		c.PostForm("accountId"),
		10,
		64,
	)

	if err != nil {
		return dto.TransactionInput{},
			fmt.Errorf("บัญชีไม่ถูกต้อง")
	}

	transactionDate, err := time.Parse(
		"2006-01-02",
		c.PostForm("transactionDate"),
	)

	if err != nil {
		return dto.TransactionInput{},
			fmt.Errorf("วันที่ไม่ถูกต้อง")
	}

	return dto.TransactionInput{
		Type:            transactionType,
		Amount:          amount,
		Category:        category,
		AccountID:       uint(accountID64),
		Title:           title,
		Note:            note,
		TransactionDate: transactionDate,
	}, nil
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

func removeTransactionFromBalance(
	tx *gorm.DB,
	accountID uint,
	transactionType string,
	amount float64,
) error {

	var expression string

	switch transactionType {
	case "income":
		// เดิม + → ย้อนกลับด้วย -
		expression = "balance - ?"

	case "expense":
		// เดิม - → ย้อนกลับด้วย +
		expression = "balance + ?"

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

func parseBudgetForm(c *gin.Context) (dto.BudgetInput, error) {
	category := c.PostForm("category")
	amount, err := strconv.ParseFloat(
		c.PostForm("amount"),
		64,
	)
	if category == "" {
		return dto.BudgetInput{},
			fmt.Errorf("กรุณากรอกข้อมูลให้ครบ")
	}

	if err != nil || amount <= 0 {
		return dto.BudgetInput{},
			fmt.Errorf("จำนวนเงินไม่ถูกต้อง")
	}

	accountID64, err := strconv.ParseUint(
		c.PostForm("accountId"),
		10,
		64,
	)

	if err != nil {
		return dto.BudgetInput{},
			fmt.Errorf("บัญชีไม่ถูกต้อง")
	}

	startDate, err := time.Parse(
		"2006-01-02",
		c.PostForm("startDate"),
	)

	if err != nil {
		return dto.BudgetInput{},
			fmt.Errorf("วันที่เริ่มต้นไม่ถูกต้อง")
	}

	endDate, err := time.Parse(
		"2006-01-02",
		c.PostForm("endDate"),
	)

	if err != nil {
		return dto.BudgetInput{},
			fmt.Errorf("วันที่สิ้นสุดไม่ถูกต้อง")
	}
	now := time.Now()

	today := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		0, 0, 0, 0,
		now.Location(),
	)

	if startDate.Before(today) {
		return dto.BudgetInput{},
			fmt.Errorf("วันที่เริ่มต้นต้องไม่ก่อนวันที่ปัจจุบัน")
	}

	// วันที่สิ้นสุดต้องไม่ก่อนวันที่เริ่มต้น
	if endDate.Before(startDate) {
		return dto.BudgetInput{},
			fmt.Errorf("วันที่สิ้นสุดต้องไม่ก่อนวันที่เริ่มต้น")
	}

	return dto.BudgetInput{
		Amount:    amount,
		Category:  category,
		AccountID: uint(accountID64),
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
}
