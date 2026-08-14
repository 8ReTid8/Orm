package models

// import "time"

type Account struct {
	ID uint `gorm:"primaryKey"`

	UserID uint
	User User

	Name string `gorm:"size:255;not null"`
	Balance float64 `gorm:"not null"`

	Transactions []Transaction `gorm:"foreignKey:AccountID"`
} 