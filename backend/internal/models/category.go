package models


type Category struct {
	ID uint `gorm:"primaryKey"`

	Name string `gorm:"size:100;not null"`
	Icon string `gorm:"size:255"`

	// Transactions []Transaction `gorm:"foreignKey:CategoryID"`
	// Budgets      []Budget      `gorm:"foreignKey:CategoryID"`
}