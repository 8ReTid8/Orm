package models

import "time"

type Budget struct {
	ID uint `gorm:"primaryKey"`

	UserID uint
	User User

	Amount float64 `gorm:"not null"`
	Category string `gorm:"size:20;not null"`
	StartDate time.Time
	EndDate time.Time

	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt *time.Time `gorm:"index"`
}