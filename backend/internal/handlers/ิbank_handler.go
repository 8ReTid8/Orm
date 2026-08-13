package handlers

import (
	"net/http"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetBanks(c *gin.Context) {
	var banks []models.Bank

	if err := database.DB.
		Order("name ASC").
		Find(&banks).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดข้อมูลธนาคารได้",
		})
		return
	}

	c.JSON(http.StatusOK, banks)
}