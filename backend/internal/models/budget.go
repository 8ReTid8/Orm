package models

import "time"

type Budget struct {
	ID uint `gorm:"primaryKey"`

	UserID uint
	User   User

	AccountID uint
	Account   Account

	Amount   float64 `gorm:"not null"`
	Category string  `gorm:"size:20;not null"`

	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null"`

	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt *time.Time `gorm:"index"`
}
