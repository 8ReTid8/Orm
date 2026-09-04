package models

import "time"

type Budget struct {
	ID uint `gorm:"primaryKey"`

	UserID uint
	User   User

	AccountID uint
	Account   Account `json:"account"`

	Amount   float64 `gorm:"not null" json:"amount"`
	Category string  `gorm:"size:20;not null" json:"category"`

	StartDate time.Time `gorm:"not null" json:"startDate"`
	EndDate   time.Time `gorm:"not null" json:"endDate"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	// UpdatedAt time.Time
	// DeletedAt *time.Time `gorm:"index"`
}
