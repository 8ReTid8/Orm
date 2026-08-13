package handlers

import (
	"net/http"
	"strings"

	"backend/internal/database"
	"backend/internal/models"
	"backend/internal/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterRequest struct {
	// Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var request RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ข้อมูลไม่ถูกต้อง",
			"error":   err.Error(),
		})
		return
	}

	email := strings.ToLower(strings.TrimSpace(request.Email))
	// name := strings.TrimSpace(request.Name)

	var existingUser models.User

	result := database.DB.
		Where("email = ?", email).
		First(&existingUser)

	if result.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "อีเมลนี้ถูกใช้งานแล้ว",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถสร้างบัญชีได้",
		})
		return
	}

	user := models.User{
		// Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถบันทึกผู้ใช้ได้",
		})
		return
	}

	// token, err := utils.GenerateToken(user.ID)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถสร้าง token ได้",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusCreated, gin.H{
		"message": "สมัครสมาชิกสำเร็จ",
		// "token":   token,
		// "user": gin.H{
		// 	"id":    user.ID,
		// 	"email": user.Email,
		// },
	})
}

func Login(c *gin.Context) {
	var request LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	email := strings.ToLower(strings.TrimSpace(request.Email))

	var user models.User

	if err := database.DB.
		Where("email = ?", email).
		First(&user).Error; err != nil {

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(request.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถสร้าง token ได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "เข้าสู่ระบบสำเร็จ",
		"token":   token,
		"user": gin.H{
			"id":    user.ID,
			// "name":  user.Name,
			"email": user.Email,
		},
	})
}