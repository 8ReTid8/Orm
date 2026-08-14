package models


type Category struct {
	ID uint `gorm:"primaryKey"`
	Type string `gorm:"size:20;not null"`
	Name string `gorm:"size:100;not null"`
	Icon string `gorm:"size:255"`

	// Transactions []Transaction `gorm:"foreignKey:CategoryID"`
	// Budgets      []Budget      `gorm:"foreignKey:CategoryID"`
}