package handlers

import (
	"backend/internal/dto"
	"backend/internal/services"
	"errors"
	"net/http"
	"strconv"
	// "time"

	"github.com/gin-gonic/gin"
)

type BudgetHandler struct {
	service *services.BudgetService
}

func NewBudgetHandler(
	service *services.BudgetService,
) *BudgetHandler {
	return &BudgetHandler{
		service: service,
	}
}

// func CreateBudget(c *gin.Context) {
func (h *BudgetHandler) CreateBudget(c *gin.Context) {
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
	budget, err := h.service.Create(
		c.Request.Context(),
		userID,
		input,
	)

	switch {
	// case errors.Is(err, services.ErrInvalidBudgetInput):
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "ข้อมูลงบประมาณไม่ถูกต้อง",
	// 	})
	// 	return

	case errors.Is(err, services.ErrAccountNotFound):
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ไม่พบบัญชี",
		})
		return

	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถสร้างงบประมาณได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "งบประมาณถูกสร้างเรียบร้อยแล้ว",
		"budget":  dto.ToBudgetResponse(*budget),
	})
	// _, err = findUserAccount(
	// 	userID,
	// 	input.AccountID,
	// )
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "ไม่พบบัญชี",
	// 	})
	// 	return
	// }
	// budget := models.Budget{
	// 	UserID:    userID,
	// 	AccountID: input.AccountID,
	// 	Amount:    input.Amount,
	// 	Category:  input.Category,
	// 	StartDate: input.StartDate,
	// 	EndDate:   input.EndDate,
	// }
	// if err := database.DB.Create(&budget).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถสร้างงบประมาณได้",
	// 	})
	// 	return
	// }
	// c.JSON(http.StatusCreated, gin.H{
	// 	"message": "งบประมาณถูกสร้างเรียบร้อยแล้ว",
	// 	"budget":  budget,
	// })
}

