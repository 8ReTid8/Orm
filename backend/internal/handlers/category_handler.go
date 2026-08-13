package handlers

import (
	"net/http"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	var categories []models.Category

	if err := database.DB.
		Order("name ASC").
		Find(&categories).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดหมวดหมู่ได้",
		})
		return
	}

	c.JSON(http.StatusOK, categories)
}