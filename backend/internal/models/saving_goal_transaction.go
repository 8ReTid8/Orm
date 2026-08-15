package models
type SavingGoalTransaction struct {
	ID uint `gorm:"primaryKey"`

	SavingGoalID uint
	SavingGoal   SavingGoal

	AccountID uint
	Account   Account

	Amount float64 `gorm:"not null"`

	// CreatedAt time.Time
}