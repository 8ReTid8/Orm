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
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	Role      string `gorm:"type:varchar(20);not null;default:'user'"`

	Accounts    []Account    `gorm:"foreignKey:UserID"`
	Budgets     []Budget     `gorm:"foreignKey:UserID"`
	SavingGoals []SavingGoal `gorm:"foreignKey:UserID"`
}
