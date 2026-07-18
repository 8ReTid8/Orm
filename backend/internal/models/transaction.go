package models

import "time"


type Transaction struct {
	
	AccountID uint
	Account Account

	ID uint `gorm:"primaryKey"`
	Title	string `gorm:"size:255;not null"`
	Note	string `gorm:"type:text"`
	Amount float64 `gorm:"not null"`
	Category string `gorm:"size:20;not null"`
	Type string `gorm:"size:20;not null"`
	TransactionDate time.Time `gorm:"not null"`
	Image string `gorm:"size:255"`

}