package models


type Category struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Type string `gorm:"size:20;not null" json:"type"`
	Name string `gorm:"size:100;not null" json:"name"`
	Icon string `gorm:"size:255" json:"icon"`

	// Transactions []Transaction `gorm:"foreignKey:CategoryID"`
	// Budgets      []Budget      `gorm:"foreignKey:CategoryID"`
}