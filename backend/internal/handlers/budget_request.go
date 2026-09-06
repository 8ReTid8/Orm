package handlers

import (
	"fmt"
	"strconv"
	"time"

	"backend/internal/dto"
	"github.com/gin-gonic/gin"
)

func parseBudgetForm(
	c *gin.Context,
) (dto.BudgetInput, error) {
	category := c.PostForm("category")

	amount, err := strconv.ParseFloat(
		c.PostForm("amount"),
		64,
	)
	if err != nil || amount <= 0 {
		return dto.BudgetInput{}, fmt.Errorf("จำนวนเงินไม่ถูกต้อง")
	}

	if category == "" {
		return dto.BudgetInput{}, fmt.Errorf("กรุณากรอกข้อมูลให้ครบ")
	}

	accountID64, err := strconv.ParseUint(
		c.PostForm("accountId"),
		10,
		64,
	)
	if err != nil || accountID64 == 0 {
		return dto.BudgetInput{}, fmt.Errorf("บัญชีไม่ถูกต้อง")
	}

	startDate, err := time.Parse(
		"2006-01-02",
		c.PostForm("startDate"),
	)
	if err != nil {
		return dto.BudgetInput{}, fmt.Errorf("วันที่เริ่มต้นไม่ถูกต้อง")
	}

	endDate, err := time.Parse(
		"2006-01-02",
		c.PostForm("endDate"),
	)
	if err != nil {
		return dto.BudgetInput{}, fmt.Errorf("วันที่สิ้นสุดไม่ถูกต้อง")
	}

	if endDate.Before(startDate) {
		return dto.BudgetInput{}, fmt.Errorf(
			"วันที่สิ้นสุดต้องไม่ก่อนวันที่เริ่มต้น",
		)
	}

	return dto.BudgetInput{
		Amount:    amount,
		Category:  category,
		AccountID: uint(accountID64),
		StartDate: startDate,
		EndDate:   endDate,
	}, nil
}

func parseBudgetFilter(
	c *gin.Context,
) (dto.BudgetFilter, error) {
	var filter dto.BudgetFilter

	yearStr := c.Query("year")
	monthStr := c.Query("month")

	if yearStr != "" {
		year, err := strconv.Atoi(yearStr)
		if err != nil || year <= 0 {
			return filter, fmt.Errorf("ปีไม่ถูกต้อง")
		}

		filter.Year = &year
	}

	if monthStr != "" {
		if filter.Year == nil {
			return filter, fmt.Errorf("ต้องเลือกปีก่อนเลือกเดือน")
		}

		month, err := strconv.Atoi(monthStr)
		if err != nil || month < 1 || month > 12 {
			return filter, fmt.Errorf("เดือนไม่ถูกต้อง")
		}

		filter.Month = &month
	}

	status := c.Query("status")
	if status != "" && status != "active" && status != "ended" {
		return filter, fmt.Errorf("status ไม่ถูกต้อง")
	}
	filter.Status = status

	accountIDStr := c.Query("accountId")
	if accountIDStr != "" {
		accountID, err := strconv.ParseUint(accountIDStr, 10, 64)
		if err != nil || accountID == 0 {
			return filter, fmt.Errorf("account id ไม่ถูกต้อง")
		}

		id := uint(accountID)
		filter.AccountID = &id
	}

	return filter, nil
}