// func GetBudgets(c *gin.Context) {
func (h *BudgetHandler) GetBudgets(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}
	filter, err := parseBudgetFilter(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// budgets, err := h.service.Get(
	// 	c.Request.Context(),
	// 	userID,
	// 	filter,
	// )
	summaries, err := h.service.Get(
		c.Request.Context(),
		userID,
		filter,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดข้อมูลงบประมาณได้",
		})
		return
	}

	responses := make(
		[]dto.BudgetOverviewResponse,
		0,
		len(summaries),
	)

	for _, summary := range summaries {
		responses = append(
			responses,
			dto.ToBudgetOverviewResponse(
				summary.Budget,
				summary.Spent,
				summary.IsActive,
			),
		)
	}

	c.JSON(http.StatusOK, gin.H{
		"budgets": responses,
	})
	// ---------------------------
	// 1. Query parameters
	// ---------------------------

	// now := time.Now()
	// yearStr := c.Query("year")
	// monthStr := c.Query("month")
	// var selectedYear *int
	// var selectedMonth *int

	// if yearStr != "" {
	// 	year, err := strconv.Atoi(yearStr)

	// 	if err != nil || year <= 0 {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "ปีไม่ถูกต้อง",
	// 		})
	// 		return
	// 	}

	// 	selectedYear = &year
	// }

	// if monthStr != "" {
	// 	if selectedYear == nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "ต้องเลือกปีก่อนเลือกเดือน",
	// 		})
	// 		return
	// 	}

	// 	month, err := strconv.Atoi(monthStr)

	// 	if err != nil || month < 1 || month > 12 {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "เดือนไม่ถูกต้อง",
	// 		})
	// 		return
	// 	}

	// 	selectedMonth = &month
	// }

	// status := c.Query("status")

	// if status != "" && status != "active" && status != "ended" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "status ไม่ถูกต้อง",
	// 	})
	// 	return
	// }

	// // ---------------------------
	// // 2. Calculate selected month
	// // ---------------------------

	// query := database.DB.
	// 	Table("budgets").
	// 	Where("budgets.user_id = ?", userID)
	// 	// Where(
	// 	// 	"budgets.start_date <= ? AND budgets.end_date >= ?",
	// 	// 	monthEnd,
	// 	// 	monthStart,
	// 	// )

	// switch status {
	// case "active":
	// 	query = query.Where(
	// 		"budgets.start_date <= ? AND budgets.end_date >= ?",
	// 		now,
	// 		now,
	// 	)

	// case "ended":
	// 	query = query.Where(
	// 		"budgets.end_date < ?",
	// 		now,
	// 	)
	// 	if selectedYear != nil {

	// 		var rangeStart time.Time
	// 		var rangeEnd time.Time

	// 		if selectedMonth != nil {
	// 			// เลือก ปี + เดือน
	// 			rangeStart = time.Date(
	// 				*selectedYear,
	// 				time.Month(*selectedMonth),
	// 				1,
	// 				0, 0, 0, 0,
	// 				time.Local,
	// 			)

	// 			rangeEnd = rangeStart.AddDate(0, 1, 0)

	// 		} else {
	// 			// เลือกเฉพาะปี
	// 			rangeStart = time.Date(
	// 				*selectedYear,
	// 				1,
	// 				1,
	// 				0, 0, 0, 0,
	// 				time.Local,
	// 			)

	// 			rangeEnd = rangeStart.AddDate(1, 0, 0)
	// 		}

	// 		// Budget ต้อง overlap กับช่วงที่เลือก
	// 		query = query.Where(
	// 			"budgets.start_date < ? AND budgets.end_date >= ?",
	// 			rangeEnd,
	// 			rangeStart,
	// 		)
	// 	}
	// }
	// // ---------------------------
	// // 3. Account filter
	// // ---------------------------

	// accountIDStr := c.Query("accountId")

	// if accountIDStr != "" {
	// 	accountID, err := strconv.ParseUint(
	// 		accountIDStr,
	// 		10,
	// 		64,
	// 	)

	// 	if err != nil || accountID == 0 {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"message": "account id ไม่ถูกต้อง",
	// 		})
	// 		return
	// 	}

	// 	query = query.Where(
	// 		"budgets.account_id = ?",
	// 		uint(accountID),
	// 	)
	// }

	// // ---------------------------
	// // 4. Load budgets
	// // ---------------------------

	// var budgets []models.Budget

	// if err := query.
	// 	// Preload("Account").
	// 	Order("budgets.start_date ASC").
	// 	Find(&budgets).Error; err != nil {

	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถโหลดข้อมูลงบประมาณได้",
	// 	})
	// 	return
	// }

	// // ---------------------------
	// // 5. Build response
	// // ---------------------------

	// responses := make([]dto.BudgetResponse, 0, len(budgets))

	// for _, budget := range budgets {

	// 	var spent float64

	// 	err := database.DB.
	// 		Table("transactions").
	// 		Select("COALESCE(SUM(amount), 0)").
	// 		Where(
	// 			"account_id = ?",
	// 			budget.AccountID,
	// 		).
	// 		Where(
	// 			"category = ?",
	// 			budget.Category,
	// 		).
	// 		Where(
	// 			"type = ?",
	// 			"expense",
	// 		).
	// 		Where(
	// 			"transaction_date >= ?",
	// 			budget.StartDate,
	// 		).
	// 		Where(
	// 			"transaction_date <= ?",
	// 			budget.EndDate,
	// 		).
	// 		Scan(&spent).Error

	// 	if err != nil {
	// 		c.JSON(http.StatusInternalServerError, gin.H{
	// 			"message": "ไม่สามารถคำนวณยอดใช้จ่ายได้",
	// 		})
	// 		return
	// 	}

	// 	isActive :=
	// 		!now.Before(budget.StartDate) &&
	// 			!now.After(budget.EndDate)

	// 	responses = append(
	// 		responses,
	// 		dto.BudgetResponse{
	// 			ID:        budget.ID,
	// 			AccountID: budget.AccountID,
	// 			Category:  budget.Category,
	// 			Amount:    budget.Amount,
	// 			StartDate: budget.StartDate,
	// 			EndDate:   budget.EndDate,

	// 			Spent:     spent,
	// 			Remaining: budget.Amount - spent,
	// 			IsActive:  isActive,

	// 			Account: budget.Account,
	// 		},
	// 	)
	// }

	// // ---------------------------
	// // 6. Response
	// // ---------------------------

	// c.JSON(http.StatusOK, gin.H{
	// 	"budgets": responses,
	// })
}
func (h *BudgetHandler) GetBudgetDetail(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

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

	detail, err := h.service.GetDetail(
		c.Request.Context(),
		userID,
		uint(budgetID64),
	)

	switch {
	case errors.Is(err, services.ErrBudgetNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบงบประมาณ",
		})
		return

	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดรายละเอียดงบประมาณได้",
		})
		return
	}

	transactions := make(
		[]dto.TransactionResponse,
		0,
		len(detail.Transactions),
	)

	for _, transaction := range detail.Transactions {
		transactions = append(
			transactions,
			dto.ToTransactionResponse(transaction),
		)
	}

	response := dto.BudgetDetailResponse{
		Budget: dto.ToBudgetOverviewResponse(
			detail.Budget,
			detail.Spent,
			detail.IsActive,
		),
		Transactions: transactions,
	}

	c.JSON(http.StatusOK, response)
}

