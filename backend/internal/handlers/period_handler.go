package handlers

import (
	"net/http"
	"sort"
	"time"

	"backend/internal/repositories"
	"github.com/gin-gonic/gin"
)

type PeriodHandler struct {
	budgetRepo      *repositories.BudgetRepository
	transactionRepo *repositories.TransactionRepository
}

// 👈 รับ repositories ทั้ง 2 ตัวเข้ามา
func NewPeriodHandler(
	budgetRepo *repositories.BudgetRepository,
	transactionRepo *repositories.TransactionRepository,
) *PeriodHandler {
	return &PeriodHandler{
		budgetRepo:      budgetRepo,
		transactionRepo: transactionRepo,
	}
}

func (h *PeriodHandler) GetAvailableYears(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	resource := c.Query("resource")
	ctx := c.Request.Context()

	var dbYears []int
	var err error

	// 👈 เรียกผ่าน Repository ของแต่ละตารางโดยตรง สะอาดมาก!
	switch resource {
	case "budget":
		dbYears, err = h.budgetRepo.GetAvailableYears(ctx, userID)

	case "transaction":
		dbYears, err = h.transactionRepo.GetAvailableYears(ctx, userID)

	default:
		// ถ้าไม่ระบุ ให้เรียกทั้งสองตารางแล้วนำมารวมกัน
		budgetYears, bErr := h.budgetRepo.GetAvailableYears(ctx, userID)
		if bErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "ไม่สามารถโหลดข้อมูลปีได้"})
			return
		}

		txYears, tErr := h.transactionRepo.GetAvailableYears(ctx, userID)
		if tErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "ไม่สามารถโหลดข้อมูลปีได้"})
			return
		}

		dbYears = append(budgetYears, txYears...)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "ไม่สามารถโหลดข้อมูลปีได้"})
		return
	}

	// กรองปีซ้ำ และบวกปีปัจจุบัน + ปีหน้าเข้าไป
	currentYear := time.Now().Year()
	yearSet := map[int]bool{currentYear: true, currentYear + 1: true}
	for _, y := range dbYears {
		yearSet[y] = true
	}

	years := make([]int, 0, len(yearSet))
	for y := range yearSet {
		years = append(years, y)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(years)))

	c.JSON(http.StatusOK, gin.H{"years": years})
}