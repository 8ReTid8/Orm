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

func (h *AuthHandler) Register(c *gin.Context) {
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
	case errors.Is(err, services.ErrVerificationEmailFailed):
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "สร้างบัญชีแล้ว แต่ไม่สามารถส่งอีเมลยืนยันได้",
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
}

func (h *AuthHandler) Login(c *gin.Context) {
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
	case errors.Is(err, services.ErrEmailNotVerified):
		c.JSON(http.StatusForbidden, gin.H{
			"message": "กรุณายืนยันอีเมลก่อนเข้าสู่ระบบ",
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
}

func (h *AuthHandler) VerifyEmail(
    c *gin.Context,
) {
    token := c.Query("token")

    err := h.service.VerifyEmail(
        c.Request.Context(),
        token,
    )

    switch {
    case errors.Is(err, services.ErrInvalidVerificationToken):
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "ลิงก์ยืนยันไม่ถูกต้องหรือหมดอายุแล้ว",
        })
        return

    case err != nil:
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "ไม่สามารถยืนยันอีเมลได้",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "ยืนยันอีเมลสำเร็จ คุณสามารถเข้าสู่ระบบได้แล้ว",
    })
}