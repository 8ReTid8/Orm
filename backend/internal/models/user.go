package models

import "time"

// import "gorm.io/gorm"

type User struct {
	// gorm.Model
	// ID        uint   `gorm:"primaryKey"`
	// Email     string `gorm:"unique;not null"`
	// Password  string `gorm:"not null"`
	ID        uint   `gorm:"primaryKey" json:"id"`
	Email     string `gorm:"size:255;uniqueIndex;not null" json:"email"`
	Password  string `gorm:"size:255;not null" json:"-"`
	CreatedAt time.Time

	Transactions []Transaction `gorm:"foreignKey:UserID"`
	Budgets      []Budget      `gorm:"foreignKey:UserID"`
	SavingGoals  []SavingGoal  `gorm:"foreignKey:UserID"`
}
