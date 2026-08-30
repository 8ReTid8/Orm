package handlers

import (
	"backend/internal/database"
	"backend/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBudget(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}
	
	input, err := parseBudgetForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	_, err = findUserAccount(
		userID,
		input.AccountID,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ไม่พบบัญชี",
		})
		return
	}
	budget := models.Budget{
		UserID:    userID,
		AccountID: input.AccountID,
		Amount:    input.Amount,
		Category:  input.Category,
		StartDate: input.StartDate,
		EndDate:   input.EndDate,
	}
	if err := database.DB.Create(&budget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถสร้างงบประมาณได้",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "งบประมาณถูกสร้างเรียบร้อยแล้ว",
		"budget":  budget,
	})
}

func GetBudgets(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}
	// 1. รับ Query Parameters (year, month, status, accountId)
	yearString := c.Query("year")
	monthString := c.Query("month")
	status := c.Query("status")       // "active" หรือ "ended"
	// accountIDStr := c.Query("accountId")

	now := time.Now()
	year := now.Year()
	month := int(now.Month())
	if y, err := strconv.Atoi(yearString); err == nil && y > 0 {
		year = y
	}
	if m, err := strconv.Atoi(monthString); err == nil && m >= 1 && m <= 12 {
		month = m
	}
	// ช่วงเวลาของเดือนที่เลือก (วันแรก ถึง วันสิ้นสุดเดือน)
	monthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Nanosecond)
	// 2. Query ดึง Budget + คำนวณ Spent จาก Transactions อัตโนมัติในคำสั่งเดียว
	query := database.DB.Table("budgets").
		Select(`
			budgets.*,
			COALESCE((
				SELECT SUM(transactions.amount)
				FROM transactions
				WHERE transactions.account_id = budgets.account_id
				  AND transactions.category = budgets.category
				  AND transactions.type = 'expense'
				  AND transactions.transaction_date >= budgets.start_date
				  AND transactions.transaction_date <= budgets.end_date
			), 0) AS spent
		`).
		Where("budgets.user_id = ?", userID).
		// เงื่อนไข Date Overlap: อยู่ในช่วงเดือนที่เลือก
		Where("budgets.start_date <= ? AND budgets.end_date >= ?", monthEnd, monthStart)
	// 3. กรองตามสถานะ Active / Ended
	if status == "active" {
		query = query.Where("budgets.end_date >= ?", now)
	} else if status == "ended" {
		query = query.Where("budgets.end_date < ?", now)
	}
	// 4. กรองตาม Account ID (ถ้ามีส่งมา)
	// if accountID, err := strconv.ParseUint(accountIDStr, 10, 64); err == nil && accountID > 0 {
	// 	query = query.Where("budgets.account_id = ?", uint(accountID))
	// }
	// var budgets []BudgetResponse
	var budgets []models.Budget
	// Preload ข้อมูล Account และเรียงตามวันเริ่มต้น
	if err := query.
		Preload("Account").
		Order("budgets.start_date ASC").
		Find(&budgets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดข้อมูลงบประมาณได้",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"budgets": budgets,
	})
}