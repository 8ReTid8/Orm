package handlers

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )
import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"backend/internal/dto"

	"github.com/gin-gonic/gin"
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

func parseTransactionForm(c *gin.Context,) (dto.TransactionInput, error) {

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