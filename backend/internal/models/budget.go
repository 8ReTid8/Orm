package models

// import "time"

type Budget struct {
	ID uint `gorm:"primaryKey"`

	UserID uint
	User User

	LimitAmount float64 `gorm:"not null"`
	Category string `gorm:"size:20;not null"`
	Month int `gorm:"not null"`
	Year int `gorm:"not null"`

	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt *time.Time `gorm:"index"`
}