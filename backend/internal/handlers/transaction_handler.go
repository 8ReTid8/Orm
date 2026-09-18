package handlers

import (
	"errors"

	"net/http"
	"strconv"

	"backend/internal/dto"
	"backend/internal/services"
	"backend/internal/utils"
	"github.com/gin-gonic/gin"
)

type TransactionHandler struct {
	service *services.TransactionService
}

func NewTransactionHandler(
	service *services.TransactionService,
) *TransactionHandler {
	return &TransactionHandler{
		service: service,
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {
	// 1. User จาก JWT
	// userIDValue, exists := c.Get("userID")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "Unauthorized",
	// 	})
	// 	return
	// }

	// userID, ok := userIDValue.(uint)
	// if !ok {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Invalid user ID",
	// 	})
	// 	return
	// }
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	input, err := parseTransactionForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// 4. Upload slip (optional)
	slipPath, err := utils.SaveSlip(
		c,
		userID,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถบันทึกรูปสลิปได้",
		})
		return
	}
	// service := services.NewTransactionService(database.DB)

	transaction, err := h.service.Create(
		c.Request.Context(),
		userID,
		input,
		slipPath,
	)
	if err != nil {
		switch {
		case errors.Is(err, services.ErrAccountNotFound):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "ไม่พบบัญชี",
			})

		case errors.Is(err, services.ErrInsufficientBalance):
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "ยอดเงินในบัญชีไม่เพียงพอ",
			})

		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ไม่สามารถบันทึกรายการได้",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "เพิ่มรายการสำเร็จ",
		"transaction": dto.ToTransactionResponse(*transaction),
	})
}


func (h *TransactionHandler) GetTransactions(c *gin.Context) {

	userID, ok := getUserID(c)
	if !ok {
		return
	}

	filter, err := parseTransactionFilter(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}


	transactions, err := h.service.Get(
		c.Request.Context(),
		userID,
		filter,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดรายการได้",
		})
		return
	}
	responses := make(
		[]dto.TransactionResponse,
		0,
		len(transactions),
	)

	for _, transaction := range transactions {
		responses = append(
			responses,
			dto.ToTransactionResponse(transaction),
		)
	}
	c.JSON(http.StatusOK, gin.H{
		"transactions": responses,
	})
}


func (h *TransactionHandler) UpdateTransaction(c *gin.Context) {
	// 1. user จาก JWT
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	// 2. transaction id จาก URL
	transactionID64, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || transactionID64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "transaction id ไม่ถูกต้อง",
		})
		return
	}

	// 3. รับข้อมูลใหม่
	input, err := parseTransactionForm(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	slipPath, err := utils.SaveSlip(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถบันทึกรูปสลิปได้",
		})
		return
	}
	// service := services.NewTransactionService(database.DB)

	transaction, err := h.service.Update(
		c.Request.Context(),
		userID,
		uint(transactionID64),
		input,
		slipPath,
	)

	switch {
	case errors.Is(err, services.ErrTransactionNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบ transaction",
		})
		return

	case errors.Is(err, services.ErrAccountNotFound):
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ไม่พบบัญชี",
		})
		return

	case errors.Is(err, services.ErrInsufficientBalance):
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ยอดเงินในบัญชีไม่เพียงพอ",
		})
		return

	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถแก้ไขรายการได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "แก้ไขรายการสำเร็จ",
		"transaction": dto.ToTransactionResponse(*transaction),
	})
}

func (h *TransactionHandler) DeleteTransaction(c *gin.Context) {
	// 1. user จาก JWT
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	// 2. transaction id จาก URL
	transactionID64, err := strconv.ParseUint(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || transactionID64 == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "transaction id ไม่ถูกต้อง",
		})
		return
	}
	// service := services.NewTransactionService(database.DB)

	err = h.service.Delete(
		c.Request.Context(),
		userID,
		uint(transactionID64),
	)

	switch {
	case errors.Is(err, services.ErrTransactionNotFound):
		c.JSON(http.StatusNotFound, gin.H{
			"message": "ไม่พบ transaction",
		})
		return

	case errors.Is(err, services.ErrAccountNotFound):
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ไม่พบบัญชี",
		})
		return

	case errors.Is(err, services.ErrInsufficientBalance):
		// กรณีข้อมูลเดิมผิด เช่น income มากกว่า balance ปัจจุบัน
		c.JSON(http.StatusConflict, gin.H{
			"message": "ยอดเงินในบัญชีไม่ถูกต้อง",
		})
		return

	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถลบรายการได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ลบรายการสำเร็จ",
	})

}

