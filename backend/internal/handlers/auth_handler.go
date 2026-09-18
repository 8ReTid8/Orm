package handlers

import (
	"errors"
	"net/http"
	
	"backend/internal/dto"
	"backend/internal/services"
	
	"github.com/gin-gonic/gin"
)


type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(
	service *services.AuthService,
) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// func Register(c *gin.Context) {
func (h *AuthHandler) Register(c *gin.Context) {
	// var request RegisterRequest
	var request dto.RegisterRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}

	err := h.service.Register(
		c.Request.Context(),
		request.Email,
		request.Password,
	)

	switch {
	case errors.Is(err, services.ErrEmailAlreadyUsed):
		c.JSON(http.StatusConflict, gin.H{
			"message": "อีเมลนี้ถูกใช้งานแล้ว",
		})
		return

	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถสร้างบัญชีได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "สมัครสมาชิกสำเร็จ",
	})
	// email := strings.ToLower(strings.TrimSpace(request.Email))

	// var existingUser models.User

	// result := database.DB.
	// 	Where("email = ?", email).
	// 	First(&existingUser)

	// if result.Error == nil {
	// 	c.JSON(http.StatusConflict, gin.H{
	// 		"message": "อีเมลนี้ถูกใช้งานแล้ว",
	// 	})
	// 	return
	// }

	// hashedPassword, err := bcrypt.GenerateFromPassword(
	// 	[]byte(request.Password),
	// 	bcrypt.DefaultCost,
	// )

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถสร้างบัญชีได้",
	// 	})
	// 	return
	// }

	// user := models.User{
	// 	// Name:     name,
	// 	Email:    email,
	// 	Password: string(hashedPassword),
	// 	Role:     "user",
	// }

	// if err := database.DB.Create(&user).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถบันทึกผู้ใช้ได้",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{
	// 	"message": "สมัครสมาชิกสำเร็จ",
	// })
}

// func Login(c *gin.Context) {
func (h *AuthHandler) Login(c *gin.Context) {
	// var request LoginRequest
	var request dto.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ข้อมูลไม่ถูกต้อง",
		})
		return
	}
	result, err := h.service.Login(
		c.Request.Context(),
		request.Email,
		request.Password,
	)

	switch {
	case errors.Is(err, services.ErrInvalidCredentials):
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
		})
		return

	case err != nil:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถสร้าง token ได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "เข้าสู่ระบบสำเร็จ",
		"token":   result.Token,
		"user":    dto.ToAuthUserResponse(*result.User),
	})
	// email := strings.ToLower(strings.TrimSpace(request.Email))

	// var user models.User

	// if err := database.DB.
	// 	Where("email = ?", email).
	// 	First(&user).Error; err != nil {

	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
	// 	})
	// 	return
	// }

	// err := bcrypt.CompareHashAndPassword(
	// 	[]byte(user.Password),
	// 	[]byte(request.Password),
	// )

	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{
	// 		"message": "อีเมลหรือรหัสผ่านไม่ถูกต้อง",
	// 	})
	// 	return
	// }

	// token, err := utils.GenerateToken(user.ID, user.Role)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถสร้าง token ได้",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "เข้าสู่ระบบสำเร็จ",
	// 	"token":   token,
	// 	"user": gin.H{
	// 		"id": user.ID,
	// 		// "name":  user.Name,
	// 		"email": user.Email,
	// 		"role":  user.Role,
	// 	},
	// })
}
