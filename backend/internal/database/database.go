package database

import (
	"fmt"
	"log"
	"os"

	"backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Account{},
		&models.Transaction{},
		&models.Budget{},
		&models.SavingGoal{},
	)

	if err != nil {
		log.Fatal(err)
	}

	// return db
	DB = db

	log.Println("Database connected")
}