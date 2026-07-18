package models

import "time"

type SavingGoal struct {
	ID uint `gorm:"primaryKey"`

	UserID uint
	User User

	Name string `gorm:"size:255;not null"`

	TargetAmount float64 `gorm:"not null"`

	CurrentAmount float64 `gorm:"default:0"`

	Deadline *time.Time

	Status string `gorm:"size:50;default:'ongoing'"`

}