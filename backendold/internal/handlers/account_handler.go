package handlers

import (
	"net/http"
	"strings"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
)

type CreateAccountRequest struct {
	Name string `json:"name" binding:"required"`
}

func CreateAccount(c *gin.Context) {
	
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var request CreateAccountRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	name := strings.TrimSpace(request.Name)

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "กรุณากรอกชื่อบัญชี",
		})
		return
	}

	account := models.Account{
		UserID:  userID,
		Name:    name,
		Balance: 0,
	}

	if err := database.DB.Create(&account).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถสร้างบัญชีได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "สร้างบัญชีสำเร็จ",
		"account": account,
	})
}

func GetAccounts(c *gin.Context) {
	
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	var accounts []models.Account

	if err := database.DB.
		Where("user_id = ?", userID).
		Order("created_at ASC").
		Find(&accounts).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดบัญชีได้",
		})
		return
	}

	c.JSON(http.StatusOK, accounts)
}

