package models

// import "time"

type Account struct {
	ID uint `gorm:"primaryKey" json:"id"` 

	UserID uint `gorm:"not null;index" json:"userId"`
	User User `json:"-"`

	Name string `gorm:"size:255;not null" json:"name"`
	Balance float64 `gorm:"not null" json:"balance"`

	Transactions []Transaction `gorm:"foreignKey:AccountID" json:"-"`
} 