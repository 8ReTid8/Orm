package handlers

import (
	"net/http"

	"backend/internal/dto"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)

type SummaryHandler struct {
	service *services.SummaryService
}

func NewSummaryHandler(
	service *services.SummaryService,
) *SummaryHandler {
	return &SummaryHandler{
		service: service,
	}
}

func (h *SummaryHandler) GetSummary(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	filter, err := parseSummaryFilter(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	summary, err := h.service.Get(
		c.Request.Context(),
		userID,
		filter,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดข้อมูลสรุปได้",
		})
		return
	}

	response := dto.SummaryResponse{
		TotalIncome:       summary.TotalIncome,
		TotalExpense:      summary.TotalExpense,
		NetBalance:        summary.NetBalance,
		SavingsRate:       summary.SavingsRate,
		IncomeByCategory:  summary.IncomeByCategory,
		ExpenseByCategory: summary.ExpenseByCategory,
		// ComparisonPeriod:  summary.ComparisonPeriod,
		// Comparison:        summary.Comparison,
	}

	c.JSON(http.StatusOK, response)
}


func (h *SummaryHandler) GetComparison(
    c *gin.Context,
) {
    userID, ok := getUserID(c)
    if !ok {
        return
    }

    filter, err := parseComparisonFilter(c)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": err.Error(),
        })
        return
    }

    comparison, err := h.service.GetComparison(
        c.Request.Context(),
        userID,
        filter,
    )
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "ไม่สามารถโหลดข้อมูลเปรียบเทียบได้",
        })
        return
    }

    c.JSON(http.StatusOK, dto.ComparisonResponse{
        Period: comparison.Period,
        Items:  comparison.Items,
    })
}