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
}


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
}

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

