package models
import "time"
// import "gorm.io/gorm"

type User struct {
	// gorm.Model
	ID uint `gorm:"primaryKey"`
	Name     string `gorm:"size:100;not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	CreatedAt time.Time

	Accounts	[]Account `gorm:"foreignKey:UserID"`
	Budgets		[]Budget `gorm:"foreignKey:UserID"`
	SavingGoals	[]SavingGoal `gorm:"foreignKey:UserID"`
}