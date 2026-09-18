package handlers

import (
	"fmt"
	"strconv"

	"backend/internal/dto"

	"github.com/gin-gonic/gin"
)

func parseSummaryFilter(
	c *gin.Context,
) (dto.SummaryFilter, error) {
	var filter dto.SummaryFilter

	year, err := strconv.Atoi(c.Query("year"))
	if err != nil || year < 1 {
		return filter, fmt.Errorf("ปีไม่ถูกต้อง")
	}
	filter.Year = year

	monthString := c.Query("month")

	if monthString != "" {
		month, err := strconv.Atoi(monthString)

		if err != nil || month < 1 || month > 12 {
			return filter, fmt.Errorf("เดือนไม่ถูกต้อง")
		}

		filter.Month = &month
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

	return filter, nil
}