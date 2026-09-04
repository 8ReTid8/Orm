// internal/handlers/transaction_request.go
package handlers

import (
	"fmt"
	"strconv"
	"time"

	"backend/internal/dto"

	"github.com/gin-gonic/gin"
)

func parseTransactionForm(c *gin.Context) (dto.TransactionInput, error) {
	transactionType := c.PostForm("type")
	category := c.PostForm("category")
	title := c.PostForm("title")
	note := c.PostForm("note")

	if transactionType != "income" && transactionType != "expense" {
		return dto.TransactionInput{}, fmt.Errorf("ประเภท transaction ไม่ถูกต้อง")
	}

	if title == "" || category == "" {
		return dto.TransactionInput{}, fmt.Errorf("กรุณากรอกข้อมูลให้ครบ")
	}

	amount, err := strconv.ParseFloat(c.PostForm("amount"), 64)
	if err != nil || amount <= 0 {
		return dto.TransactionInput{}, fmt.Errorf("จำนวนเงินไม่ถูกต้อง")
	}

	accountID, err := strconv.ParseUint(c.PostForm("accountId"), 10, 64)
	if err != nil || accountID == 0 {
		return dto.TransactionInput{}, fmt.Errorf("บัญชีไม่ถูกต้อง")
	}

	transactionDate, err := time.Parse(
		"2006-01-02",
		c.PostForm("transactionDate"),
	)
	if err != nil {
		return dto.TransactionInput{}, fmt.Errorf("วันที่ไม่ถูกต้อง")
	}

	return dto.TransactionInput{
		Type:            transactionType,
		Amount:          amount,
		Category:        category,
		AccountID:       uint(accountID),
		Title:           title,
		Note:            note,
		TransactionDate: transactionDate,
	}, nil
}

func parseTransactionFilter(
	c *gin.Context,
) (dto.TransactionFilter, error) {
	var filter dto.TransactionFilter

	startDateString := c.Query("startDate")
	endDateString := c.Query("endDate")

	if startDateString != "" && endDateString != "" {
		startDate, err := time.Parse(
			"2006-01-02",
			startDateString,
		)
		if err != nil {
			return filter, fmt.Errorf("วันที่เริ่มต้นไม่ถูกต้อง")
		}

		endDate, err := time.Parse(
			"2006-01-02",
			endDateString,
		)
		if err != nil {
			return filter, fmt.Errorf("วันที่สิ้นสุดไม่ถูกต้อง")
		}

		filter.StartDate = startDate
		filter.EndDate = endDate.AddDate(0, 0, 1)

	} else {
		year, err := strconv.Atoi(c.Query("year"))
		if err != nil || year < 1 {
			return filter, fmt.Errorf("ปีไม่ถูกต้อง")
		}

		month, err := strconv.Atoi(c.Query("month"))
		if err != nil || month < 1 || month > 12 {
			return filter, fmt.Errorf("เดือนไม่ถูกต้อง")
		}

		filter.StartDate = time.Date(
			year,
			time.Month(month),
			1,
			0, 0, 0, 0,
			time.Local,
		)

		filter.EndDate = filter.StartDate.AddDate(0, 1, 0)
	}

	accountIDString := c.Query("accountId")
	if accountIDString != "" {
		accountID, err := strconv.ParseUint(
			accountIDString,
			10,
			64,
		)
		if err != nil || accountID == 0 {
			return filter, fmt.Errorf("account id ไม่ถูกต้อง")
		}

		id := uint(accountID)
		filter.AccountID = &id
	}

	filter.Category = c.Query("category")

	return filter, nil
}