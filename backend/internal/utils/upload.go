package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveSlip(
	c *gin.Context,
	userID uint,
) (string, error) {

	file, err := c.FormFile("slipImage")

	// ไม่มีไฟล์ ไม่ถือว่า error
	if err != nil {
		return "", nil
	}

	return saveUploadedSlip(
		c,
		file,
		userID,
	)
}

func saveUploadedSlip(
	c *gin.Context,
	file *multipart.FileHeader,
	userID uint,
) (string, error) {

	uploadDir := "uploads/slips"

	if err := os.MkdirAll(
		uploadDir,
		0755,
	); err != nil {
		return "", err
	}

	extension := filepath.Ext(file.Filename)

	filename := fmt.Sprintf(
		"%d_%d%s",
		userID,
		time.Now().UnixNano(),
		extension,
	)

	filePath := filepath.Join(
		uploadDir,
		filename,
	)

	if err := c.SaveUploadedFile(
		file,
		filePath,
	); err != nil {
		return "", err
	}

	return "/" + filepath.ToSlash(filePath), nil
}