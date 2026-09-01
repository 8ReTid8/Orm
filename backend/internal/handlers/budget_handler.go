package handlers

import (
	"backend/internal/database"
	"backend/internal/dto"
	"backend/internal/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

// func GetBudgets(c *gin.Context) {
// 	userID, ok := getUserID(c)
// 	if !ok {
// 		return
// 	}
// 	// 1. รับ Query Parameters (year, month, status, accountId)
// 	yearString := c.Query("year")
// 	monthString := c.Query("month")
// 	status := c.Query("status")       // "active" หรือ "ended"
// 	// accountIDStr := c.Query("accountId")

// 	now := time.Now()
// 	year := now.Year()
// 	month := int(now.Month())
// 	if y, err := strconv.Atoi(yearString); err == nil && y > 0 {
// 		year = y
// 	}
// 	if m, err := strconv.Atoi(monthString); err == nil && m >= 1 && m <= 12 {
// 		month = m
// 	}
// 	// ช่วงเวลาของเดือนที่เลือก (วันแรก ถึง วันสิ้นสุดเดือน)
// 	monthStart := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
// 	monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Nanosecond)
// 	// 2. Query ดึง Budget + คำนวณ Spent จาก Transactions อัตโนมัติในคำสั่งเดียว
// 	query := database.DB.Table("budgets").
// 		Select(`
// 			budgets.*,
// 			COALESCE((
// 				SELECT SUM(transactions.amount)
// 				FROM transactions
// 				WHERE transactions.account_id = budgets.account_id
// 				  AND transactions.category = budgets.category
// 				  AND transactions.type = 'expense'
// 				  AND transactions.transaction_date >= budgets.start_date
// 				  AND transactions.transaction_date <= budgets.end_date
// 			), 0) AS spent
// 		`).
// 		Where("budgets.user_id = ?", userID).
// 		// เงื่อนไข Date Overlap: อยู่ในช่วงเดือนที่เลือก
// 		Where("budgets.start_date <= ? AND budgets.end_date >= ?", monthEnd, monthStart)
// 	// 3. กรองตามสถานะ Active / Ended
// 	if status == "active" {
// 		query = query.Where("budgets.end_date >= ?", now)
// 	} else if status == "ended" {
// 		query = query.Where("budgets.end_date < ?", now)
// 	}
// 	// 4. กรองตาม Account ID (ถ้ามีส่งมา)
// 	// if accountID, err := strconv.ParseUint(accountIDStr, 10, 64); err == nil && accountID > 0 {
// 	// 	query = query.Where("budgets.account_id = ?", uint(accountID))
// 	// }
// 	// var budgets []BudgetResponse
// 	var budgets []models.Budget
// 	// Preload ข้อมูล Account และเรียงตามวันเริ่มต้น
// 	if err := query.
// 		Preload("Account").
// 		Order("budgets.start_date ASC").
// 		Find(&budgets).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "ไม่สามารถโหลดข้อมูลงบประมาณได้",
// 			"error":   err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"budgets": budgets,
// 	})
// }

func GetBudgets(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	// ---------------------------
	// 1. Query parameters
	// ---------------------------

	now := time.Now()

	year := now.Year()
	month := int(now.Month())
	if value := c.Query("year"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed > 0 {
			year = parsed
		}
	}

	if value := c.Query("month"); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil && parsed >= 1 && parsed <= 12 {
			month = parsed
		}
	}
	status := c.Query("status")

	if status != "" && status != "active" && status != "ended" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "status ไม่ถูกต้อง",
		})
		return
	}

	// ---------------------------
	// 2. Calculate selected month
	// ---------------------------

	monthStart := time.Date(
		year,
		time.Month(month),
		1,
		0, 0, 0, 0,
		time.Local,
	)

	monthEnd := monthStart.
		AddDate(0, 1, 0).
		Add(-time.Nanosecond)

	query := database.DB.
		Table("budgets").
		Where("budgets.user_id = ?", userID).
		Where(
			"budgets.start_date <= ? AND budgets.end_date >= ?",
			monthEnd,
			monthStart,
		)
		
	switch status {
	case "active":
		query = query.Where(
			"budgets.start_date <= ? AND budgets.end_date >= ?",
			now,
			now,
		)

	case "ended":
		query = query.Where(
			"budgets.end_date < ?",
			now,
		)
	}
	// ---------------------------
	// 3. Account filter
	// ---------------------------

	accountIDStr := c.Query("accountId")

	if accountIDStr != "" {
		accountID, err := strconv.ParseUint(
			accountIDStr,
			10,
			64,
		)

		if err != nil || accountID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "account id ไม่ถูกต้อง",
			})
			return
		}

		query = query.Where(
			"budgets.account_id = ?",
			uint(accountID),
		)
	}

	// ---------------------------
	// 4. Load budgets
	// ---------------------------

	var budgets []models.Budget

	if err := query.
		Preload("Account").
		Order("budgets.start_date ASC").
		Find(&budgets).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดข้อมูลงบประมาณได้",
		})
		return
	}

	// ---------------------------
	// 5. Build response
	// ---------------------------

	responses := make([]dto.BudgetResponse, 0, len(budgets))

	for _, budget := range budgets {

		var spent float64

		err := database.DB.
			Table("transactions").
			Select("COALESCE(SUM(amount), 0)").
			Where(
				"account_id = ?",
				budget.AccountID,
			).
			Where(
				"category = ?",
				budget.Category,
			).
			Where(
				"type = ?",
				"expense",
			).
			Where(
				"transaction_date >= ?",
				budget.StartDate,
			).
			Where(
				"transaction_date <= ?",
				budget.EndDate,
			).
			Scan(&spent).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ไม่สามารถคำนวณยอดใช้จ่ายได้",
			})
			return
		}

		isActive :=
			!now.Before(budget.StartDate) &&
				!now.After(budget.EndDate)

		responses = append(
			responses,
			dto.BudgetResponse{
				ID:        budget.ID,
				AccountID: budget.AccountID,
				Category:  budget.Category,
				Amount:    budget.Amount,
				StartDate: budget.StartDate,
				EndDate:   budget.EndDate,

				Spent:     spent,
				Remaining: budget.Amount - spent,
				IsActive:  isActive,

				Account: budget.Account,
			},
		)
	}

	// ---------------------------
	// 6. Response
	// ---------------------------

	c.JSON(http.StatusOK, gin.H{
		"budgets": responses,
	})
}

