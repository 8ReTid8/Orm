package models

// import "time"

type Bank struct {
	ID uint `gorm:"primaryKey"`

	Name string `gorm:"size:255;not null"`
	Logo string `gorm:"size:255"`

	Transactions []Transaction `gorm:"foreignKey:BankID"`
}