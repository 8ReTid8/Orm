package handlers

import (
	"errors"
	"net/http"

	"backend/internal/dto"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
)
type AccountHandler struct {
	service *services.AccountService
}

func NewAccountHandler(
	service *services.AccountService,
) *AccountHandler {
	return &AccountHandler{
		service: service,
	}
}
type CreateAccountRequest struct {
	Name string `json:"name" binding:"required"`
}

// func CreateAccount(c *gin.Context) {
func (h *AccountHandler) CreateAccount(c *gin.Context) {
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
	account, err := h.service.Create(
		c.Request.Context(),
		userID,
		request.Name,
	)

	if errors.Is(err, services.ErrInvalidAccountName) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "กรุณากรอกชื่อบัญชี",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถสร้างบัญชีได้",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "สร้างบัญชีสำเร็จ",
		"account": dto.ToAccountResponse(*account),
	})
	// name := strings.TrimSpace(request.Name)

	// if name == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "กรุณากรอกชื่อบัญชี",
	// 	})
	// 	return
	// }

	// account := models.Account{
	// 	UserID:  userID,
	// 	Name:    name,
	// 	Balance: 0,
	// }

	// if err := database.DB.Create(&account).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถสร้างบัญชีได้",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{
	// 	"message": "สร้างบัญชีสำเร็จ",
	// 	"account": account,
	// })
}

// func GetAccounts(c *gin.Context) {
func (h *AccountHandler) GetAccounts(c *gin.Context) {
	userID, ok := getUserID(c)
	if !ok {
		return
	}

	accounts, err := h.service.Get(
		c.Request.Context(),
		userID,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดบัญชีได้",
		})
		return
	}
	
	responses := make(
		[]dto.AccountResponse,
		0,
		len(accounts),
	)

	for _, account := range accounts {
		responses = append(
			responses,
			dto.ToAccountResponse(account),
		)
	}

	c.JSON(http.StatusOK, responses)

	// var accounts []models.Account

	// if err := database.DB.
	// 	Where("user_id = ?", userID).
	// 	Order("created_at ASC").
	// 	Find(&accounts).Error; err != nil {

	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถโหลดบัญชีได้",
	// 	})
	// 	return
	// }

	// c.JSON(http.StatusOK, accounts)
}