func UpdateBudget(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	budgetID64, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "budget id ไม่ถูกต้อง",
		})
		return
	}
	budgetID := uint(budgetID64)

	input, err := parseBudgetForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	var oldBudget models.Budget

	if err := database.DB.
		// Joins(
		// 	"JOIN accounts ON accounts.id = budgets.account_id",
		// ).
		// Where(
		// 	"budgets.id = ? AND accounts.user_id = ?",
		// 	budgetID,
		// 	userID,
		// ).
		Where("id = ? AND user_id = ?", budgetID, userID).
		First(&oldBudget).Error; err != nil {

		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบ งบประมาณ",
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
	// 5. อัปเดตค่าฟิลด์ต่างๆ
	oldBudget.AccountID = input.AccountID
	oldBudget.Amount = input.Amount
	oldBudget.Category = input.Category
	oldBudget.StartDate = input.StartDate
	oldBudget.EndDate = input.EndDate
	// 6. บันทึกลง Database
	if err := database.DB.Save(&oldBudget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถแก้ไขงบประมาณได้",
		})
		return
	}
	// 7. ส่ง Response สำเร็จ (HTTP 200 OK)
	c.JSON(http.StatusOK, gin.H{
		"message": "แก้ไขงบประมาณสำเร็จ",
		"budget":  oldBudget,
	})
}

func DeleteBudget(c *gin.Context) {
	// 1. user จาก JWT
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	// 2. budget id จาก URL
	budgetID64, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || budgetID64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "budget id ไม่ถูกต้อง",
		})
		return
	}

	budgetID := uint(budgetID64)

	// 3. หา budget
	// และต้องเป็น budget ของ user คนนี้
	var budget models.Budget

	if err := database.DB.
		Where("id = ? AND user_id = ?", budgetID, userID).
		First(&budget).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "ไม่พบ budget",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถค้นหางบประมาณได้",
		})
		return
	}

	// 4. Database transaction
	if err := database.DB.Delete(&budget).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถลบงบประมาณได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ลบงบประมาณสำเร็จ",
	})
}
