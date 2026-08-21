package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"backend/internal/database"
	"backend/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateTransaction(c *gin.Context) {
	// 1. User จาก JWT
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user ID",
		})
		return
	}

	// 2. รับค่าจาก multipart/form-data
	transactionType := c.PostForm("type")
	amountString := c.PostForm("amount")
	category := c.PostForm("category")
	accountIDString := c.PostForm("accountId")
	title := c.PostForm("title")
	note := c.PostForm("note")
	transactionDateString := c.PostForm("transactionDate")
	fmt.Printf(
		"transactionType = [%s]\n",
		transactionType,
	)
	if transactionType != "income" && transactionType != "expense" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ประเภท transaction ไม่ถูกต้อง",
		})
		return
	}

	if title == "" || category == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "กรุณากรอกข้อมูลให้ครบ",
		})
		return
	}

	// amount string -> float64
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil || amount <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "จำนวนเงินไม่ถูกต้อง",
		})
		return
	}

	accountID64, err := strconv.ParseUint(accountIDString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "บัญชีไม่ถูกต้อง",
		})
		return
	}

	accountID := uint(accountID64)

	// วันที่
	transactionDate, err := time.Parse(
		"2006-01-02",
		transactionDateString,
	)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "วันที่ไม่ถูกต้อง",
		})
		return
	}

	var account models.Account

	if err := database.DB.
		Where("id = ? AND user_id = ?", accountID, userID).
		First(&account).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ไม่พบบัญชี",
		})
		return
	}

	// 4. Upload slip (optional)
	slipPath := ""

	file, err := c.FormFile("slipImage")

	if err == nil {
		uploadDir := "uploads/slips"

		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ไม่สามารถสร้างโฟลเดอร์รูปได้",
			})
			return
		}

		extension := filepath.Ext(file.Filename)

		filename := fmt.Sprintf(
			"%d_%d%s",
			userID,
			time.Now().UnixNano(),
			extension,
		)

		filePath := filepath.Join(uploadDir, filename)

		if err := c.SaveUploadedFile(file, filePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ไม่สามารถบันทึกรูปสลิปได้",
			})
			return
		}

		slipPath = "/" + filepath.ToSlash(filePath)
	}

	// 5. Create Transaction
	transaction := models.Transaction{
		AccountID:       accountID,
		Category:        category,
		Type:            transactionType,
		Amount:          amount,
		Title:           title,
		Note:            note,
		Image:           slipPath,
		TransactionDate: transactionDate,
	}

	// if err := database.DB.Create(&transaction).Error; err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "ไม่สามารถบันทึกรายการได้",
	// 	})
	// 	return
	// }
	err = database.DB.Transaction(func(tx *gorm.DB) error {

		// expense ต้องเช็กยอดก่อน
		if transactionType == "expense" {
			if account.Balance < amount {
				return fmt.Errorf("ยอดเงินในบัญชีไม่เพียงพอ")
			}

			if err := tx.Model(&models.Account{}).
				Where("id = ?", accountID).
				Update(
					"balance",
					gorm.Expr("balance - ?", amount),
				).Error; err != nil {
				return err
			}
		}

		// income
		if transactionType == "income" {
			if err := tx.Model(&models.Account{}).
				Where("id = ?", accountID).
				Update(
					"balance",
					gorm.Expr("balance + ?", amount),
				).Error; err != nil {
				return err
			}
		}

		// สร้าง transaction
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		if err.Error() == "ยอดเงินในบัญชีไม่เพียงพอ" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถบันทึกรายการได้",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message":     "เพิ่มรายการสำเร็จ",
		"transaction": transaction,
	})
}

func GetTransactions(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Invalid user ID",
		})
		return
	}

	yearString := c.Query("year")
	monthString := c.Query("month")

	year, err := strconv.Atoi(yearString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ปีไม่ถูกต้อง",
		})
		return
	}

	month, err := strconv.Atoi(monthString)
	if err != nil || month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "เดือนไม่ถูกต้อง",
		})
		return
	}

	startDate := time.Date(
		year,
		time.Month(month),
		1,
		0, 0, 0, 0,
		time.Local,
	)

	endDate := startDate.AddDate(0, 1, 0)

	var transactions []models.Transaction

	if err := database.DB.
		Preload("Account").
		Joins("JOIN accounts ON accounts.id = transactions.account_id").
		Where("accounts.user_id = ?", userID).
		Where(
			"transactions.transaction_date >= ? AND transactions.transaction_date < ?",
			startDate,
			endDate,
		).
		Order("transactions.transaction_date DESC").
		Find(&transactions).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "ไม่สามารถโหลดรายการได้",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}
