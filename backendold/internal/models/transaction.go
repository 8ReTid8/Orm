package models

import "time"

type Transaction struct {

	AccountID uint    `gorm:"not null;index" json:"accountId"`
	// Account   Account `json:"account"`
	Account   Account `json:"-"`


	ID              uint      `gorm:"primaryKey" json:"id"`
	Title           string    `gorm:"size:255;not null" json:"title"`
	Note            string    `gorm:"type:text" json:"note"`
	Amount          float64   `gorm:"not null" json:"amount"`
	Category        string    `gorm:"size:20;not null" json:"category"`
	Type            string    `gorm:"size:20;not null" json:"type"`
	TransactionDate time.Time `gorm:"not null" json:"transactionDate"`
	Image           string    `gorm:"size:255" json:"image"`

	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
}