// func UpdateBudget(c *gin.Context) {
func (h *BudgetHandler) UpdateBudget(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

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
	// budgetID := uint(budgetID64)

	input, err := parseBudgetForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	budget, err := h.service.Update(
		c.Request.Context(),
		userID,
		uint(budgetID64),
		input,
	)

	switch {
	case errors.Is(err, services.ErrBudgetNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบงบประมาณ",
		})
		return

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "แก้ไขงบประมาณสำเร็จ",
		"budget":  dto.ToBudgetResponse(*budget),
	})

	// var oldBudget models.Budget

	// if err := database.DB.
	// 	// Joins(
	// 	// 	"JOIN accounts ON accounts.id = budgets.account_id",
	// 	// ).
	// 	// Where(
	// 	// 	"budgets.id = ? AND accounts.user_id = ?",
	// 	// 	budgetID,
	// 	// 	userID,
	// 	// ).
	// 	Where("id = ? AND user_id = ?", budgetID, userID).
	// 	First(&oldBudget).Error; err != nil {

	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"message": "ไม่พบ งบประมาณ",
	// 	})
	// 	return
	// }
	// _, err = findUserAccount(
	// 	userID,
	// 	input.AccountID,
	// )
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "ไม่พบบัญชี",
	// 	})
	// 	return
	// }
	// // 5. อัปเดตค่าฟิลด์ต่างๆ
	// oldBudget.AccountID = input.AccountID
	// oldBudget.Amount = input.Amount
	// oldBudget.Category = input.Category
	// oldBudget.StartDate = input.StartDate
	// oldBudget.EndDate = input.EndDate
	// // 6. บันทึกลง Database
	// if err := database.DB.Save(&oldBudget).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถแก้ไขงบประมาณได้",
	// 	})
	// 	return
	// }
	// // 7. ส่ง Response สำเร็จ (HTTP 200 OK)
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "แก้ไขงบประมาณสำเร็จ",
	// 	"budget":  oldBudget,
	// })
}

// func DeleteBudget(c *gin.Context) {
func (h *BudgetHandler) DeleteBudget(c *gin.Context) {
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
	err = h.service.Delete(
		c.Request.Context(),
		userID,
		uint(budgetID64),
	)

	switch {
	case errors.Is(err, services.ErrBudgetNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบงบประมาณ",
		})
		return

	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถลบงบประมาณได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ลบงบประมาณสำเร็จ",
	})
}

