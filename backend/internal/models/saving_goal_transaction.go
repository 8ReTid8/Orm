package models

import "time"

type SavingGoalTransaction struct {
	ID uint `gorm:"primaryKey"`

	SavingGoalID uint
	SavingGoal   SavingGoal

	AccountID uint
	Account   Account

	Amount    float64   `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	// CreatedAt time.Time
}
