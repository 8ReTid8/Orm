package handlers

import (
    "fmt"
    "strconv"

    "backend/internal/dto"

    "github.com/gin-gonic/gin"
)

func parsePageRequest(
    c *gin.Context,
) (dto.PageRequest, error) {
    page := 1
    limit := 20

    if pageString := c.Query("page"); pageString != "" {
        value, err := strconv.Atoi(pageString)
        if err != nil || value < 1 {
            return dto.PageRequest{}, fmt.Errorf("หน้าที่ต้องการไม่ถูกต้อง")
        }

        page = value
    }

    if limitString := c.Query("limit"); limitString != "" {
        value, err := strconv.Atoi(limitString)
        if err != nil || value < 1 || value > 100 {
            return dto.PageRequest{}, fmt.Errorf("จำนวนรายการต่อหน้าไม่ถูกต้อง")
        }

        limit = value
    }

    return dto.PageRequest{
        Page:  page,
        Limit: limit,
    }, nil
}