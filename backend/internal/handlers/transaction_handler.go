package handlers

// import (
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"time"

// 	"backend/internal/database"
// 	"backend/internal/models"

// 	"github.com/gin-gonic/gin"
// )

// func CreateTransaction(c *gin.Context) {
// 	// 1. User จาก JWT
// 	userIDValue, exists := c.Get("userID")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "Unauthorized",
// 		})
// 		return
// 	}

// 	userID, ok := userIDValue.(uint)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "Invalid user ID",
// 		})
// 		return
// 	}

// 	// 2. รับค่าจาก multipart/form-data
// 	transactionType := c.PostForm("type")
// 	amountString := c.PostForm("amount")
// 	category := c.PostForm("category")
// 	bankIDString := c.PostForm("bankId")
// 	title := c.PostForm("title")
// 	note := c.PostForm("note")
// 	transactionDateString := c.PostForm("transactionDate")

// 	if transactionType != "income" && transactionType != "expense" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "ประเภท transaction ไม่ถูกต้อง",
// 		})
// 		return
// 	}

// 	if title == "" || category == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "กรุณากรอกข้อมูลให้ครบ",
// 		})
// 		return
// 	}

// 	// amount string -> float64
// 	amount, err := strconv.ParseFloat(amountString, 64)
// 	if err != nil || amount <= 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "จำนวนเงินไม่ถูกต้อง",
// 		})
// 		return
// 	}

// 	// bankId string -> uint
// 	bankID64, err := strconv.ParseUint(bankIDString, 10, 64)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "ธนาคารไม่ถูกต้อง",
// 		})
// 		return
// 	}

// 	bankID := uint(bankID64)

// 	// วันที่
// 	transactionDate, err := time.Parse(
// 		"2006-01-02",
// 		transactionDateString,
// 	)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "วันที่ไม่ถูกต้อง",
// 		})
// 		return
// 	}

// 	// 3. เช็ก Bank ก่อน
// 	var bank models.Bank

// 	if err := database.DB.First(&bank, bankID).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": "ไม่พบธนาคาร",
// 		})
// 		return
// 	}

// 	// 4. Upload slip (optional)
// 	slipPath := ""

// 	file, err := c.FormFile("slipImage")

// 	if err == nil {
// 		uploadDir := "uploads/slips"

// 		if err := os.MkdirAll(uploadDir, 0755); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"message": "ไม่สามารถสร้างโฟลเดอร์รูปได้",
// 			})
// 			return
// 		}

// 		extension := filepath.Ext(file.Filename)

// 		filename := fmt.Sprintf(
// 			"%d_%d%s",
// 			userID,
// 			time.Now().UnixNano(),
// 			extension,
// 		)

// 		filePath := filepath.Join(uploadDir, filename)

// 		if err := c.SaveUploadedFile(file, filePath); err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"message": "ไม่สามารถบันทึกรูปสลิปได้",
// 			})
// 			return
// 		}

// 		slipPath = "/" + filepath.ToSlash(filePath)
// 	}

// 	// 5. Create Transaction
// 	transaction := models.Transaction{
// 		UserID:          userID,
// 		BankID:          bankID,
// 		Category:        category,
// 		Type:            models.TransactionType(transactionType),
// 		Amount:          amount,
// 		Title:           title,
// 		Note:            note,
// 		SlipImage:       slipPath,
// 		TransactionDate: transactionDate,
// 	}

// 	if err := database.DB.Create(&transaction).Error; err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"message": "ไม่สามารถบันทึกรายการได้",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"message":     "เพิ่มรายการสำเร็จ",
// 		"transaction": transaction,
// 	})
// }