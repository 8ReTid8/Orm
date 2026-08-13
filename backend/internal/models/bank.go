package models

// import "time"

type Bank struct {
	// ID uint `gorm:"primaryKey"`

	// Name string `gorm:"size:255;not null"`
	// Logo string `gorm:"size:255"`

	ID uint `gorm:"primaryKey" json:"id"`

	Name string `gorm:"size:255;not null" json:"name"`
	Logo string `gorm:"size:255;not null" json:"logo"`


	Transactions []Transaction `gorm:"foreignKey:BankID"`
}